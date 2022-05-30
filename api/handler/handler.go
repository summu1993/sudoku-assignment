package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

/*
 	This reverse proxy will forward the request so the same network no external RPC calls
	UseCase:
	  1. If my service was deplpoyed somewhere else then this proxy architecture would be a scalable approach
	  2. right now reverse proxy will loop back to the same host and port but different end point
*/
func ReverseProxy(proxyUrl string) gin.HandlerFunc {

	remote, err := url.Parse(proxyUrl)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.URL.Scheme = remote.Scheme
			req.Host = remote.Host
			req.URL.Host = remote.Host
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func SolveSudoku() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, response)
	}
}

func CheckSudokuValidity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.JSON(http.StatusOK, response)
	}
}
