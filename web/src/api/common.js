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
export function getArticleSampleList(searchKey, then) {
    getQuery("/common/article/query_sample", {search_key: searchKey}).then((res) => {
        if (res.data.code !== 2000) {
            openNotification(res.data.message);
        }
        then(res.data.data);
    }).catch((error) => {
        openNotification(error.toString());
    });
}















