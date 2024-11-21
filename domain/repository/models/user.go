package models

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

type User struct {
	ID       int       `json:"id,omitempty"`
	Username string    `json:"username" binding:"required"`
	Phone    string    `json:"phone" binding:"required,phone_format"`
	Birthday time.Time `json:"-"`
}

// 自定義 JSON 序列化
func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Alias
		Birthday string `json:"birthday"`
	}{
		Alias:    Alias(u),
		Birthday: u.Birthday.Format("2006-01-02"),
	})
}

// 用於創建/更新用戶的請求結構
type UserRequest struct {
	Username string `json:"username" binding:"required" example:"derek"`
	Phone    string `json:"phone" binding:"required,phone_format" example:"0987-123456"`
	Birthday string `json:"birthday" binding:"required,datetime_format" example:"2012-01-01"`
}

// 驗證手機號碼格式
func ValidatePhoneFormat(phone string) bool {
	pattern := `^0\d{3}-\d{6}$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}

// 驗證日期格式
func ValidateDateFormat(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, fmt.Errorf("日期格式必須為 YYYY-MM-DD")
	}
	return t, nil
}

// 全局用戶存儲
var UserStore = map[int]User{
	1: {
		ID:       1,
		Username: "jared",
		Phone:    "0912-123456",
		Birthday: time.Now().Truncate(24 * time.Hour),
	},
}

// 用於生成新的ID
var LastID = 1
