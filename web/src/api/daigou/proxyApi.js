import { respCode } from '@/utils/errorCode'
import service from '@/utils/request'

export const proxyRespCode = {
  ERROR: -1, //请求错误
  SUCCESS: 0, //请求成功
  FAIL: 1, //请求失败,一般指业务方面
  UNAUTHORIZED: 2, //鉴权失败
  EXPIRED: 3, //token失效(只在访问登录接口时使用)
  UNAUTHOPERATE: 4 //没有操作权限
}

const bathPath = 'daigou/proxyApi/'

// admin使用post请求代理访问server
async function proxyPost({ proxyURL, proxyData }) {
  const { code, msg, data } = await service({
    url: bathPath + 'proxyPost',
    method: 'post',
    data: {
      url: proxyURL,
      data: JSON.stringify(proxyData)
    }
  })
  if (code != respCode.SUCCESS) {
    this.$message.error(`proxyPost代理请求${proxyURL}错误,${msg}`)
    return
  }
  return data
}

export async function proxyGetShopCode(data) {
  return proxyPost({
    proxyURL: 'admin/getShopCode',
    proxyData: data
  })
}

export async function proxyUseShopCode(data) {
  return proxyPost({
    proxyURL: 'admin/useShopCode',
    proxyData: data
  })
}
