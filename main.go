package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	// ConfigRuntime()
	// StartWorkers()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers start starsWorker by goroutine.
func StartWorkers() {
	go StatsWorker()
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(RateLimit, gin.Recovery())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", Index)
	router.GET("/room/:roomid", RoomGET)
	router.POST("/room-post/:roomid", RoomPOST)
	router.GET("/stream/:roomid", StreamRoom)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
