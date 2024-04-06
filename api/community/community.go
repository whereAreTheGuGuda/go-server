package community

import (
	"database/sql"
	"go-server/api/community/dto"
	"go-server/api/community/vo"
	"go-server/enum"
	"go-server/global"
	"go-server/model"
	"go-server/utils"
	"strconv"

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
	// query := ctx.DefaultQuery("q", "")
	var total int64
	if result := a.db.Model(&model.CommunityEntity{}).Count(&total).Error; result != nil {
		global.Logger.Error("查询条数失败" + result.Error())
		utils.Fail(ctx, "查询失败")
	} else {
		var communityList []vo.CommunityListVo
		if result := a.db.Model(&model.CommunityEntity{}).Scopes(utils.Paginate(ctx.Request)).Find(&communityList).Error; result != nil {
			global.Logger.Error("查询列表失败" + result.Error())
			utils.Fail(ctx, "查询失败")
		} else {
			utils.Success(ctx, utils.PageVo{
				Data:  communityList,
				Total: total,
			})
		}
	}
}

func (a Community) Send(ctx *gin.Context) {
	accountId := ctx.GetInt("accountId")
	var communityDto dto.CommunityDto
	if err := ctx.ShouldBindJSON(&communityDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(ctx, message)
		return
	}

	if result := a.db.Create(&model.CommunityEntity{
		AccountId: accountId,
		Context:   communityDto.Context,
		Status:    sql.NullInt64{Valid: true, Int64: enum.Normal},
	}).Error; result != nil {
		global.Logger.Error("创建帖子失败" + result.Error())
		utils.Fail(ctx, "创建帖子失败")
	} else {
		utils.Success(ctx, "创建帖子成功")
	}
}

func (a Community) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	if result := a.db.Where("id=?", idInt).Delete(&model.CommunityEntity{}).Error; result != nil {
		global.Logger.Error("根据id删除帖子失败" + result.Error())
		utils.Fail(ctx, "删除失败")
	} else {
		utils.Success(ctx, "删除成功")
	}
}

func NewCommunity(db *gorm.DB) ICommunity {
	db.AutoMigrate(&model.CommunityEntity{})
	return Community{
		db: db,
	}
}
