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
	var wg sync.WaitGroup
	wg.Add(1)
	log.Println("main starts")
	go StartGin("3000", "srv1")
	log.Println("server are running")
	wg.Wait()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartGin starts gin web server with setting router.
func StartGin(port string, message string) {
	//gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	//router.Use(RateLimit, gin.Recovery())
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
