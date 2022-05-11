package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/junbin-yang/see"
	"svc-user/model"
)

type UpdateAppParam struct {
	Id        uint64 `json:"id" validate:"required"`
	AppSecret string `json:"appsecret" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Remark    string `json:"remark" validate:"-"`
	Status    *bool  `json:"status" validate:"required"`
}

// @Summary 更新应用
// @Description 更新应用
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body UpdateAppParam true "参数"
// @Success 200 {object} Response
// @Router /admin/app/updateApp [post]
func (this *Svr) AdminUpdateApp(c *see.Context) Response {
	var parameters UpdateAppParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	result := model.M.Model(model.App{}).Where("id = ?", parameters.Id).Updates(map[string]interface{}{"name": parameters.Name, "app_secret": parameters.AppSecret, "status": *parameters.Status, "remark": parameters.Remark})
	if result.Error != nil {
		return Failure(result.Error.Error())
	}
	return Success("修改成功")
}

type IdParam struct {
	Id uint64
}

// @Summary 应用详情
// @Description 获取应用详情
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body PaginateParam true "注册参数"
// @Success 200 {object} Response
// @Router /admin/app/getAppById [post]
func (this *Svr) AdminGetAppById(c *see.Context) Response {
	var parameters IdParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	appInfo := new(model.App)
	appInfo.FirstById(parameters.Id)

	return RespData(see.H{"app": appInfo})
}

// @Summary 获取应用列表
// @Description 获取应用列表
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body PaginateParam true "注册参数"
// @Success 200 {object} Response
// @Router /admin/app/list [post]
func (this *Svr) GetAppList(c *see.Context) Response {
	var parameters PaginateParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	u := []model.App{}
	data, e := model.Paginate(&u, int(parameters.Page), int(parameters.Limit), model.ConversionSearchType2(parameters.Search), "")
	if e != nil {
		return Failure(e.Error())
	}

	return RespData(data)
}

type RegisterAppParam struct {
	ValidateCodeId string `json:"validateCodeId" validate:"required"` // 验证码Id
	ValidateCode   string `json:"validateCode" validate:"required"`   // 验证码
	Name           string `json:"name" validate:"required"`
	Remark         string `json:"remark" validate:"-"`
}

// @Summary 注册应用
// @Description 注册新应用
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body RegisterAppParam true "注册参数"
// @Success 200 {object} Response
// @Router /admin/app/register [post]
func (this *Svr) RegisterApp(c *see.Context) Response {
	var parameters RegisterAppParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	if !this.ValidateCodeVerify(parameters.ValidateCodeId, parameters.ValidateCode) {
		return Failure("验证码错误")
	}

	AppKey := EncryptMd5By16Bit(fmt.Sprint(carbon.Now().TimestampWithNanosecond()))
	str := EncryptMd5(fmt.Sprint(carbon.Now().TimestampWithNanosecond()))
	AppSecret := fmt.Sprintf("%x-%x-%x-%x-%x", str[0:4], str[4:6], str[6:8], str[8:10], str[10:14])

	app := model.App{
		Name:      parameters.Name,
		Remark:    parameters.Remark,
		AppKey:    AppKey,
		AppSecret: AppSecret,
	}
	e = app.Create()
	if e != nil {
		return Failure(e.Error())
	}
	return Success("注册成功")
}

func EncryptMd5(input string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(input))
	cipherStr := md5Ctx.Sum(nil)
	encryptedData := hex.EncodeToString(cipherStr)
	return encryptedData
}

func EncryptMd5By16Bit(data string) string {
	return EncryptMd5(data)[8:24]
}

// @Summary 系统信息
// @Description 系统信息
// @Tags Management
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /admin/systemInfo [get]
func (this *Svr) GetAdminSystemInfo(c *see.Context) Response {
	var count, loginTaday int64
	model.M.Model(&model.User{}).Count(&count)
	model.M.Model(&model.User{}).Where("logined_at like ?", carbon.Now().ToDateString()+"%").Count(&loginTaday)
	return RespData(see.H{"userTotal": count, "loginTaday": loginTaday})
}
