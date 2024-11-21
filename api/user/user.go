package user

import (
	"mygosimplerestfulapi/domain/repository/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 獲取所有用戶
// @Description 獲取系統中所有用戶的列表
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users [get]
func GetUsers(c *gin.Context) {
	users := make([]models.User, 0)
	for _, user := range models.UserStore {
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}

// @Summary 獲取單個用戶
// @Description 根據ID獲取用戶信息
// @Tags users
// @Produce json
// @Param id path int true "用戶ID"
// @Success 200 {object} models.User
// @Failure 404 {object} models.ErrorResponse
// @Router /api/users/{id} [get]
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if user, exists := models.UserStore[id]; exists {
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "用戶不存在"})
}

// @Summary 創建用戶
// @Description 創建新用戶
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserRequest true "用戶信息"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /api/users [post]
func CreateUser(c *gin.Context) {
	var userReq models.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.Error(err)
		return
	}

	models.LastID++
	birthday, _ := models.ValidateDateFormat(userReq.Birthday)
	user := models.User{
		ID:       models.LastID,
		Username: userReq.Username,
		Phone:    userReq.Phone,
		Birthday: birthday,
	}

	models.UserStore[user.ID] = user
	c.JSON(http.StatusCreated, user)
}

// @Summary 更新用戶
// @Description 更新指定用戶的所有信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "用戶ID"
// @Param user body models.UserRequest true "用戶信息"
// @Success 200 {object} models.User
// @Failure 404 {object} models.ErrorResponse
// @Router /api/users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, exists := models.UserStore[id]; !exists {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Message: "用戶不存在",
			Code:    404,
		})
		return
	}

	var userReq models.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.Error(err)
		return
	}

	birthday, _ := models.ValidateDateFormat(userReq.Birthday)
	user := models.User{
		ID:       id,
		Username: userReq.Username,
		Phone:    userReq.Phone,
		Birthday: birthday,
	}

	models.UserStore[id] = user
	c.JSON(http.StatusOK, user)
}

// @Summary 刪除用戶
// @Description 刪除指定用戶
// @Tags users
// @Produce json
// @Param id path int true "用戶ID"
// @Success 200 {object} string
// @Failure 404 {object} models.ErrorResponse
// @Router /api/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, exists := models.UserStore[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "用戶不存在"})
		return
	}

	delete(models.UserStore, id)
	c.JSON(http.StatusOK, gin.H{"message": "用戶已刪除"})
}
