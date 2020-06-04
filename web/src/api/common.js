import axios from "axios";
import openNotification from "../utils/MessageBoard";
import {baseUrl} from "./baseUrl";
import {getQuery} from "./api";
import qs from 'qs'

// =====================================================================================================

// 非api 外层返回结构可能不统一
const apiCommon = axios.create({
    baseURL: baseUrl
})

const commonGet = (url, query) => {
    console.log(url)
    return apiCommon({
        method: "get",
        url: query === undefined ? `${url}` : `${url}?${qs.stringify(query)}`,
    });
};

// ====================================================================================================

// {"code":2000,"message":"successful","data":{"name":"1","link":"2","label":"3"}}
export const getIcpInfo = (then) => {
    commonGet("api/common/icp_info").then((res) => {
        if (res.data.code !== 2000) {
            openNotification(res.data.message);
        }
        then(res.data.data);
    }).catch((error) => {
        openNotification(error.toString());
    });
}

export const getVersion = (then) => {
    commonGet("version").then((res) => {
        then(res.data)
    }).catch((error) => {
        openNotification(error.toString());
    });
}

// ===================================================================================================

// 获取文章列表sample
export function getArticleSampleList(pageNum, pageSize, searchKey, funcOk, funcFinally) {
    getQuery("/common/article/query_sample", {
        page_num: pageNum,
        page_size: pageSize,
        search_key: searchKey
    }).then((res) => {
        // http 200
        if (res.data.code !== 2000) {
            openNotification(res.data.message);
        } else {
            funcOk(res.data.data);
        }
    }).catch((error) => {
        // http !200
        openNotification(error.toString());
    }).finally(() => {
        funcFinally()
    });
}