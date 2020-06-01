package main

import (
	"github.com/gin-gonic/gin"
	"edbproxy/mock"
	"fmt"
	"encoding/xml"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/edb", func(context *gin.Context) {
		method := context.Request.FormValue("method")
		switch method {
		case "edbTradeGet":
			eorderNo := context.Request.PostFormValue("eorderId")
			xmlContent:= mock.DeliveryXml(eorderNo)
			bytes, _ := xml.MarshalIndent(xmlContent, "  ", "    ")
			fmt.Println(string(bytes))
			context.XML(200, xmlContent)
		case "edbReturnTradeGet":
			edbResp := mock.EdbProxy(context.Request, context.Writer)
			fmt.Println(string(edbResp))
			context.Data(200, "application/xml; charset=utf-8", edbResp)
		case "edbTradeAdd":
			edbResp := mock.EdbProxy(context.Request, context.Writer)
			fmt.Println(string(edbResp))
			context.Data(200, "application/xml; charset=utf-8", edbResp)
		case "edbBlankOrderAdd":
			edbResp := mock.EdbProxy(context.Request, context.Writer)
			fmt.Println(string(edbResp))
			context.Data(200, "application/xml; charset=utf-8", edbResp)
		}

	})
	r.Run("0.0.0.0:8181") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}