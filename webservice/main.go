package main

import (
	"webservice/controllers"
	"webservice/helpers"

	"github.com/gin-gonic/gin"
)
func main(){
	
	mongo_uri := helpers.GetEnvByName("MONGO_URI")
	db := helpers.InitDB(mongo_uri)
	
	router := gin.Default()
	controllers.InitRoutes(db, router)
	
	port := helpers.GetEnvByName("API_PORT")
	router.Run(port)
}
