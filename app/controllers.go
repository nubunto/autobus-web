// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/nubunto/autobus-web/design
// --out=$(GOPATH)/src/github.com/nubunto/autobus-web
// --version=v1.1.0-dirty
//
// API "autobus-web": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// GPSController is the controller interface for the GPS actions.
type GPSController interface {
	goa.Muxer
	Show(*ShowGPSContext) error
}

// MountGPSController "mounts" a GPS resource controller on the given service.
func MountGPSController(service *goa.Service, ctrl GPSController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowGPSContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/gps", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "GPS", "action", "Show", "route", "GET /gps")
}
