import axios from "axios";
import qs from 'qs'
import {ToastErr, ToastWarn} from "../tools/toast";

export const apiBaseUrl = 'https://lizo.top'
export const imageProxy = 'https://imageproxy.lizo.top' // 缩略图地址

// res.data is returned data from api
// with have filed [code.message.data]
export const checkResData = (res) => res.data.code === 2000 ? res.data.data : undefined;
export const getResData = (res) => checkResData(res) ? res.data.data : undefined;
export const checkIsNotFound = (res)=> res.data.code === 4004;
export const checkResDataWithToast = (res) => {
    if (res.data.code !== 2000) {
        ToastWarn(res.data.message);
        return undefined;
    } else {
        return res.data.data;
    }
}

// =====================================================================================================

// 非api 外层返回结构可能不统一
const apiCommon = axios.create({baseURL: apiBaseUrl})

// 返回后拦截
apiCommon.interceptors.response.use(
    data => {
        return data
    },
    err => {
        console.log(err.toString() + "," + err.response.toString())
        if (err.response !== undefined && err.response !== null) {
            if (err.response.status === 502 || err.response.status === 504) {
                ToastErr("网关错误");
            } else if (err.response.status === 401) {
                ToastErr("认证信息无效");
            } else if (err.response.status === 500) {
                ToastErr("服务器错误");
            } else if (err.response.status === 404) {
                ToastErr("未找到内容");
            } else {
                ToastErr("未知错误");
            }
        } else {
            ToastErr(err.toString())
        }
        return Promise.reject(err);
    }
);

const commonGet = (url, query) => {
    return apiCommon({
        method: "get",
        url: query === undefined ? `${url}` : `${url}?${qs.stringify(query)}`,
    });
};


// @RequestBody请求
export const commonPostJson = (url, params) => {
    return apiCommon({
        method: "post",
        url: `${url}`,
        data: params,
        headers: {
            "Content-Type": "application/json",
            charset: "utf-8"
        }
    });
};

// @RequestParam请求
export const commonPostQuery = (url, params) => {
    return apiCommon({
        params: params,
        method: "post",
        url: `${url}`,
    });
};


// @RequestParam请求
export const commonPostForm = (url, params) => {
    return apiCommon({
        method: "post",
        url: `${url}`,
        data: qs.stringify({...params}),
        headers: {"Content-Type": "application/x-www-form-urlencoded"}
    });
};

export const commonGetQuery = (url, query) => {
    console.log(query)
    return apiCommon({
        method: "get",
        url: query === undefined ? `${url}` : `${url}?${qs.stringify(query)}`,
    });
};

export const commonPostMultiForm = (url, params) => {
    let param = new window.FormData();
    for (let i in params) {
        param.append(i, params[i]);
    }
    return apiCommon({
        method: 'post',
        url: `${url}`,
        data: param,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
};

// ====================================================================================================
// {"code":2000,"message":"successful","data":{"name":"1","link":"2","label":"3"}}
export const getIcpInfo = () => commonGet("/api/common/icp_info");
export const getVersion = () => commonGet("/api/common/version");
export const postVisited = (traceId) => commonPostForm("/api/common/visited", {traceId: traceId});
// ===================================================================================================
// 获取文章列表sample
export function getArticleSampleList(pageNum, pageSize, searchKey) {
    return commonGet("/api/common/article/query_sample", {
        page_num: pageNum,
        page_size: pageSize,
        search_key: searchKey
    });
}

export function getArticleById(articleId) {
    return commonGet("/api/common/article/get_article", {article_id: articleId});
}

// =============================================================
export const GetCarouselPicData = () => commonGet("/api/common/carousel/query")
