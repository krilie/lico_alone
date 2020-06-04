import axios from "axios";
import {message} from "antd";
import qs from 'qs'
import {GetUserToken, ClearToken} from "../utils/LocalStorageUtil";
import {baseUrl} from "./baseUrl";

// api请求组 外层返回结构终一
const apiRequest = axios.create({
    baseURL: baseUrl
})

let base = "/api";
// 请求拦截器
apiRequest.interceptors.request.use(
    config => {
        const token = GetUserToken();
        if (token !== '')
            config.headers.Authorization = token;
        return config;
    },
    err => {
        return Promise.reject(err);
    }
);


// 返回后拦截
apiRequest.interceptors.response.use(
    data => {
        // 请求成功
        if (data.data.code === 2000) {
            return data;
        }
        // token有误
        if (data.data.code === 4002) {
            message.warning(data.data.message)
            ClearToken()
            window.location.reload()
            return Promise.reject(data)
        }
        // 其它错误
        message.error(data.data.message)
        return Promise.reject(data)
    },
    err => {
        if (err.response.status === 504 || err.response.status === 404) {
            message.error("服务器被吃了⊙﹏⊙∥");
        } else if (err.response.status === 401) {
            message.error("登录信息失效⊙﹏⊙∥");
        } else if (err.response.status === 500) {
            message.error("服务器开小差了⊙﹏⊙∥");
        }
        return Promise.reject(err);
    }
);

// @RequestBody请求
export const postJson = (url, params) => {
    return apiRequest({
        method: "post",
        url: `${base}${url}`,
        data: params,
        headers: {
            "Content-Type": "application/json",
            charset: "utf-8"
        }
    });
};

// @RequestParam请求
export const postQuery = (url, params) => {
    return apiRequest({
        params: params,
        method: "post",
        url: `${base}${url}`,
    });
};


// @RequestParam请求
export const postForm = (url, params) => {
    return apiRequest({
        method: "post",
        url: `${base}${url}`,
        data: qs.stringify({...params}),
        headers: {"Content-Type": "application/x-www-form-urlencoded"}
    });
};

export const getQuery = (url, query) => {
    console.log(query)
    return apiRequest({
        method: "get",
        url: query === undefined ? `${base}${url}` : `${base}${url}?${qs.stringify(query)}`,
    });
};

export const postMultiForm = (url, params) => {
    let param = new window.FormData();
    for (let i in params) {
        param.append(i, params[i]);
    }
    return apiRequest({
        method: 'post',
        url: `${base}${url}`,
        data: param,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
};