package http

import (
	"github.com/whosonfirst/go-rasterzen/mvt"
	gohttp "net/http"
)

func SVGHandler() (gohttp.HandlerFunc, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		fh, err := GetTileForRequest(req)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		defer fh.Close()

		rsp.Header().Set("Content-Type", "image/svg+xml")
		rsp.Header().Set("Access-Control-Allow-Origin", "*")

		err = mvt.ToSVG(fh, rsp)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		return
	}

	return gohttp.HandlerFunc(fn), nil
}
