package controllers

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"

	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/persistence/entities"
)

// Handle is used to handle the request to any endpoint
func Handle(ctx *fasthttp.RequestCtx) {
	// first get the method, path and tag
	method := string(ctx.Method())
	path := string(ctx.Path())
	tag := string(ctx.Request.Header.Peek(commons.TagHeader))

	// now for this method and path, fetch the entity
	mock := entities.Mock{
		Tag:    tag,
		Method: method,
		Path:   path,
	}

	// try to fetch the mock
	err := PersistenceView(&mock)
	if err == nil {
		// this means that the mock exists
		ctx.SetStatusCode(mock.Status)
		// writing the response here is tricky
		if s, ok := mock.Body.(string); ok {
			// write as string
			ctx.SetContentType(commons.TextStringContentType)
			ctx.SetBodyString(s)
		} else {
			// first parse to bytes
			body, err := EncodeJSON(mock.Body)
			if err == nil {
				// now write
				ctx.SetContentType(commons.ApplicationJSONContentType)
				ctx.SetBody(body)
			}
		}
		// return from here
		return
	}

	// get the url
	url := string(ctx.Request.Header.Peek(commons.URLHeader))
	body := ctx.Request.Body()

	// now that mock does not exist, call the actual endpoint
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(fmt.Sprintf("%s/%s", url, path))
	req.Header.SetMethod(method)
	req.SetBody(body)
	res := fasthttp.AcquireResponse()
	err = fasthttp.Do(req, res)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetBody([]byte(err.Error()))
		return
	}
	ctx.SetStatusCode(res.StatusCode())
	res.Header.CopyTo(&ctx.Response.Header)
	ctx.SetBody(res.Body())
}
