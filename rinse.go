package rinse

//go:generate go run github.com/cparta/makeversion/cmd/mkver@latest -name rinse -out version.gen.go

import (
	"bytes"
	"embed"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/linkdata/deadlock"
	"github.com/linkdata/jaws"
	"github.com/linkdata/jaws/staticserve"
	"github.com/linkdata/webserv"
)

//go:embed assets
var assetsFS embed.FS

const PodmanImage = "ghcr.io/linkdata/rinse:latest"

type Rinse struct {
	Config     *webserv.Config
	Jaws       *jaws.Jaws
	PodmanBin  string
	RunscBin   string
	FaviconURI string
	mu         deadlock.Mutex // protects following
	jobs       []*Job
}

func New(cfg *webserv.Config, mux *http.ServeMux, jw *jaws.Jaws) (rns *Rinse, err error) {
	var tmpl *template.Template
	var faviconuri string
	if tmpl, err = template.New("").ParseFS(assetsFS, "assets/ui/*.html"); err == nil {
		jw.AddTemplateLookuper(tmpl)
		var extraFiles []string
		addStaticFiles := func(filename string, ss *staticserve.StaticServe) (err error) {
			uri := path.Join("/static", ss.Name)
			if strings.HasSuffix(filename, "favicon.png") {
				faviconuri = uri
			}
			extraFiles = append(extraFiles, uri)
			mux.Handle(uri, ss)
			return
		}
		if err = staticserve.WalkDir(assetsFS, "assets/static", addStaticFiles); err == nil {
			if err = jw.GenerateHeadHTML(extraFiles...); err == nil {
				var podmanbin string
				if podmanbin, err = exec.LookPath("podman"); err == nil {
					slog.Info("podman", "bin", podmanbin)
					var runscbin string
					if s, e := exec.LookPath("runsc"); e == nil {
						if os.Getuid() == 0 && cfg.User == "" {
							runscbin = s
							slog.Info("gVisor", "bin", runscbin)
						} else {
							slog.Warn("gVisor needs root", "bin", s)
						}
					} else {
						slog.Info("gVisor not found", "err", e)
					}
					rns = &Rinse{
						Config:     cfg,
						Jaws:       jw,
						PodmanBin:  podmanbin,
						RunscBin:   runscbin,
						FaviconURI: faviconuri,
					}
					rns.addRoutes(mux)
				}
			}
		}
	}

	return
}

func (rns *Rinse) addRoutes(mux *http.ServeMux) {
	mux.Handle("GET /{$}", rns.Jaws.Handler("index.html", rns))
	mux.Handle("GET /setup/{$}", rns.Jaws.Handler("setup.html", rns))
	mux.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		slog.Info(r.RequestURI)
	})
}

func (rns *Rinse) Close() {
}

func (rns *Rinse) Pull() (err error) {
	slog.Info("podman pull", "image", PodmanImage)
	var out []byte
	cmd := exec.Command(rns.PodmanBin, "pull", PodmanImage)
	if out, err = cmd.CombinedOutput(); err != nil {
		for _, line := range bytes.Split(bytes.TrimSpace(out), []byte{'\n'}) {
			slog.Error("podman", "msg", string(bytes.TrimSpace(line)))
		}
	}
	return
}

func (rns *Rinse) PkgName() string {
	return PkgName
}

func (rns *Rinse) PkgVersion() string {
	return PkgVersion
}

func (rns *Rinse) NewJob(name, lang string) (job *Job, err error) {
	if job, err = NewJob(name, lang, rns.PodmanBin, rns.RunscBin); err == nil {
		rns.mu.Lock()
		rns.jobs = append(rns.jobs, job)
		rns.mu.Unlock()
	}
	return
}

// JawsContains implements jaws.Container.
func (rns *Rinse) JawsContains(e *jaws.Element) (contents []jaws.UI) {
	rns.mu.Lock()
	for _, job := range rns.jobs {
		contents = append(contents, jaws.NewTemplate("job.html", job))
	}
	rns.mu.Unlock()
	return
}
