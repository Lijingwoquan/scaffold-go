import axios from "axios";
import { useRouter } from 'vue-router';
const service = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_API,
})
const router = useRouter()

// 添加请求拦截器
service.interceptors.request.use(function (config) {

  return config;
}, function (error) {
  // 对请求错误做些什么
  return Promise.reject(error);
});


// 添加响应拦截器
service.interceptors.response.use(function (response) {
  //对响应数据进行处理
  return response.data.data; //后续请求响应数据写起来更加优雅
});

export default service
