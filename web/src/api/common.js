import axios from "axios";
import openNotification from "../utils/MessageBoard";

const apiCommon = axios.create({
    baseURL: 'https://api.lizo.top'
})

const get = (url) => {
    console.log(url)
    return apiCommon({
        method: "get",
        url: `${url}`,
    });
};

// {"code":2000,"message":"successful","data":{"name":"1","link":"2","label":"3"}}
export const getIcpInfo = () => {
    get("api/common/icp_info").then((res) => {
        if (res.code !== 2000) {
            openNotification(res.message);
        }
        return res.data;
    }).catch((error) => {
        openNotification(error.toString());
        return {name: "unknown", link: "unknown", label: "unknown"};
    });
}

export const getVersion = ()=>{
    get("api/common/icp_info").then((res) => {
        if (res.code !== 2000) {
            openNotification(res.message);
        }
        return res.data;
    }).catch((error) => {
        openNotification(error.toString());
        return {name: "unknown", link: "unknown", label: "unknown"};
    });
}