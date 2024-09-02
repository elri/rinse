package rinse

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"

	"github.com/linkdata/deadlock"
)

type JobState int

const (
	JobNew JobState = iota
	JobStarting
	JobPdfToPPm
	JobTesseract
	JobFinished
	JobFailed
)

type Job struct {
	Name      string
	PodmanBin string
	RunscBin  string
	Workdir   string
	mu        deadlock.Mutex
	state     JobState
	resultCh  chan error
	started   time.Time
	stopped   time.Time
	closed    bool
}

func NewJob(name, podmanbin, runscbin string) (job *Job, err error) {
	var workdir string
	if workdir, err = os.MkdirTemp("", "rinse-"); err == nil {
		job = &Job{
			Name:      filepath.Base(name),
			PodmanBin: podmanbin,
			RunscBin:  runscbin,
			Workdir:   workdir,
			state:     JobNew,
			resultCh:  make(chan error, 1),
		}
	}
	return
}

// podman --runtime=/usr/bin/runsc run --rm -v $DIR_WITH_INPUT_PDF:/var/rinse -it ghcr.io/linkdata/rinse-pdftoppm:latest
// podman run --rm -v $DIR_WITH_OUTPUT_PPM:/var/rinse -it ghcr.io/linkdata/rinse-tesseract:latest

func (job *Job) Start() (err error) {
	if err = job.transition(JobNew, JobStarting); err == nil {
		if job.Name != "input.pdf" {
			err = os.Rename(path.Join(job.Workdir, job.Name), path.Join(job.Workdir, "input.pdf"))
		}
		if err == nil {
			job.mu.Lock()
			job.started = time.Now()
			job.mu.Unlock()
			go job.runPdfToPpm()
		}
	}
	job.checkErr(err)
	return
}

func (job *Job) State() (state JobState) {
	job.mu.Lock()
	state = job.state
	job.mu.Unlock()
	return
}

func (job *Job) ResultCh() (ch <-chan error) {
	job.mu.Lock()
	ch = job.resultCh
	job.mu.Unlock()
	return
}

func (job *Job) checkErrLocked(err error) {
	if err != nil {
		job.state = JobFailed
		job.resultCh <- err
	}
}

func (job *Job) checkErr(err error) {
	if err != nil {
		job.mu.Lock()
		defer job.mu.Unlock()
		job.checkErrLocked(err)
	}
}

func (job *Job) transition(fromState, toState JobState) (err error) {
	job.mu.Lock()
	if job.state == fromState {
		job.state = toState
	} else {
		err = fmt.Errorf("wrong job state (%d)", job.state)
	}
	job.mu.Unlock()
	return
}

func (job *Job) runPdfToPpm() {
	err := job.transition(JobStarting, JobPdfToPPm)

	if err == nil {
		var args []string
		if job.RunscBin != "" {
			args = append(args, "--runtime="+job.RunscBin)
		}
		args = append(args, "--log-level=error", "run", "--rm", "--tty", "-v", job.Workdir+":/var/rinse", "ghcr.io/linkdata/rinse-pdftoppm:latest")
		cmd := exec.Command(job.PodmanBin, args...)
		// we expect no output from pdftoppm
		var output []byte
		output, err = cmd.CombinedOutput()
		if len(output) > 0 {
			slog.Info("pdftoppm", "msg", string(output))
		}
		if err == nil {
			if err = job.runTesseract(); err == nil {
				if err = job.transition(JobTesseract, JobFinished); err == nil {
					job.mu.Lock()
					job.stopped = time.Now()
					ch := job.resultCh
					job.mu.Unlock()
					select {
					case ch <- nil:
					default:
					}
				}
			}
		}
		job.mu.Lock()
		job.stopped = time.Now()
		job.mu.Unlock()
	}

	job.checkErr(err)
}

func (job *Job) runTesseract() (err error) {
	if err = job.transition(JobPdfToPPm, JobTesseract); err == nil {
		if err = os.Remove(path.Join(job.Workdir, "input.pdf")); err == nil {
			var args []string
			args = append(args, "--log-level=error", "run", "--rm", "--tty", "-v", job.Workdir+":/var/rinse", "ghcr.io/linkdata/rinse-tesseract:latest")
			cmd := exec.Command(job.PodmanBin, args...)
			var stdout io.ReadCloser
			if stdout, err = cmd.StdoutPipe(); err == nil {
				if err = cmd.Start(); err == nil {
					fileScanner := bufio.NewScanner(stdout)
					fileScanner.Split(bufio.ScanLines)
					for fileScanner.Scan() {
						fmt.Println(fileScanner.Text())
					}
					err = cmd.Wait()
				}
			}
		}
	}
	return
}

func (job *Job) Result() (err error) {
	if err = os.Rename(path.Join(job.Workdir, "output.pdf"), path.Join(job.Workdir, job.Name)); err == nil {
	}
	return
}

func (job *Job) Close() (err error) {
	job.mu.Lock()
	defer job.mu.Unlock()
	if !job.closed {
		job.closed = true
		close(job.resultCh)
		err = os.RemoveAll(job.Workdir)
	}
	return
}