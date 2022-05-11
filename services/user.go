package services

import (
	"context"
	"fmt"
	"github.com/junbin-yang/golib/json"
	"github.com/junbin-yang/golib/logger"
	"github.com/junbin-yang/see"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"math/rand"
	"regexp"
	"svc-user/jwt"
	"svc-user/model"
	"svc-user/proto"
	"svc-user/utils/encoding"
	"time"
)

// @Summary 系统初始化
// @Description 判断是否已初始化系统
// @Tags Initialized
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /admin/initialized [get]
func (this *Svr) Initialized(c *see.Context) Response {
	if find := model.M.Last(&model.User{}).RowsAffected; find > 0 {
		return RespData(true)
	}
	return RespData(false)
}

type FirstUserParam struct {
	Password       string `json:"password" validate:"required"`       // 密码
	Account        string `json:"account" validate:"required"`        // 账号
	NickName       string `json:"nickName" validate:"-"`              // 昵称
	ValidateCodeId string `json:"validateCodeId" validate:"required"` // 验证码Id
	ValidateCode   string `json:"validateCode" validate:"required"`   // 验证码
}

// @Summary 初始化管理员
// @Description 创建第一个管理员用户
// @Tags Initialized
// @Accept json
// @Produce json
// @Param Body body FirstUserParam true "初始化参数"
// @Success 200 {object} Response "响应信息"
// @Router /admin/createFirstUser [post]
func (this *Svr) CreateFirstUser(c *see.Context) Response {
	var parameters FirstUserParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	if !this.ValidateCodeVerify(parameters.ValidateCodeId, parameters.ValidateCode) {
		return Failure("验证码错误")
	}

	if find := model.M.Last(&model.User{}).RowsAffected; find > 0 {
		return Failure("系统已完成初始化")
	}

	if parameters.NickName == "" {
		parameters.NickName = parameters.Account
	}

	u := model.User{
		Password: parameters.Password,
		Account:  parameters.Account,
		NickName: parameters.NickName,
		Admin:    true,
	}
	e = u.Create()
	if e != nil {
		return Failure(e.Error())
	}
	return Success("初始化成功")
}

type ChangPassParam struct {
	Password       string `json:"password" validate:"required"`
	NewPassword    string `json:"newPassword" validate:"required"`
	ValidateCodeId string `json:"validateCodeId" validate:"required"` // 验证码Id
	ValidateCode   string `json:"validateCode" validate:"required"`   // 验证码
}

// @Summary 修改密码
// @Description 通过旧密码设置新密码
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body ChangPassParam true "修改密码参数"
// @Success 200 {object} Response "响应信息"
// @Router /admin/changingPassword [post]
func (this *Svr) ChangingPasswordBySelf(c *see.Context) Response {
	userid := c.MustGet("userId").(uint64)
	var parameters ChangPassParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	if !this.ValidateCodeVerify(parameters.ValidateCodeId, parameters.ValidateCode) {
		return Failure("验证码错误")
	}

	if find := model.M.Where("id = ? AND password = ?", userid, parameters.Password).First(&model.User{}).RowsAffected; find > 0 {
		result := model.M.Model(model.User{}).Where("id = ?", userid).Updates(map[string]interface{}{"password": parameters.NewPassword})
		if result.Error != nil {
			return Failure(result.Error.Error())
		}
	} else {
		return Failure("旧密码错误")
	}

	return Success("修改成功")
}

// @Summary 重置用户密码
// @Description 重置用户密码，新密码发送邮件
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body SendEmailParam true "修改密码参数"
// @Success 200 {object} Response "响应信息"
// @Router /admin/retrievePassword [post]
func (this *Svr) ChangingPasswordByEmail(c *see.Context) Response {
	var parameters SendEmailParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	var newPass string = GetRandStr(12)
	result := model.M.Model(model.User{}).Where("email = ?", parameters.Email).Updates(map[string]interface{}{"password": EncryptMd5(newPass)})
	if result.Error != nil {
		return Failure(result.Error.Error())
	}

	content := "你好！ 你的新密码为：" + newPass + "，请重新登录！"
	this.SendEmail(parameters.Email, "密码找回", content)

	return RespData(newPass)
}

type RetrievePasswordParam struct {
	Email          string `json:"email" validate:"required,email"`
	ValidateCodeId string `json:"validateCodeId" validate:"required"` // 验证码Id
	ValidateCode   string `json:"validateCode" validate:"required"`   // 验证码
	Password       string `json:"password" validate:"required"`       // 密码
}

