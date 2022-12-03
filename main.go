package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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
	sendRequest("5000", "TEST 5000 OK")
	sendRequest("8080", "TEST 8080 OK")
	wg.Wait()
}

func startServers() {
	log.Println("main starts")
	go StartGin("8080", "srv p 8080")
	time.Sleep(2 * time.Second)
	go StartGin("5000", "srv p 5000")
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
	router.POST("/test", func(c *gin.Context) {
		message = c.Request.PostFormValue("message")
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

func sendRequest(port string, message string) {
	response, err := http.PostForm(
		"http://localhost:"+port+"/test",
		url.Values{"message": {message}})

	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
