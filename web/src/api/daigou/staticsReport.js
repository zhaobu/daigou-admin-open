import service from '@/utils/request'
const bathPath = 'daigou/staticsReport/'

export const newUser = async (params) => {
  return service({
    url: bathPath + 'newUser',
    method: 'get',
    params
  })
}

export const activeUser = async (params) => {
  return service({
    url: bathPath + 'activeUser',
    method: 'get',
    params
  })
}
