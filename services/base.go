package services

import (
	"crypto/tls"
	"fmt"
	"github.com/junbin-yang/golib/captcha"
	"github.com/junbin-yang/golib/redisx"
	"github.com/junbin-yang/see"
	seeSwagger "github.com/junbin-yang/see-swagger"
	"github.com/soheilhy/cmux"
	swaggerFiles "github.com/swaggo/files"
	"google.golang.org/grpc"
	"gopkg.in/gomail.v2"
	"net"
	"net/http"
	"os"
	"os/signal"
	"svc-user/config"
	"svc-user/constant"
	"svc-user/jwt"
	"svc-user/proto"
	"syscall"
	"time"
)

var App *Svr

func init() {
	App = &Svr{
		Name: constant.AppName,
		Conf: config.Parse(),
	}

	if App.Conf.Release {
		see.SetMode(see.ReleaseMode)
	}

	if App.Conf.Redis.Cluster {
		App.Redis = &redisx.Cluster{Host: App.Conf.Redis.Addr, Pass: App.Conf.Redis.Pass}
	} else {
		App.Redis = &redisx.Alone{Host: App.Conf.Redis.Addr[0], Pass: App.Conf.Redis.Pass}
	}
	App.Redis.Connect()

	App.Email = gomail.NewDialer(App.Conf.Smtp.Host, App.Conf.Smtp.Port, App.Conf.Smtp.User, App.Conf.Smtp.Pass)
	App.Email.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

type Svr struct {
	Name  string
	Conf  *config.Config
	See   *see.Engine
	Redis redisx.Rediser
	Email *gomail.Dialer
}

func (this *Svr) Run() {
	(&captcha.Options{}).New()

	l, e := net.Listen("tcp", this.Conf.Listen)
	if e != nil {
		panic(e)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	this.See = see.Enable("access", this.Conf.Logger.Dir, this.Conf.Logger.Rotate, this.Conf.Logger.KeepDays)
	if !this.Conf.Release {
		this.See.GET("/", func(c *see.Context) {
			c.Redirect(http.StatusMovedPermanently, "/dist/#/")
		})
		this.See.GET("/docs/swagger/*any", seeSwagger.WrapHandler(swaggerFiles.Handler, seeSwagger.URL("/docs/swagger.json")))
		this.See.StaticFile("/docs/swagger.json", "./docs/swagger.json")
		this.See.StaticFile("/docs/proto/", "./docs/index.html")
		this.See.Static("/dist", "./web/dist")
	}
	this.Bind()

	var opts []grpc.ServerOption
	// 注册拦截器
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	s := grpc.NewServer(opts...)
	proto.RegisterUserServer(s, this)

	go s.Serve(grpcL)
	go this.See.RunListener(httpL)
	go m.Serve()
	fmt.Println("Start Server ...")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	fmt.Println("Shutdown Server ...")
}

func ProtoSucceed() *proto.Reply {
	return &proto.Reply{Status: true}
}

func ProtoFailure() *proto.Reply {
	return &proto.Reply{Status: false}
}

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func RespData(data interface{}) Response {
	return Response{http.StatusOK, "success", data}
}

func Failure(msg interface{}) Response {
	return Response{http.StatusBadRequest, msg, map[string]interface{}{}}
}

func Success(msg interface{}) Response {
	return Response{http.StatusOK, msg, map[string]interface{}{}}
}

type WrapperFuncType func(c *see.Context) Response

func Wrapper(handle WrapperFuncType) see.HandlerFunc {
	return func(c *see.Context) {
		var resp Response = handle(c)
		c.PureJSON(http.StatusOK, resp)
	}
}

func JWTAuth() func(c *see.Context) {
	return func(c *see.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, see.H{
				"code":    -2,
				"message": "Authorization not found",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, see.H{
				"code":    -2,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		if (token.ExpiresAt - int64(time.Now().Unix())) < 1800 {
			newtoken, _ := token.Refresh()
			c.SetHeader("new-token", newtoken)
		}

		if !token.Admin {
			c.JSON(http.StatusOK, see.H{
				"code":    -2,
				"message": "非管理员账号",
			})
			c.Abort()
			return
		}

		// 将信息保存到请求的上下文c上
		c.Set("userId", token.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("account")来获取当前请求的用户
	}
}

func (this *Svr) SendEmail(receiver, title, content string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", fmt.Sprintf("%s<%s>", this.Name, this.Conf.Smtp.User))
	msg.SetHeader("To", receiver)
	//msg.SetAddressHeader("Cc", "dan@example.com", "Dan")
	msg.SetHeader("Subject", title)
	msg.SetBody("text/html", content)
	//msg.Attach("/home/Alex/lolcat.jpg")

	return this.Email.DialAndSend(msg)
}
