package interceptor

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-disk/common"
	"go-disk/common/rpcinterface/authinterface"
	"go-disk/services/upload/rpc"
	"log"
	"net/http"
)

func AuthorizeInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")

		if token == "" {
			log.Printf("request param error")
			ctx.Abort()
			ctx.JSON(http.StatusBadRequest,
				common.NewServiceResp(common.RespCodeBindReParamError, nil))
			return
		}
		resp, err := rpc.GetAuthCli().Authentication(context.TODO(), &authinterface.AuthenticationReq{
			AccessToken: token,
		})
		if err != nil || resp.Code != int64(common.RespCodeSuccess.Code) {
			log.Printf("token validate error")
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized,
				common.NewServiceResp(common.RespCodeUnauthorizedError, nil))
			return
		}
		ctx.Next()
	}
}


