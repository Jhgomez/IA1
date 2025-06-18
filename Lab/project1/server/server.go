package main

import (
    "github.com/gin-gonic/gin"

	"unimatchserver/services"
)

func main() {
    router := gin.Default()

	knowledgeServiceInstance := knowledgeservice.GetKnowledgeService()

    router.GET("/getFacts", knowledgeServiceInstance.GetFacts)
    router.POST("/addFact", knowledgeServiceInstance.AddFact)
	router.POST("/deleteFact", knowledgeServiceInstance.DeleteFact)
	router.POST("/updateFact", knowledgeServiceInstance.UpdateFact)

    router.Run("localhost:8000")
}