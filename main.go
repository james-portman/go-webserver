package main

import (
	"time"
	"fmt"
	// "io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"os/exec"
	"runtime"
)

func main() {
	go runWebserver()

	time.Sleep(1 * time.Second)
	if runtime.GOOS == "windows" {
		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:8080").Start()
	} else {
		// fmt.Println("Cannot auto open in browser yet for this OS")
	}

	for true {
		time.Sleep(1 * time.Second)
	}
}

func runWebserver() {

	gin.SetMode(gin.ReleaseMode)
	// gin.DefaultWriter = ioutil.Discard // to disable web hits output to console

	router := gin.Default()
	router.Use(cors.Default()) // allow all origins

	router.Use(static.Serve("/", static.LocalFile("web-static", false)))

	fmt.Println("Starting webserver at: http://localhost:8080")
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
