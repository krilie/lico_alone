import axios from "axios";
import openNotification from "../utils/MessageBoard";

// 非api 外层返回结构可能不统一
const apiCommon = axios.create({
    // baseURL: 'https://lizo.top'
    // baseURL: ''
})

const get = (url) => {
    console.log(url)
    return apiCommon({
        method: "get",
        url: `${url}`,
    });
};

// {"code":2000,"message":"successful","data":{"name":"1","link":"2","label":"3"}}
export const getIcpInfo = (then) => {
    get("api/common/icp_info").then((res) => {
        if (res.data.code !== 2000) {
            openNotification(res.data.message);
        }
        then(res.data.data);
    }).catch((error) => {
        openNotification(error.toString());
    });
}

export const getVersion = (then)=>{
    get("version").then((res) => {
        then(res.data)
    }).catch((error) => {
        openNotification(error.toString());
    });
}