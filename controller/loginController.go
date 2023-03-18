package controller

import (
	"github.com/PangQiMing/InterviewItem/config"
	"github.com/PangQiMing/InterviewItem/dto"
	"github.com/PangQiMing/InterviewItem/entity"
	"github.com/PangQiMing/InterviewItem/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginHandler 登录账号
func LoginHandler(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	err := ctx.ShouldBind(&loginDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user entity.User
	config.DB.Where("account = ? && password = ?", loginDTO.Account, loginDTO.Password).Take(&user)
	if loginDTO.Account != user.Account && loginDTO.Password != user.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "账号密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(user.Account)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Token = token
	config.DB.Save(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   user.Token,
	})
}
