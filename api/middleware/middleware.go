package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"sudoku-assignment/response"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		log.Println("request = ", readBody(rdr1)) // Print request body

		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

func Headers() gin.HandlerFunc {
	return func(c *gin.Context) {

		// check for specific header
		userCheck := c.Request.Header.Get("X-User")
		if len(userCheck) > 0 {
			// do something with header
		} else {
			var response response.HttpResponse
			response.ErrorResponse.Message = "X-User header is missing"
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		// continue with the next pending handlers
		c.Next()
	}
}

func CORSHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
