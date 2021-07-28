package controllers

import (
	"github.com/valyala/fasthttp"
	"net/http"

	"github.com/sinhashubham95/moxy/models"
	"github.com/sinhashubham95/moxy/persistence/entities"
)

// HandleMock is used to mock and endpoint
func HandleMock(ctx *fasthttp.RequestCtx) {
	var request models.MockRequest
	err := DecodeJSON(ctx.PostBody(), &request)
	if err != nil {
		// invalid request body
		ctx.Error("invalid json request body provided", http.StatusBadRequest)
		return
	}
	(&request).Clean()
	err = (&request).Validate()
	if err != nil {
		// invalid request
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}
	(&request).Default()
	// now time to save it
	err = PersistenceSave(&entities.Mock{
		Tag:           request.Tag,
		Method:        request.Method,
		Path:          request.Path,
		DelayInMillis: request.ResponseDelayInMillis,
		Status:        request.ResponseStatus,
		Body:          request.ResponseBody,
	})
	if err != nil {
		// unable to save the mock configurations
		ctx.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	// success
	ctx.SetStatusCode(http.StatusOK)
}
