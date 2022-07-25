package main

import (
	"fakerAPI/main/config"
	"fakerAPI/main/infrastructure"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	env := config.GetEnvironment()

	r.Use(func(context *gin.Context) {
		rapidAPISecret := context.GetHeader("X-RapidAPI-Proxy-Secret")
		if rapidAPISecret != env.RapidAPIKey {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	})
	r.Use(func(context *gin.Context) {
		user := context.GetHeader("X-Rapidapi-User")
		if user == "" {
			context.AbortWithStatus(http.StatusUnauthorized)
		} else {
			context.Set(config.UserContext, user)
		}
	})
	r.GET("/health", func(context *gin.Context) {
		user := context.Value(config.UserContext)
		if user == "" {
			user = " World"
		}
		context.String(http.StatusOK, fmt.Sprintf("Hello %s!", user))
		context.Status(http.StatusOK)
	})

	useCases := infrastructure.NewUseCases(env)
	version1 := r.Group("v1")
	infrastructure.BuildControllersV1(version1, useCases)

	err := r.Run(":" + env.Port)
	if err != nil {
		log.Panicln(err)
	}
}
