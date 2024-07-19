package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func initiateRoutes() *gin.Engine {

	route := gin.Default()
	route.LoadHTMLGlob("templates/**")

	route.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusAccepted, "index.html", nil)
	})

	//route.GET("/contacts", func(context *gin.Context) {
	//	//context.IndentedJSON(http.StatusOK, contacts)
	//})

	route.POST("/contacts", func(context *gin.Context){
		var newContact contact_string
		if err := context.BindJSON(&newContact); err != nil {
			fmt.Printf("%s", err)
		}
		//contacts = append(contacts, convert_contact(newContact)) //todo
	})
	return route
}
