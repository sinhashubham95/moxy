package controllers

import (
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/models"
	"github.com/sinhashubham95/moxy/persistence"
	"github.com/sinhashubham95/moxy/persistence/entities"
	"github.com/valyala/fasthttp"
	"net/http"
)

// HandleMock is used to mock and endpoint
func HandleMock(ctx *fasthttp.RequestCtx) {
	var request models.MockRequest
	err := commons.DecodeJSON(ctx.PostBody(), &request)
	if err != nil {
		// invalid request body
		ctx.Error("invalid json request body provided", http.StatusBadRequest)
		return
	}
	err = (&request).Validate()
	if err != nil {
		// invalid request
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}
	(&request).Default()
	// now time to save it
	err = persistence.Save(&entities.Mock{
		Tag:    request.Tag,
		Method: request.Method,
		Path:   request.Path,
		Status: request.ResponseStatus,
		Body:   request.ResponseBody,
	})
	if err != nil {
		// unable to save the mock configurations
		ctx.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	// success
	ctx.SetStatusCode(http.StatusOK)
}
