package main

import (
	"fmt"
	"github.com/valyala/fasthttp"

	"github.com/sinhashubham95/moxy/flags"
)

func main() {
	err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", flags.Port()), getFastHTTPHandler())
	if err != nil {
		return
	}
}
