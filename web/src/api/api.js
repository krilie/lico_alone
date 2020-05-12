import axios from "axios";
import qs from 'qs'
import {GetUserToken} from "../utils/LocalStorageUtil";
import openNotification from "../utils/MessageBoard"

// api请求组 外层返回结构终一
const apiRequest = axios.create({
    baseURL: 'https://api.lizo.top'
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
        if (data.data.code !== 2000) {
            openNotification(data.data.message)
            return Promise.reject(data)
        }
        return data;
    },
    err => {
        if (err.response.status === 504 || err.response.status === 404) {
            openNotification("服务器被吃了⊙﹏⊙∥");
        } else if (err.response.status === 401) {
            openNotification("登录信息失效⊙﹏⊙∥");
        } else if (err.response.status === 500) {
            openNotification("服务器开小差了⊙﹏⊙∥");
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
        url: `${base}${url}?${qs.stringify(query)}`,
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