package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	// ConfigRuntime()
	// StartWorkers()
	var wg sync.WaitGroup
	wg.Add(1)
	log.Println("main starts")
	go StartGin("3000", "srv1")
	go StartGin("5000", "srv2")
	log.Println("server are running")
	wg.Wait()
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
func StartGin(port string, message string) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(RateLimit, gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, message)
	})

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
