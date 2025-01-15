import request from '@/utils/request'

// 模板设计图片上传
export function uploadImage(data) {
  return request({
    url: '/tool/design/image',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data',
    },
    data: data
  })
}


