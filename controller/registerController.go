package controller

import (
	"github.com/PangQiMing/InterviewItem/config"
	"github.com/PangQiMing/InterviewItem/dto"
	"github.com/PangQiMing/InterviewItem/entity"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"net/http"
)

// RegisterHandler 注册账号
func RegisterHandler(ctx *gin.Context) {
	var registerDTO dto.Register
	err := ctx.ShouldBind(&registerDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user entity.User
	err = smapping.FillStruct(&user, smapping.MapFields(&registerDTO))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx := config.DB.Save(&user)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": tx.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}
