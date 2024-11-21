package api

import (
	"mygosimplerestfulapi/api/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 設置路由
func (s *Server) setupRouter() {
	// Swagger 文檔
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API 路由組
	api := s.router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("", user.GetUsers)
			users.GET("/:id", user.GetUser)
			users.POST("", user.CreateUser)
			users.PUT("/:id", user.UpdateUser)
			users.DELETE("/:id", user.DeleteUser)
		}
	}
}
