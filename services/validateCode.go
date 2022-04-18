package services

import (
	"context"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/junbin-yang/golib/captcha"
	"github.com/junbin-yang/see"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"svc-user/model"
	"svc-user/proto"
)

// @Summary 验证码
// @Description 获取图形验证码
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /validateCodeByBase64 [get]
func (this *Svr) ValidateCodeByBase64(c *see.Context) Response {
	id, b64s, val, e := captcha.GenerateDefaultCaptcha(true)
	if e != nil {
		return Failure(e.Error())
	}
	rediskey := "ValidateCode.Base64." + id + "." + val
	this.Redis.Set(rediskey, val, 600)
	return RespData(see.H{"key": id, "b64s": b64s})
}

func (this *Svr) ValidateCodeVerify(id, val string, by ...string) bool {
	rediskey := "ValidateCode.Base64." + id + "." + val

	if by != nil {
		switch by[0] {
		case "Email":
			rediskey = "ValidateCode.Email." + id + "." + val
		case "Mobile":
			rediskey = "ValidateCode.Mobile." + id + "." + val
		}
	}

	exists, _ := this.Redis.Exists(rediskey)
	return exists
}

func (this *Svr) CreateValidateCode(ctx context.Context, in *proto.NullRequest) (*proto.ValidateCodeReply, error) {
	id, b64s, val, e := captcha.GenerateDefaultCaptcha(true)
	if e != nil {
		return nil, grpc.Errorf(codes.DataLoss, e.Error())
	}
	rediskey := "ValidateCode.Base64." + id + "." + val
	this.Redis.Set(rediskey, val, 600)

	return &proto.ValidateCodeReply{CodeId: id, ImageData: b64s}, nil
}

func (this *Svr) VerifyValidateCode(ctx context.Context, in *proto.VerifyValidateCodeRequest) (*proto.Reply, error) {
	status := this.ValidateCodeVerify(in.CodeId, in.CodeValue)
	if status {
		return ProtoSucceed(), nil
	}
	return ProtoFailure(), nil
}

type SendEmailParam struct {
	Email          string `json:"email" validate:"required,email"`
	InvitationCode string `json:"invitationCode" validate:"-"`
}

// @Summary 验证码
// @Description 发送邮件验证码
// @Tags Public
// @Accept json
// @Produce json
// @Param Body body SendEmailParam true "收件人邮箱地址"
// @Success 200 {object} Response
// @Router /validateCodeByEmail [post]
func (this *Svr) ValidateCodeByEmail(c *see.Context) Response {
	var parameters SendEmailParam
	e := c.ShouldBindJSON(&parameters)
	if e != nil {
		return Failure(e.Error())
	}

	if parameters.InvitationCode == "" {
		if find := model.M.Where("email = ?", parameters.Email).First(&model.User{}).RowsAffected; find == 0 {
			return Failure("邮箱地址不存在")
		}
	} else {
		invitationCode := "InvitationCode.RegisterUser." + parameters.InvitationCode
		exists, _ := this.Redis.Exists(invitationCode)
		if !exists {
			return Failure("邀请码无效")
		}
	}

	id, _, val, e := captcha.GenerateDefaultCaptcha(true)
	if e != nil {
		return Failure(e.Error())
	}
	rediskey := "ValidateCode.Email." + id + "." + val
	this.Redis.Set(rediskey, val, 600)

	content := "你好！，你的邮箱为：" + parameters.Email + "。请回填验证码：" + val + "。验证码有效期10分钟，请尽快校验！"
	e = this.SendEmail(parameters.Email, "验证码校验", content)
	if e != nil {
		return Failure(e.Error())
	}

	return RespData(see.H{"key": id})
}

// @Summary 创建邀请码
// @Description 创建邀请码
// @Tags Management
// @Accept json
// @Produce json
// @Success 200 {object} Response "响应信息"
// @Router /admin/createInvitationCode [get]
func (this *Svr) CreateInvitationCode(c *see.Context) Response {
	str := EncryptMd5(fmt.Sprint(carbon.Now().TimestampWithNanosecond()))
	code := fmt.Sprintf("%x-%x-%x-%x-%x", str[0:4], str[4:6], str[6:8], str[8:10], str[10:14])
	rediskey := "InvitationCode.RegisterUser." + code
	this.Redis.Set(rediskey, code, 86400)
	return RespData(code)
}
