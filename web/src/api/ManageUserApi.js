import {postMultiForm} from "./ApiManage";

// 用户登录
export function userLogin({phone, password}) {
    return postMultiForm("/user/login", {phone, password});
}

// 用户注册
export function userRegister({phone,password,valid_code}) {
    return postMultiForm("/user/register",{phone,password,valid_code});
}