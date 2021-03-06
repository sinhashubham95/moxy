package main

import (
	"github.com/sinhashubham95/go-actuator"
	actuatorCore "github.com/sinhashubham95/go-actuator/core"
	actuatorModels "github.com/sinhashubham95/go-actuator/models"
	"github.com/valyala/fasthttp"
	"net/http"
	"strings"

	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/controllers"
)

func getFastHTTPHandler() fasthttp.RequestHandler {
	actuatorHandler := actuator.GetFastHTTPActuatorHandler(&actuatorModels.Config{
		Prefix: commons.ActuatorPrefix,
	})
	return actuatorCore.WrapFastHTTPHandler(func(ctx *fasthttp.RequestCtx) {
		// actuator
		if strings.HasPrefix(string(ctx.Path()), commons.ActuatorPrefix) && string(ctx.Method()) == http.MethodGet {
			// this means it is an actuator endpoint
			actuatorHandler(ctx)
			return
		}

		// cors
		ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Add("Access-Control-Allow-Headers", "*")
		ctx.Response.Header.Add("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")
		if strings.ToUpper(string(ctx.Method())) == http.MethodOptions {
			ctx.SetStatusCode(http.StatusNoContent)
			return
		}

		switch string(ctx.Path()) {
		case commons.BasePath:
			handle(ctx, http.MethodGet, controllers.HandleBase)
		case commons.MockEndpointPath:
			handle(ctx, http.MethodPost, controllers.HandleMock)
		case commons.UnMockEndpointPath:
			handle(ctx, http.MethodDelete, controllers.HandleUnMock)
		default:
			controllers.Handle(ctx)
		}
	})
}

func handle(ctx *fasthttp.RequestCtx, method string, handler fasthttp.RequestHandler) {
	if string(ctx.Method()) == method {
		handler(ctx)
		return
	}
	ctx.SetStatusCode(http.StatusNotFound)
}
