import service from '@/utils/request'

export function getList(data) {
  return service({
    url: '/table/getList',
    method: 'post',
    data
  })
}

export function doEdit(data) {
  return service({
    url: '/table/doEdit',
    method: 'post',
    data
  })
}

export function doDelete(data) {
  return service({
    url: '/table/doDelete',
    method: 'post',
    data
  })
}
