package controllers

import (
	"github.com/valyala/fasthttp"
	"net/http"

	"github.com/sinhashubham95/moxy/models"
	"github.com/sinhashubham95/moxy/persistence/entities"
)

// HandleUnMock is used to un mock and endpoint
func HandleUnMock(ctx *fasthttp.RequestCtx) {
	var request models.UnMockRequest
	err := DecodeJSON(ctx.PostBody(), &request)
	if err != nil {
		// invalid request body
		ctx.Error("invalid json request body provided", http.StatusBadRequest)
		return
	}
	err = PersistenceDelete(&entities.Mock{
		Tag:    request.Tag,
		Method: request.Method,
		Path:   request.Path,
	})
	if err != nil {
		// unable to save the mock configurations
		ctx.Error(err.Error(), http.StatusInternalServerError)
		return
	}
	// success
	ctx.SetStatusCode(http.StatusOK)
}
