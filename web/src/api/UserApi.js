import "./api"
import {postMultiForm} from "./api";

// 用户登录
export function userLogin({name, password}) {
    return postMultiForm("/user/login", {name, password});
}

// 用户注册
export function userRegister({phone,password,valid_code}) {
    return postMultiForm("/user/register",{phone,password,valid_code});
}