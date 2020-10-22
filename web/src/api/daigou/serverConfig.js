import service from '@/utils/request'

const bathPath = 'daigou/serverConfig/'

// 获取redis配置列表
export async function getConnectList() {
  return service({
    url: bathPath + 'getConnectList',
    method: 'get'
  })
}

// 修改redis连接
export async function changeConnect(data) {
  return service({
    url: bathPath + 'changeConnect',
    method: 'post',
    data: data
  })
}

// // 获取redis配置列表
// export async function getRedisList() {
//   return service({
//     url: bathPath + 'getRedisList',
//     method: 'get'
//   })
// }

// // 修改redis连接
// export async function changeRedis(data) {
//   return service({
//     url: bathPath + 'changeRedis',
//     method: 'post',
//     data: data
//   })
// }

// // 获取mysql配置列表
// export async function getMysqlList() {
//   return service({
//     url: bathPath + 'getMysqlList',
//     method: 'get'
//   })
// }

// // 修改mysql连接
// export async function changeMysql(data) {
//   return service({
//     url: bathPath + 'changeMysql',
//     method: 'post',
//     data: data
//   })
// }
