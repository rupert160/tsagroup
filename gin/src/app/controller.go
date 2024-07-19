package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/go-pg/pg/v10"
)

func initiateRoutes() *gin.Engine {

	route := gin.Default()
	route.LoadHTMLGlob("templates/**")

	route.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusAccepted, "index.html", nil)
	})

	route.GET("/contacts", func(context *gin.Context) {
		db := pg.Connect(&pg.Options{
			User: "tsauser",
       			Password: "tsapass",
       			Database: "tsagroup",
       			Addr: "postgres:5432",

    		})
		// Select all users.
		var users []Contact
		err := db.Model(&users).Select()
		if err != nil {
		    panic(err)
		}
		context.IndentedJSON(http.StatusOK, users)
    		db.Close()
	})

	route.POST("/contacts", func(context *gin.Context){
		var newContact contact_string
		if err := context.BindJSON(&newContact); err != nil {
			fmt.Printf("%s", err)
		}
		c := convert_contact(newContact)
		contacts = append(contacts, c) //todo
		db := pg.Connect(&pg.Options{
			User: "tsauser",
       			Password: "tsapass",
       			Database: "tsagroup",
       			Addr: "postgres:5432",
    		})

                _, err := db.Model(&Contact{
                    Name:  c.Full_Name,
                    Email: c.Email,
                    Phone_Numbers: c.Phone_Numbers,
                }).Insert()
                if err != nil {
                    panic(err)
                }

    		db.Close()
	})
	return route
}
