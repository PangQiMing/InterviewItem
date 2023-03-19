package middleware

import (
	"github.com/PangQiMing/InterviewItem/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.VerificationToken(ctx)
	}
}
