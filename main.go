package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ConfigRuntime()
	var wg sync.WaitGroup
	wg.Add(1)
	startServers()
	time.Sleep(2 * time.Second)
	sendRequest("8080")
	wg.Wait()
}

func startServers() {
	log.Println("main starts")
	go StartServer2("8080", "srv p 8080")
	time.Sleep(2 * time.Second)
	go StartServer1("5000", "srv p 5000")
	log.Println("server are running")
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartGin starts gin web server with setting router.
func StartGin(port string, message string) {
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	//router.Use(RateLimit, gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, message)
	})

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }
	//log.Println("ENV PORT", os.Getenv("PORT"))
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func StartServer1(port string, message string) {
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, message)
	})

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func StartServer2(port string, message string) {
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, message)
	})
	router.POST("/test", func(c *gin.Context) {
		c.JSON(200, "TEST OK")
	})
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func sendRequest(port string) {
	response, err := http.Post(
		"http://localhost:"+port+"/test",
		"text/plain",
		bytes.NewBuffer([]byte("test")))
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
