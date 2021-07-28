package controllers

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

// HandleBase is used to handle the request to the base endpoint
func HandleBase(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBodyString("Naruto rocks, and so does Moxy!!")
}
