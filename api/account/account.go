package account

import (
	"database/sql"
	"go-server/api/account/dto"
	"go-server/api/account/vo"
	"go-server/enum"
	"go-server/global"
	"go-server/model"
	"go-server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAccount interface {
	Register(ctx *gin.Context)                     // 用户注册
	Login(ctx *gin.Context)                        // 用户登录
	DeleteAccountById(ctx *gin.Context)            // 根据id修改账号
	ModifyPasswordById(ctx *gin.Context)           // 根据id修改账号密码
	UpdateStatusById(ctx *gin.Context)             // 根据id修改状态
	UpdateCurrentAccountPassword(ctx *gin.Context) // 修改当前账号密码
	GetAccountById(ctx *gin.Context)               // 根据id获取账号信息
	GetAccountPage(ctx *gin.Context)               // 分页获取账号数据
	UpdateCurrentAccountInfo(ctx *gin.Context)
}

type Account struct {
	db *gorm.DB
}

// @Summary 注册
// @Produce json

func (a Account) Register(ctx *gin.Context) {
	var createAccountDto dto.CreateAccountDto
	if err := ctx.ShouldBindJSON(&createAccountDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(ctx, message)
		return
	}
	// 1.判断两次密码是否一致
	if createAccountDto.Password != createAccountDto.ConfirmPassword {
		utils.Fail(ctx, "两次密码不一致")
		return
	}
	// 2.对密码加密
	password, err := utils.MakePassword(createAccountDto.Password)
	if err != nil {
		global.Logger.Error("密码加密失败" + err.Error())
		utils.Fail(ctx, "创建账号失败")
		return
	}
	// 3.创建账号信息
	if result := a.db.Create(&model.AccountEntity{
		Phone:    createAccountDto.Phone,
		Password: password,
		Status:   sql.NullInt64{Valid: true, Int64: enum.Normal},
	}).Error; result != nil {
		global.Logger.Error("创建账号失败" + result.Error())
		utils.Fail(ctx, "创建账号失败")
	} else {
		utils.Success(ctx, "创建成功")
	}
}

func (a Account) Login(ctx *gin.Context) {
	var accountDto dto.AccountDto
	if err := ctx.ShouldBindJSON(&accountDto); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(ctx, message)
		return
	}
	// 1.根据账号名去查询密码信息
	var accountEntity model.AccountEntity
	if result := a.db.Where("phone=?", accountDto.Phone).Select([]string{"password", "id", "phone", "status"}).First(&accountEntity); result.RowsAffected == 0 {
		global.Logger.Error("根据用户名查询数据失败" + result.Error.Error())
		utils.Fail(ctx, "账号或密码错误")
		return
	}
	if accountEntity.Status.Int64 == enum.Forbid {
		utils.Fail(ctx, "当前账号不允许登录,请联系管理员")
		return
	}
	// 2.判断密码是否正确
	isOk, err := utils.CheckPassword(accountEntity.Password, accountDto.Password)
	if err != nil {
		global.Logger.Error("校验密码错误" + err.Error())
		utils.Fail(ctx, "账号或密码错误")
		return
	}
	if !isOk {
		utils.Fail(ctx, "账号或密码错误")
		return
	}
	// 3.生产token返回给前端
	hmacUser := utils.HmacUser{
		Id:    int(accountEntity.Id),
		Phone: accountEntity.Phone,
	}
	if token, err := utils.GenerateToken(hmacUser); err == nil {
		utils.Success(ctx, gin.H{
			"id":        accountEntity.Id,
			"nickname":  accountEntity.Nickname,
			"token":     token,
			"birthday":  accountEntity.Birthday,
			"coin":      accountEntity.Coin,
			"city":      accountEntity.City,
			"country":   accountEntity.Country,
			"province":  accountEntity.Province,
			"star":      accountEntity.Star,
			"status":    accountEntity.Status,
			"vpi_level": accountEntity.Vip_Level,
			"phone":     accountEntity.Phone,
			"sex":       accountEntity.Sex,
		})
	} else {
		global.Logger.Error("生成token失败")
		utils.Fail(ctx, "账号或密码错误")
	}
}

func (a Account) DeleteAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	if result := a.db.Where("id=?", idInt).Delete(&model.AccountEntity{}).Error; result != nil {
		global.Logger.Error("根据id删除账号失败" + result.Error())
		utils.Fail(ctx, "删除失败")
	} else {
		utils.Success(ctx, "删除成功")
	}
}

func (a Account) ModifyPasswordById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	var modifyAccountPassword dto.ModifyAccountPassword
	if err := ctx.ShouldBindJSON(&modifyAccountPassword); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(ctx, message)
		return
	}
	// 1.判断两次密码是否一致
	if modifyAccountPassword.Password != modifyAccountPassword.ConfirmPassword {
		utils.Fail(ctx, "两次密码不一致")
		return
	}
	// 2.对密码加密
	password, err := utils.MakePassword(modifyAccountPassword.Password)
	if err != nil {
		global.Logger.Error("密码加密失败" + err.Error())
		utils.Fail(ctx, "创建账号失败")
		return
	}
	if result := a.db.Where("id=?", idInt).Updates(&model.AccountEntity{
		Password: password,
	}).Error; result != nil {
		global.Logger.Error("修改密码失败" + result.Error())
		utils.Fail(ctx, "修改密码失败")
	} else {
		utils.Success(ctx, "修改密码成功")
	}
}

