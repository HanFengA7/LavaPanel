package service

import "github.com/gin-gonic/gin"

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/"+GetSystemConfig().AuthCode, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = router.Run(":" + GetSystemConfig().PanelPort)
}
