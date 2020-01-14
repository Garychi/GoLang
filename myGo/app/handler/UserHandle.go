package handler

import (
	"Lv/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchAllUser 取得所有使用者資訊
// @Summary 取得所有使用者資訊
// @Description 取得所有使用者資訊
// @Tags 使用者
// @Success 200 {array} model.UserAccount "使用者資訊"
// @Router /api/allUser [get]
func FetchAllUser(c *gin.Context) {
	userList := []model.UserAccount{}
	userList = append(userList, model.UserAccount{ID: 1, Account: "admin", Type: 1, Password: "1234", Name: "Admin"})
	userList = append(userList, model.UserAccount{ID: 2, Account: "admin2", Type: 1, Password: "12345", Name: "Admin2"})
	//userList[0] = model.UserAccount{ID: 1, Account: "admin", Type: 1, Password: "1234", Name: "Admin"}
	//userList[1] = model.UserAccount{ID: 2, Account: "admin2", Type: 1, Password: "12345", Name: "Admin2"}

	c.JSON(http.StatusOK, gin.H{
		"Result": userList,
	})
}

// FetchUser 取得所有使用者資訊2
// @Summary 取得所有使用者資訊2
// @Description 取得所有使用者資訊2
// @Tags 使用者2
// @Success 200 {object} model.UserAccount "使用者資訊2"
// @Router /api/user [get]
func FetchUser(c *gin.Context) {
	userAccount := model.UserAccount{
		ID:       1,
		Type:     1,
		Account:  "user001",
		Password: "1234",
		Name:     "Alice",
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userAccount,
	})

}
