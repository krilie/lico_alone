import "./api"
import {postMultiForm} from "./api";

// 用户登录
export function userLogin({name, password}) {
    return postMultiForm("/user/login", {name, password});
}