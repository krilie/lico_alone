import axios from "axios";
import {apiBaseUrl} from "./ApiBaseUrl";
import qs from 'qs'
import {message} from "antd";

// =====================================================================================================

// 非api 外层返回结构可能不统一
const apiCommon = axios.create({
    baseURL: apiBaseUrl
})

const commonGet = (url, query) => {
    console.log(url)
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
export const getIcpInfo = (then) => {
    commonGet("/api/common/icp_info").then((res) => {
        if (res.data.code !== 2000) {
            message.warning(res.data.message);
        }
        then(res.data.data);
    }).catch((error) => {
        message.error(error.toString());
    });
}

export const getVersion = (then) => {
    commonGet("/api/common/version").then((res) => {
        then(res.data)
    }).catch((error) => {
        message.error(error.toString());
    });
}

export const postVisited = (traceId, then) => {
    commonPostForm("/api/common/visited", {traceId: traceId}).then((res) => {
        then(res.data)
    }).catch((error) => {
        message.error(error.toString());
    });
}

// ===================================================================================================

// 获取文章列表sample
export function getArticleSampleList(pageNum, pageSize, searchKey, funcOk, funcFinally) {
    commonGet("/api/common/article/query_sample", {
        page_num: pageNum,
        page_size: pageSize,
        search_key: searchKey
    }).then((res) => {
        // http 200
        if (res.data.code !== 2000) {
            message.warning(res.data.message);
        } else {
            funcOk(res.data.data);
        }
    }).catch((error) => {
        // http !200
        message.error(error.toString());
    }).finally(() => {
        funcFinally()
    });
}

export function getArticleById(articleId, then) {
    commonGet("/api/common/article/get_article", {article_id: articleId}).then(res => {
        if (res.data.code !== 2000) {
            message.warning(res.data.message);
        } else {
            then(res.data.data);
        }
    }).catch(err => {
        message.error(err.toString());
    })
}

// =============================================================

export function GetCarouselPicData(then) {
    commonGet("/api/common/carousel/query").then(res => {
        if (res.data.code !== 2000) {
            message.warning(res.data.message);
        } else {
            then(res.data.data);
        }
    }).catch(err => {
        message.error(err.toString());
    })
}