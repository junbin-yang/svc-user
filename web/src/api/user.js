import service from '@/utils/request'
// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
  return service({
    url: '/api/login',
    method: 'post',
    data: data
  })
}

export const getEmailCaptcha = (data) => {
  return service({
    url: '/api/validateCodeByEmail',
    method: 'post',
    data: data
  })
}

export const retrievePassword = (data) => {
  return service({
    url: '/api/user/retrievePassword',
    method: 'post',
    data: data
  })
}

export const GegSystemInfo = (data) => {
  return service({
    url: '/api/admin/systemInfo',
    method: 'get',
    data: data
  })
}

export const getInvitationCode = (data) => {
  return service({
    url: '/api/admin/createInvitationCode',
    method: 'get',
    data: data
  })
}

export const registerUser = (data) => {
  return service({
    url: '/api/user/register',
    method: 'post',
    data: data
  })
}

export const captcha = (data) => {
  return service({
    url: '/api/validateCodeByBase64',
    method: 'get',
    data: data
  })
}

// @Summary 修改当前用户密码
// @Produce  application/json
// @Param data body {password:"string",newPassword:"string",validateCodeId:"string",validateCode:"string"}
// @Router /api/admin/changingPassword [post]
export const changePassword = (data) => {
  return service({
    url: '/api/admin/changingPassword',
    method: 'post',
    data: data
  })
}

export const getUserList = (data) => {
  return service({
    url: '/api/admin/account/list',
    method: 'post',
    data: data
  })
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
export const deleteUser = (data) => {
  return service({
    url: '/user/deleteUser',
    method: 'delete',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserInfo [put]
export const setUserInfo = (data) => {
  return service({
    url: '/user/setUserInfo',
    method: 'put',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setSelfInfo [put]
export const setSelfInfo = (data) => {
  return service({
    url: '/user/setSelfInfo',
    method: 'put',
    data: data
  })
}

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.setUserAuthorities true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthorities [post]
export const setUserAuthorities = (data) => {
  return service({
    url: '/user/setUserAuthorities',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/admin/getUserInfo [get]
export const getUserInfo = () => {
  return service({
    url: '/api/admin/getUserInfo',
    method: 'get'
  })
}

export const resetPassword = (data) => {
  return service({
    url: '/api/admin/retrievePassword',
    method: 'post',
    data: data
  })
}

export const updateUser = (data) => {
  return service({
    url: '/api/admin/updateUser',
    method: 'post',
    data: data
  })
}