func (a Account) UpdateStatusById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	// 1.根据id查询之前的状态
	var accountEntity model.AccountEntity
	if result := a.db.Where("id=?", idInt).Select([]string{"status"}).First(&accountEntity).Error; result != nil {
		global.Logger.Error("根据id查询账号数据失败" + result.Error())
		utils.Fail(ctx, "修改失败")
		return
	}
	status := 0
	if accountEntity.Status.Int64 == enum.Forbid {
		status = enum.Normal
	} else {
		status = enum.Forbid
	}
	if result := a.db.Where("id=?", idInt).Updates(&model.AccountEntity{
		Status: sql.NullInt64{Valid: true, Int64: int64(status)},
	}).Error; result != nil {
		global.Logger.Error("根据id修改状态失败" + result.Error())
		utils.Fail(ctx, "更新失败")
	} else {
		utils.Success(ctx, "更新成功")
	}
}

// 更新用户信息
func (a Account) UpdateCurrentAccountInfo(ctx *gin.Context) {
	accountId := ctx.GetInt("accountId")
	nickname := ctx.GetString("nickname")
	nation := ctx.GetString("nation")
	province := ctx.GetString("province")
	city := ctx.GetString("city")
	country := ctx.GetString("country")
	birthday := ctx.GetString("birthday")
	var editAccountInfo dto.EditAccountInfo
	if err := ctx.ShouldBindJSON(&editAccountInfo); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(ctx, message)
		return
	}
	if result := a.db.Where("id=?", accountId).Updates(&model.AccountEntity{
		Nickname: nickname,
		Nation:   nation,
		Province: province,
		City:     city,
		Birthday: birthday,
		Country:  country,
	}).Error; result != nil {
		global.Logger.Error("更新失败" + result.Error())
		utils.Fail(ctx, "更新失败")
	} else {
		utils.Success(ctx, "修改密码成功")
	}
}

func (a Account) UpdateCurrentAccountPassword(ctx *gin.Context) {
	accountId := ctx.GetInt("accountId")
	var modifyAccountPassword dto.ModifyAccountPassword
	if err := ctx.ShouldBindJSON(&modifyAccountPassword); err != nil {
		message := utils.ShowErrorMessage(err)
		utils.Fail(ctx, message)
		return
	}
	// 1.判断两次密码是否一致
	if modifyAccountPassword.Password != modifyAccountPassword.ConfirmPassword {
		utils.Fail(ctx, "两次密码不一致")
		return
	}
	// 2.对密码加密
	password, err := utils.MakePassword(modifyAccountPassword.Password)
	if err != nil {
		global.Logger.Error("密码加密失败" + err.Error())
		utils.Fail(ctx, "创建账号失败")
		return
	}
	if result := a.db.Where("id=?", accountId).Updates(&model.AccountEntity{
		Password: password,
	}).Error; result != nil {
		global.Logger.Error("修改密码失败" + result.Error())
		utils.Fail(ctx, "修改密码失败")
	} else {
		utils.Success(ctx, "修改密码成功")
	}
}

func (a Account) GetAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	var accountVo vo.AccountVo
	if result := a.db.Model(&model.AccountEntity{}).Where("id=?", idInt).
		Select([]string{"id", "phone", "status", "created_at", "updated_at"}).
		First(&accountVo).Error; result != nil {
		global.Logger.Error("根据id查询账号信息失败" + result.Error())
		utils.Fail(ctx, "根据id查询账号信息失败")
	} else {
		utils.Success(ctx, accountVo)
	}
}

func (a Account) GetAccountPage(ctx *gin.Context) {
	phone := ctx.DefaultQuery("phone", "")
	tx := a.db
	if phone != "" {
		tx = tx.Where("phone like ?", "%"+phone+"%")
	}
	var accountList []vo.AccountVo
	if result := tx.Model(&model.AccountEntity{}).Scopes(utils.Paginate(ctx.Request)).Find(&accountList).Error; result != nil {
		global.Logger.Error("查询列表失败" + result.Error())
	}
	var total int64
	if result := tx.Model(&model.AccountEntity{}).Count(&total).Error; result != nil {
		global.Logger.Error("查询条数失败" + result.Error())
	} else {
		utils.Success(ctx, utils.PageVo{
			Data:  accountList,
			Total: total,
		})
	}
}

func NewAccount(db *gorm.DB) IAccount {
	db.AutoMigrate(&model.AccountEntity{})
	return Account{
		db: db,
	}
}