// @Summary 找回用户密码
// @Description 找回用户密码
// @Tags Public
// @Accept json
// @Produce json
// @Param Body body RetrievePasswordParam true "找回密码参数"
// @Success 200 {object} Response "响应信息"
// @Router /user/retrievePassword [post]
func (this *Svr) RetrievePassword(c *see.Context) Response {
	var parameters RetrievePasswordParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	if !this.ValidateCodeVerify(parameters.ValidateCodeId, parameters.ValidateCode, "Email") {
		return Failure("验证码错误")
	}

	passErr := JudgePassword(parameters.Password)
	if passErr != nil {
		return Failure(passErr.Error())
	}

	result := model.M.Model(model.User{}).Where("email = ?", parameters.Email).Updates(map[string]interface{}{"password": encoding.EncryptMd5(parameters.Password)})
	if result.Error != nil {
		return Failure(result.Error.Error())
	}

	return Success("修改成功")
}

func JudgePassword(ps string) error {
	e := fmt.Errorf("密码长度要求为6-30位包含大小写字母+数字+特殊符号")
	if len(ps) < 6 || len(ps) > 30 {
		return e
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, _ := regexp.MatchString(num, ps); !b {
		return e
	}
	if b, _ := regexp.MatchString(a_z, ps); !b {
		return e
	}
	if b, _ := regexp.MatchString(A_Z, ps); !b {
		return e
	}
	if b, _ := regexp.MatchString(symbol, ps); !b {
		return e
	}
	return nil
}

func GetRandStr(length int) string {
	baseStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#@^&*"
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, length)
	l := len(baseStr)
	for i := 0; i < length; i++ {
		bytes[i] = baseStr[r.Intn(l)]
	}
	return string(bytes)
}

type RegisterParam struct {
	InvitationCode string `json:"invitationCode" validate:"required"` // 邀请码
	ValidateCodeId string `json:"validateCodeId" validate:"required"` // 验证码Id
	ValidateCode   string `json:"validateCode" validate:"required"`   // 验证码
	Password       string `json:"password" validate:"required"`       // 密码
	Account        string `json:"account" validate:"required"`        // 账号
	NickName       string `json:"nickName" validate:"required"`       // 昵称
	Mobile         string `json:"mobile" validate:"required"`         // 手机号
	Email          string `json:"email" validate:"required"`          // 邮箱
	Info           struct {
		RealName string `json:"realName" validate:"-"` // 详情
		Age      uint   `json:"age" validate:"-"`      // 年龄
		Sex      uint   `json:"sex" validate:"-"`      // 性别
		Nation   string `json:"nation" validate:"-"`   // 民族
		Address  string `json:"address" validate:"-"`  // 地址
		IdCard   string `json:"idCard" validate:"-"`   // 身份证
	} `json:"info" validate:"-"` // 详情
}

func FindUser(email, mobile, account string) error {
	if ckEmail := model.M.Where("email = ?", email).First(&model.User{}).RowsAffected; ckEmail > 0 {
		return fmt.Errorf("邮箱地址已有关联用户")
	}
	if ckMobile := model.M.Where("mobile = ?", mobile).First(&model.User{}).RowsAffected; ckMobile > 0 {
		return fmt.Errorf("手机号码已有关联用户")
	}
	if ckAccount := model.M.Where("account = ?", account).First(&model.User{}).RowsAffected; ckAccount > 0 {
		return fmt.Errorf("账号已存在")
	}
	return nil
}

// @Summary 注册用户
// @Description 注册新用户
// @Tags Public
// @Accept json
// @Produce json
// @Param Body body RegisterParam true "注册参数"
// @Success 200 {object} Response
// @Router /user/register [post]
func (this *Svr) RegisterUser(c *see.Context) Response {
	var parameters RegisterParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	if !this.ValidateCodeVerify(parameters.ValidateCodeId, parameters.ValidateCode, "Email") {
		return Failure("验证码错误")
	}

	rediskey := "InvitationCode.RegisterUser." + parameters.InvitationCode
	exists, _ := this.Redis.Exists(rediskey)
	if !exists {
		return Failure("邀请码无效")
	}

	e = FindUser(parameters.Email, parameters.Mobile, parameters.Account)
	if e != nil {
		return Failure(e.Error())
	}

	u := model.User{
		Password: encoding.EncryptMd5(parameters.Password),
		Account:  parameters.Account,
		NickName: parameters.NickName,
		Mobile:   parameters.Mobile,
		Email:    parameters.Email,
		Info: model.Info{
			RealName: parameters.Info.RealName,
			Age:      parameters.Info.Age,
			Sex:      parameters.Info.Sex,
			Nation:   parameters.Info.Nation,
			Address:  parameters.Info.Address,
			IdCard:   parameters.Info.IdCard,
		},
	}
	e = u.Create()
	if e != nil {
		return Failure(e.Error())
	}

	if exists {
		this.Redis.Delete(rediskey)
	}

	return Success("注册成功")
}

