package main

import (
	"github.com/aasumitro/goms/web"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const ServiceAddress = "localhost:3000"

func main() {
	appEngine := gin.Default()
	appEngine.StaticFS("/",
		http.FS(web.Resource))
	log.Fatal(appEngine.Run(ServiceAddress))
}
