import axios from "axios"

export default defineNuxtPlugin((NuxtApp) => {

    // 允许跨域请求时携带 Cookie
    axios.defaults.withCredentials = true;
    // TODO:确认路由
    axios.defaults.baseURL = 'http://localhost:8888'

    return {
        provide: { 
            axios: axios
        },
    }
})