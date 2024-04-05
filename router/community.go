package router

import (
	"go-server/api/community"
	"go-server/global"
	"go-server/middleware"

	"github.com/gin-gonic/gin"
)

func InitCommunityRouter(Router *gin.RouterGroup) {
	communityRouter := Router.Group("community")
	newCommunity := community.NewCommunity(global.DB)
	communityRouter.GET("list", newCommunity.CommunityList)                          // 帖子列表
	communityRouter.POST("send", middleware.AuthMiddleWare(), newCommunity.Send)     // 发表帖子
	communityRouter.DELETE("/:id", middleware.AuthMiddleWare(), newCommunity.Delete) // 根据id删除
}
