package restserver

import (
	"fmt"
	"net/http"

	"github.com/abdullahjankhan/go_sample/models"
	"github.com/abdullahjankhan/go_sample/service"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	ValidateJWTToken() gin.HandlerFunc
}

type middleware struct {
	jwtService service.JWTService
}

func NewMiddleware(jwtService service.JWTService) Middleware {
	return &middleware{
		jwtService: jwtService,
	}
}

func (m *middleware) ValidateJWTToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Query("token")

		if len(token) <= 0 {
			token = ctx.GetHeader("token")
		}

		if len(token) == 0 {
			var err error
			token, err = ctx.Cookie("token")
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusOK, models.NewStandardResponse(false, models.INVALID_TOKEN, "token not founded", nil))
				return

			}
		}

		uid, err := m.jwtService.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, models.NewStandardResponse(false, models.INVALID_TOKEN, err.Error(), nil))
			return
		}

		if ctx == nil {
			ctx.AbortWithStatusJSON(http.StatusOK, models.NewStandardResponse(false, models.INVALID_TOKEN, err.Error(), nil))
			return
		}

		if len(uid) <= 0 {
			ctx.AbortWithStatusJSON(http.StatusOK, models.NewStandardResponse(false, models.INVALID_TOKEN, err.Error(), nil))
			return
		}

		ctx.Set(models.KeyCurrentUUID, uid)
		ctx.Next()
	}
}

func AttachBody[BodyType any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body BodyType

		err := ctx.BindJSON(body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, models.NewStandardResponse(false, models.INVALID_REQUEST_BODY, fmt.Sprintf(models.INVALID_REQUEST_BODY_MSG, err), nil))
		}
		ctx.Set(models.RequestKey, body)
		ctx.Next()
	}
}
