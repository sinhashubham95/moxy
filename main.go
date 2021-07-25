package main

import (
	"fmt"
	"github.com/valyala/fasthttp"

	"github.com/sinhashubham95/moxy/flags"
)

func main() {
	if flags.TLSEnabled() {
		_ = fasthttp.ListenAndServeTLS(fmt.Sprintf(":%d", flags.Port()), flags.CertFilePath(),
			flags.KeyFilePath(), getFastHTTPHandler())
	} else {
		_ = fasthttp.ListenAndServe(fmt.Sprintf(":%d", flags.Port()), getFastHTTPHandler())
	}
}
