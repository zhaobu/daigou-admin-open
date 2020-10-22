import { store } from '@/store/index'
import axios from 'axios' // 引入axios
import { Loading, Message } from 'element-ui'
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 99999
})
let acitveAxios = 0
let loadingInstance
let timer
const showLoading = () => {
  acitveAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (acitveAxios > 0) {
      loadingInstance = Loading.service({ fullscreen: true })
    }
  }, 400)
}

const closeLoading = () => {
  acitveAxios--
  if (acitveAxios <= 0) {
    clearTimeout(timer)
    loadingInstance && loadingInstance.close()
  }
}
//http request 拦截器
service.interceptors.request.use(
  (config) => {
    showLoading()
    const token = store.getters['user/token']
    const user = store.getters['user/userInfo']
    config.data = JSON.stringify(config.data)
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': token,
      'x-user-id': user.ID
    }
    if (process.env.NODE_ENV !== 'production') {
      console.log({
        type: 'request',
        baseurl: config.baseURL,
        relurl: config.url,
        method: config.method,
        data: config.data,
        params: config.params
      })
    }
    return config
  },
  (error) => {
    closeLoading()
    Message({
      showClose: true,
      message: error,
      type: 'error'
    })
    return Promise.reject(error)
  }
)

//http response 拦截器
service.interceptors.response.use(
  (response) => {
    if (process.env.NODE_ENV !== 'production') {
      console.dir({
        type: 'response',
        baseurl: GetUrlBase(response.request.responseURL),
        relurl: GetUrlRelative(response.request.responseURL),
        data: response.data
      })
    }
    closeLoading()
    if (response.data.code == 0 || response.headers.success === 'true') {
      return response.data
    } else {
      Message({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg),
        type: 'error'
      })
      if (response.data.data && response.data.data.reload) {
        store.commit('user/LoginOut')
      }
      return Promise.reject(response.data.msg)
    }
  },
  (error) => {
    closeLoading()
    Message({
      showClose: true,
      message: error,
      type: 'error'
    })
    return Promise.reject(error)
  }
)

function GetUrlBase(url) {
  const regex = new RegExp('/')
  const arrUrl = url.split(regex)
  const res = `${arrUrl[0]}//${arrUrl[2]}`
  return res
}

function GetUrlRelative(url) {
  const arrUrl = url.split('//')

  const start = arrUrl[1].indexOf('/')
  let relUrl = arrUrl[1].substring(start) // stop省略，截取从start开始到结尾的所有字符

  if (relUrl.indexOf('?') !== -1) {
    relUrl = relUrl.split('?')[0]
  }
  return relUrl
}

export default service
