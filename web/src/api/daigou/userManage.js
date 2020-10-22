import service from '@/utils/request'
const bathPath = 'daigou/userManage/'

export const deleteUserInfoByIds = async (data) => {
  return service({
    url: bathPath + 'deleteUserInfoByIds',
    method: 'delete',
    data
  })
}

export const operateUser = async (data) => {
  return service({
    url: bathPath + 'operateUser',
    method: 'put',
    data
  })
}

export const findUserInfoList = async (data) => {
  return service({
    url: bathPath + 'findUserInfoList',
    method: 'post',
    data
  })
}
