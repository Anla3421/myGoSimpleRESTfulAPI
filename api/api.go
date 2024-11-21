package api

import (
	"mygosimplerestfulapi/domain/repository/models"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	router *gin.Engine
}

// 創建新的 API 服務器
func NewServer() *Server {
	server := &Server{
		router: gin.Default(),
	}

	// 註冊自定義驗證器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone_format", validatePhoneFormat)
		v.RegisterValidation("datetime_format", validateDateFormat)
	}

	// 設置路由
	server.setupRouter()

	return server
}

// 啟動服務器
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

// 驗證手機號碼格式
func validatePhoneFormat(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	pattern := `^0\d{3}-\d{6}$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}

// 驗證日期格式
func validateDateFormat(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	_, err := models.ValidateDateFormat(date)
	return err == nil
}
