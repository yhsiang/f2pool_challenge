package handlers

import (
	_ "embed"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/yhsiang/f2pool-challenge/models"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/root"

	"k8s.io/client-go/rest"
)

//go:embed version.txt
var version string

func Root(params root.RootParams) middleware.Responder {
	isInK8s := new(bool)

	_, err := rest.InClusterConfig()
	*isInK8s = err == nil

	return root.NewRootOK().WithPayload(&models.HandlerRootResponse{
		Date:       time.Now().Unix(),
		Kubernetes: isInK8s,
		Version:    version,
	})
}
