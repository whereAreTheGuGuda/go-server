package community

import (
	"go-server/model"
	"go-server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICommunity interface {
	CommunityList(ctx *gin.Context) // 帖子列表
	Send(ctx *gin.Context)          // 发送帖子
	Delete(ctx *gin.Context)        // 删除帖子
}

type Community struct {
	db *gorm.DB
}

func (a Community) CommunityList(ctx *gin.Context) {
	utils.Success(ctx, "请求成功")
}

func (a Community) Send(ctx *gin.Context) {

}

func (a Community) Delete(ctx *gin.Context) {

}

func NewCommunity(db *gorm.DB) ICommunity {
	db.AutoMigrate(&model.CommunityEntity{})
	return Community{
		db: db,
	}
}