func (this *Svr) CreateUser(ctx context.Context, in *proto.CreateUserRequest) (*proto.CreateUserReply, error) {
	e := FindUser(in.Email, in.Mobile, in.Account)
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}

	u := model.User{
		Password: encoding.EncryptMd5(in.Password),
		Account:  in.Account,
		NickName: in.NickName,
		Mobile:   in.Mobile,
		Email:    in.Email,
		Info: model.Info{
			RealName: in.UserInfo.RealName,
			Age:      uint(in.UserInfo.Age),
			Sex:      uint(in.UserInfo.Sex),
			Nation:   in.UserInfo.Nation,
			Address:  in.UserInfo.Address,
			IdCard:   in.UserInfo.IdCard,
		},
	}
	e = u.Create()
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}

	return &proto.CreateUserReply{UserId: uint32(u.ID)}, nil
}

func (this *Svr) UpdateUserInfo(ctx context.Context, in *proto.UpdateUserInfoRequest) (*proto.Reply, error) {
	result := model.M.Model(model.Info{}).Where("user_id = ?", in.Id).Updates(map[string]interface{}{
		"real_name": in.UserInfo.RealName,
		"age":       in.UserInfo.Age,
		"sex":       in.UserInfo.Sex,
		"nation":    in.UserInfo.Nation,
		"address":   in.UserInfo.Address,
		"id_card":   in.UserInfo.IdCard,
	})
	if result.Error != nil {
		return nil, grpc.Errorf(codes.DataLoss, result.Error.Error())
	}
	return ProtoSucceed(), nil
}

func (this *Svr) ChangingUserStatus(ctx context.Context, in *proto.ChangingUserStatusRequest) (*proto.Reply, error) {
	result := model.M.Model(model.User{}).Where("id = ?", in.Id).Updates(map[string]interface{}{"status": in.Status})
	if result.Error != nil {
		return nil, grpc.Errorf(codes.DataLoss, result.Error.Error())
	}
	return ProtoSucceed(), nil
}

type ChangingStatusParam struct {
	Id     string `json:"id" validate:"required"`
	Status bool   `json:"status" validate:"required"`
}

// @Summary 启/停用账号
// @Description 切换账号状态
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body ChangingStatusParam true "修改用户状态参数"
// @Success 200 {object} Response "响应信息"
// @Router /admin/changingUserStatus [post]
func (this *Svr) ChangingStatusByAdmin(c *see.Context) Response {
	var parameters ChangingStatusParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	result := model.M.Model(model.User{}).Where("id = ?", parameters.Id).Updates(map[string]interface{}{"status": parameters.Status})
	if result.Error != nil {
		return Failure(result.Error.Error())
	}

	return Success("修改成功")
}

func (this *Svr) GetUserList(ctx context.Context, in *proto.PaginateRequest) (*proto.GetUserListReply, error) {
	logger.Info(in)
	u := []model.User{}
	data, e := model.Paginate(&u, int(in.Page), int(in.Limit), model.ConversionSearchType(in.Search), "account != 'root'", "Info")
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}

	str, _ := json.ObjectToJson(data.Docs)
	docs := []*proto.GetUserListReply_UserListData{}
	json.JsonToObject(str, &docs)

	return &proto.GetUserListReply{
		Paginate: &proto.PaginateReply{
			Page:  uint32(data.Page),
			Pages: uint32(data.Pages),
			Total: uint32(data.Total),
			Limit: uint32(data.Limit),
		},
		Docs: docs,
	}, nil
}

type PaginateParam struct {
	Page   uint32              `json:"page" validate:"required"`
	Limit  uint32              `json:"limit" validate:"required"`
	Search map[string][]string `json:"search" validate:"-"`
}

