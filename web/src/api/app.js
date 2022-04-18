import service from '@/utils/request'

export const getAppList = (data) => {
  return service({
    url: '/api/admin/app/list',
    method: 'post',
    data
  })
}

export const getAppById = (data) => {
  return service({
    url: '/api/admin/app/getAppById',
    method: 'post',
    data
  })
}

export const updateApp = (data) => {
  return service({
    url: '/api/admin/app/updateApp',
    method: 'post',
    data
  })
}

export const createApp = (data) => {
  return service({
    url: '/api/admin/app/register',
    method: 'post',
    data
  })
}

export const getAppRoleList = (data) => {
  return service({
    url: '/api/admin/role/list',
    method: 'post',
    data
  })
}

export const updateRole = (data) => {
  return service({
    url: '/api/admin/role/updateRole',
    method: 'post',
    data
  })
}

export const createRole = (data) => {
  return service({
    url: '/api/admin/role/createRole',
    method: 'post',
    data
  })
}

export const bindRoleUser = (data) => {
  return service({
    url: '/api/admin/role/bindUser',
    method: 'post',
    data
  })
}
