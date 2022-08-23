import axios from "axios";

var instance = axios.create({
  baseURL: import.meta.env.VITE_API_URL, // 获取当前环境的域名配置
  timeout: 6000, // 设置超时时间1分钟
  header: {
    "Content-Type": "application/json;charset=UTF-8",
  }, // 基础的请求头
});

// 请求中间件
instance.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token");
    token && (config.headers.Authorization = token); // 如果localstorage存在token则将token直接写入headers
    if (config.method === "POST") {
      config.data = JSON.stringify(config.data);
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// 返回结果中间件
instance.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    // TODO: 进行错误预先处理
    return Promise.reject(error);
  }
);

export default instance;