// @Summary 获取账号列表
// @Description 管理员获取账号列表
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body PaginateParam true "分页参数"
// @Success 200 {object} Response "响应信息"
// @Router /admin/account/list [post]
func (this *Svr) AdminGetUserList(c *see.Context) Response {
	var parameters PaginateParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	u := []model.User{}
	data, e := model.Paginate(&u, int(parameters.Page), int(parameters.Limit), model.ConversionSearchType2(parameters.Search), "", "Info")
	if e != nil {
		return Failure(e.Error())
	}

	return RespData(data)
}

func GetUserPermissions(appkey string, userId uint) map[string]uint32 {
	permissions := map[string]uint32{}
	// 查询用户有多少个角色
	roleUserMap := []model.RoleUserMap{}
	model.M.Where("app_key = ? AND user_id = ?", appkey, userId).Find(&roleUserMap)
	roleIds := []uint{}
	for _, role := range roleUserMap {
		if find := model.M.Where("status = ? AND id = ?", 1, role.RoleID).First(&model.Role{}).RowsAffected; find > 0 {
			roleIds = append(roleIds, role.RoleID)
		}
	}

	rows, err := model.M.Table("role_authorities").Where("role_id IN ?", roleIds).Rows()
	if err != nil {
		logger.Error(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var auth model.RoleAuthority
		model.M.ScanRows(rows, &auth)
		permissions[auth.Code] = permissions[auth.Code] | auth.ActionValue
	}
	return permissions
}

// @Summary 用户详情
// @Description 获取当前用户详情
// @Tags Management
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /admin/getUserInfo [get]
func (this *Svr) AdminGetUserInfo(c *see.Context) Response {
	userid := c.MustGet("userId").(uint64)
	userInfo := new(model.User)
	userInfo.FirstByUserId(userid)

	if userInfo.Status == false {
		return Failure("账号已禁用")
	}
	return RespData(see.H{"userInfo": UserInfoResp{
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
	}})
}

type UpdateUserParam struct {
	Id       uint64        `json:"id" validate:"required"`
	Email    string        `json:"email" validate:"required"`
	Status   *bool         `json:"status" validate:"required"`
	Mobile   string        `json:"mobile" validate:"required"`
	NickName string        `json:"nickname" validate:"required"`
	Info     UserInfoParam `json:"info" validate:"-"`
}

type UserInfoParam struct {
	RealName string `json:"realname" validate:"-"`
	Age      *uint  `json:"age" validate:"-"`
	Sex      *uint  `json:"sex" validate:"-"`
	Nation   string `json:"nation" validate:"-"`
	IdCard   string `json:"idcard" validate:"-"`
	Address  string `json:"address" validate:"-"`
}

// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags Management
// @Accept json
// @Produce json
// @Param Body body UpdateUserParam true "参数"
// @Success 200 {object} Response
// @Router /admin/updateUser [post]
func (this *Svr) AdminUpdateUser(c *see.Context) Response {
	var parameters UpdateUserParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	userInfo := new(model.User)
	userInfo.Email = parameters.Email
	userInfo.Status = *parameters.Status
	userInfo.Mobile = parameters.Mobile
	userInfo.NickName = parameters.NickName
	userInfo.Info.RealName = parameters.Info.RealName
	userInfo.Info.Age = *parameters.Info.Age
	userInfo.Info.Sex = *parameters.Info.Sex
	userInfo.Info.Nation = parameters.Info.Nation
	userInfo.Info.IdCard = parameters.Info.IdCard
	userInfo.Info.Address = parameters.Info.Address

	err := userInfo.UpdateInfo(parameters.Id)
	if err != nil {
		return Failure(err.Error())
	}
	return Success("修改成功")
}

func (this *Svr) UserPassVerify(ctx context.Context, in *proto.UserPassVerifyRequest) (*proto.Token, error) {
	userInfo := new(model.User)
	password := encoding.EncryptMd5(in.Password)
	switch in.Type {
	case "account":
		userInfo.FirstByAccount(in.Value, password)
	case "mobile":
		userInfo.FirstByMobile(in.Value, password)
	case "email":
		userInfo.FirstByEmail(in.Value, password)
	default:
	}

	if userInfo.ID == 0 {
		return nil, grpc.Errorf(codes.DataLoss, "账号密码错误")
	}

	if userInfo.Status == false {
		return nil, grpc.Errorf(codes.DataLoss, "账号已禁用")
	}

	token, err := jwt.Create(uint64(userInfo.ID), userInfo.Admin, userInfo.Account, userInfo.NickName, "")
	if err != nil {
		return nil, grpc.Errorf(codes.DataLoss, err.Error())
	}
	return &proto.Token{Token: token}, nil
}
