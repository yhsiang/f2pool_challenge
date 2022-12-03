package handlers

import (
	_ "embed"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/yhsiang/f2pool-challenge/models"
	"github.com/yhsiang/f2pool-challenge/restapi/operations/root"
)

//go:embed version.txt
var version string

func Root(params root.RootParams) middleware.Responder {
	kubernetes := new(bool)
	*kubernetes = false

	return root.NewRootOK().WithPayload(&models.HandlerRootResponse{
		Date:       time.Now().Unix(),
		Kubernetes: kubernetes,
		Version:    version,
	})
}
