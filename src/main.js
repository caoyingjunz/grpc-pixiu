import {
    createApp
} from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import router from './router/index.js'
import {
    ElMessage
} from 'element-plus'

// 全局引入css
import './style/apps.css'

const app = createApp(App)
app.use(ElementPlus, {
    locale: zhCn
})
app.use(router)

import axios from 'axios'
axios.defaults.baseURL = 'http:/127.0.0.1:8080'

// http拦截: 在 axios 请求发送之前给每个接口携带token去后端校验
axios.interceptors.request.use(config => {
    config.headers.Authorization = window.sessionStorage.getItem("token") // 服务端从请求头中获取 Authorization 作为 token
    return config
})
// http拦截: 在 axios 请求发送之后, 最返回值进行处理
axios.interceptors.response.use(
    response => {
        return response
    },
    error => {
        if (error.response) {
            let status = error.response.status
            switch (status) {
                case 401:
                    ElMessage({
                        message: error.response.data,
                        type: 'error',
                    })

                    window.sessionStorage.clear()
                    // 跳转到登陆界面
                    router.push("/")

                    break;
            }
        }
        return Promise.reject(error.response) // 返回错误信息
    }
)
app.config.globalProperties.$http = axios

app.mount('#app')