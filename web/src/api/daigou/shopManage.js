import service from '@/utils/request'
const bathPath = 'daigou/shopManage/'

// export const createShopInfo = async (data) => {
//   return service({
//     url: bathPath + 'createShopInfo',
//     method: 'post',
//     data
//   })
// }

// export const deleteShopInfo = async (data) => {
//   return service({
//     url: bathPath + 'deleteShopInfo',
//     method: 'delete',
//     data
//   })
// }

export const deleteShopInfoByIds = async (data) => {
  return service({
    url: bathPath + 'deleteShopInfoByIds',
    method: 'delete',
    data
  })
}

// export const updateShopInfo = async (data) => {
//   return service({
//     url: bathPath + 'updateShopInfo',
//     method: 'put',
//     data
//   })
// }

export const operateShop = async (data) => {
  return service({
    url: bathPath + 'operateShop',
    method: 'put',
    data
  })
}

export const findShopInfoList = async (data) => {
  return service({
    url: bathPath + 'findShopInfoList',
    method: 'post',
    data
  })
}
