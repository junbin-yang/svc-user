package services

import ()

func (this *Svr) Bind() {
	// 初始化
	this.See.GET("/api/admin/initialized", Wrapper(this.Initialized))
	this.See.POST("/api/admin/createFirstUser", Wrapper(this.CreateFirstUser))
	// 公共接口
	this.See.GET("/api/validateCodeByBase64", Wrapper(this.ValidateCodeByBase64))
	this.See.POST("/api/validateCodeByEmail", Wrapper(this.ValidateCodeByEmail))
	this.See.POST("/api/login", Wrapper(this.Login))
	this.See.POST("/api/user/register", Wrapper(this.RegisterUser))
	this.See.POST("/api/user/retrievePassword", Wrapper(this.RetrievePassword))

	// 管理员接口
	this.See.GET("/api/admin/getMenu", JWTAuth(), Wrapper(this.GetAdminMenus))
	this.See.GET("/api/admin/getUserInfo", JWTAuth(), Wrapper(this.AdminGetUserInfo))
	this.See.POST("/api/admin/changingPassword", JWTAuth(), Wrapper(this.ChangingPasswordBySelf))
	this.See.POST("/api/admin/app/list", JWTAuth(), Wrapper(this.GetAppList))
	this.See.POST("/api/admin/app/getAppById", JWTAuth(), Wrapper(this.AdminGetAppById))
	this.See.POST("/api/admin/app/updateApp", JWTAuth(), Wrapper(this.AdminUpdateApp))
	this.See.POST("/api/admin/app/register", JWTAuth(), Wrapper(this.RegisterApp))
	this.See.POST("/api/admin/account/list", JWTAuth(), Wrapper(this.AdminGetUserList))
	this.See.POST("/api/admin/updateUser", JWTAuth(), Wrapper(this.AdminUpdateUser))
	this.See.POST("/api/admin/role/list", JWTAuth(), Wrapper(this.AdminGetRoleList))
	this.See.POST("/api/admin/role/updateRole", JWTAuth(), Wrapper(this.AdminUpdateRole))
	this.See.POST("/api/admin/role/createRole", JWTAuth(), Wrapper(this.AdminCreateRole))
	this.See.POST("/api/admin/role/bindUser", JWTAuth(), Wrapper(this.AdminBindUser))
	this.See.GET("/api/admin/createInvitationCode", JWTAuth(), Wrapper(this.CreateInvitationCode))
	this.See.POST("/api/admin/retrievePassword", JWTAuth(), Wrapper(this.ChangingPasswordByEmail))
	this.See.GET("/api/admin/systemInfo", JWTAuth(), Wrapper(this.GetAdminSystemInfo))

	this.See.POST("/api/admin/changingUserStatus", JWTAuth(), Wrapper(this.ChangingStatusByAdmin))
}
