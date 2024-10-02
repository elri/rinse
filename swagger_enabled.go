//go:build debug || race

package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/linkdata/rinse/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/linkdata/rinse/rinser"
	httpSwagger "github.com/swaggo/http-swagger"
)

func maybeSwagger(listenUrl string) {
	docs.SwaggerInfo.Version = strings.TrimPrefix(rinser.PkgVersion, "v")
	if os.Getuid() == 0 {
		if strings.HasSuffix(listenUrl, ":80") {
			listenUrl = strings.TrimSuffix(listenUrl, ":80") + ":8080"
		}
		if strings.HasSuffix(listenUrl, ":443") {
			listenUrl = strings.TrimSuffix(listenUrl, ":443") + ":8443"
		}
	}
	docs.SwaggerInfo.Host = listenUrl
	http.DefaultServeMux.Handle("GET /api/", httpSwagger.Handler(
		httpSwagger.URL(listenUrl+"/docs/swagger.json"),
	))
}
