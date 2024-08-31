package routes

import (
    "github.com/parkrealgood/gotification/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

    router.POST("/topics", handlers.CreateTopic)
		router.GET("/topics", handlers.GetTopics)
		router.GET("/topics/:id", handlers.GetTopic)
    router.POST("/topics/:id/subscribe", handlers.SubscribeTopic)
    router.POST("/topics/:id/publish", handlers.PublishTopic)
}