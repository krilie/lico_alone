import React from "react";
import {GetUserToken,SetUserToken} from "../../../utils/LocalStorageUtil";

/**
 * 1.检查token 如果token不为空或null 则跳转主管理页面
 * 2.执行用户登录操作 操作成功设置token并跳转到主管理页面
 */
export default class LoginPage extends React.Component {

    componentDidMount() {
        const token = GetUserToken();
        if (token !== "")
            this.goToPage("/management/manage");
    }

    goToPage = path => {
        this.props.history.push(path);
    };

    setToken(token){
        SetUserToken(token)
    }

    render() {
        return (
            <div>
                login page
            </div>
        );
    }
}
