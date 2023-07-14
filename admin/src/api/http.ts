import axios, { type Method } from 'axios'
import qs from 'qs'

import { ElMessage, ElLoading } from 'element-plus'
import router from '../router'
import { useUserStore } from '@/stores'

axios.defaults.withCredentials = false
axios.defaults.withCredentials = false

export const request = async function (
  url: string,
  method: Method,
  queryParams?: any,
  fromdata?: any
): Promise<any> {
  const response = await axios.request({
    url: url,
    method: method,
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
    },
    params: queryParams,
    data: fromdata,
    timeout: 30 * 1000,
    transformRequest: [
      (data: any) => {
        return qs.stringify(data)
      }
    ]
  })
  return response.data
}

export const requestBody = async function (
  url: string,
  method: Method,
  querys?: any,
  body?: any
): Promise<any> {
  let config
  if (body instanceof FormData) {
    config = {
      url: url,
      method: method,
      headers: {
        'Content-Type': 'multipart/form-data'
      },
      params: querys,
      data: body,
      timeout: 10 * 60 * 1000
    }
  } else {
    config = {
      url: url,
      method: method,
      headers: {
        'Content-Type': 'application/json;charset=UTF-8'
      },
      params: querys,
      data: body,
      timeout: 30 * 1000,
      transformRequest: [
        (data: any) => {
          return JSON.stringify(data)
        }
      ]
    }
  }
  const response = await axios.request(config)
  return response.data
}

let loading: any

function showLoading() {
  loading = ElLoading.service({
    lock: true,
    text: '加载中...',
    background: 'rgba(0, 0, 0, 0.7)'
  })
}

function hideLoading() {
  loading.close()
}

axios.interceptors.request.use(
  (value: any) => {
    showLoading()
    return value
  },
  (error: any) => {
    return Promise.reject(error)
  }
)

axios.interceptors.response.use(
  (value: any) => {
    // 处理业务层的错误
    hideLoading()
    if (!value.data || !value.data.status) {
      ElMessage.error('通信数据错误')
      return Promise.reject(value)
    }
    switch (value.data.status) {
      case 200:
        return value
      case 401: {
        let userStore = useUserStore()
        if (userStore.isLoggedIn) {
          userStore.logout()
          ElMessage.error('登录失效，请重新登录')
          router.push('/login')
        }
        break
      }
      default:
        ElMessage.error('[' + value.data.status + '] ' + value.data.message)
        break
    }
    return Promise.reject(value)
  },
  (error: any) => {
    // 处理 HTTP 层的错误
    hideLoading()
    if (error.response) {
      ElMessage.error('服务器通信失败:' + error.response.status)
    } else {
      ElMessage.error('连接服务器失败')
    }
    return Promise.reject(error)
  }
)
