package services

import (
	"github.com/junbin-yang/golib/logger"
	"github.com/junbin-yang/see"
	"svc-user/jwt"
	"svc-user/model"
	"time"
)

type LoginParam struct {
	Type           string `json:"loginType" validate:"required"`      // 登录类型，account|mobile|email
	Value          string `json:"loginValue" validate:"required"`     // 登录类型值
	Password       string `json:"password" validate:"required"`       // 密码
	ValidateCodeId string `json:"validateCodeId" validate:"required"` // 验证码Id
	ValidateCode   string `json:"validateCode" validate:"required"`   // 验证码
}

type LoginResp struct {
	Authorization string       `json:"authorization"` // Token
	UserInfo      UserInfoResp `json:"userInfo"`
}
type UserInfoResp struct {
	NickName    string                 `json:"nickname"`
	RealName    string                 `json:"realname"`
	Age         uint                   `json:"age"`
	Sex         uint                   `json:"sex"`
	Nation      string                 `json:"nation"`
	Address     string                 `json:"address"`
	IdCard      string                 `json:"idcard"`
	Mobile      string                 `json:"mobile"`
	Email       string                 `json:"email"`
	Authority   map[string]interface{} `json:"authority"`
	HeaderImg   string                 `json:"headerImg"`
	SideMode    string                 `json:"sideMode"` // 菜单主题
	ActiveColor string                 `json:"activeColor"`
	BaseColor   string                 `json:"baseColor"`
}

// @Summary 登录
// @Description 通过账号密码登录
// @Tags Public
// @Accept json
// @Produce json
// @Param Body body LoginParam true "登录参数"
// @Success 200 {object} Response{data=LoginResp} "token"
// @Router /login [post]
func (this *Svr) Login(c *see.Context) Response {
	var parameters LoginParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	if !this.ValidateCodeVerify(parameters.ValidateCodeId, parameters.ValidateCode) {
		return Failure("验证码错误")
	}

	userInfo := new(model.User)

	switch parameters.Type {
	case "account":
		userInfo.FirstByAccount(parameters.Value, parameters.Password)
	case "mobile":
		userInfo.FirstByMobile(parameters.Value, parameters.Password)
	case "email":
		userInfo.FirstByEmail(parameters.Value, parameters.Password)
	default:
	}

	if userInfo.ID == 0 {
		return Failure("账号密码错误")
	}

	if userInfo.Status == false {
		return Failure("账号已禁用")
	}

	token, err := jwt.Create(uint64(userInfo.ID), userInfo.Admin, userInfo.Account, userInfo.NickName, "basic")
	if err != nil {
		logger.Error(err.Error())
		return Failure("Create Token Fail")
	}

	userInfo.Update(map[string]interface{}{"logined_at": time.Now()})

	return RespData(
		LoginResp{
			token,
			UserInfoResp{
				userInfo.NickName,
				userInfo.Info.RealName,
				userInfo.Info.Age,
				userInfo.Info.Sex,
				userInfo.Info.Nation,
				userInfo.Info.Address,
				userInfo.Info.IdCard,
				userInfo.Mobile,
				userInfo.Email,
				map[string]interface{}{
					"defaultRouter": "dashboard",
				},
				"https://cdn.learnku.com/uploads/images/201805/11/1/ZnrA2VK0SN.png!/both/200x200",
				"dark", // 暗黑模式
				"#1890ff",
				"#fff",
			},
		},
	)
}
