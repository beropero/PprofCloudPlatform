import axios from 'axios'

// 创建axios实例
const req = axios.create({
  baseURL: '/api', 
  headers: { 
    'contentType': 'application/json',
  },
  timeout: 3000,
})

export default req;