import React from "react";
import "./Management.less"
import {Redirect, Route, Switch} from "react-router-dom";
import LoginPage from "./login-page/LoginPage";
import ManagePage from "./manage-page/ManagePage";
import {GetUserToken} from "../../utils/LocalStorageUtil";

export default class Management extends React.Component {

    componentWillMount() {
        const token = GetUserToken();
        if (token === "")
            this.goToPage("/management/login");
    }

    goToPage = path => {
        window.location.href = path;
        location.reload();
    };

    render() {
        return (
            <div>
                <div>网站管理页面</div>
                <Switch>
                    <Route exact path="/management/login" component={LoginPage}/>
                    <Route exact={false} path="/management/manage" component={ManagePage}/>
                    <Redirect path="/management" to={{pathname: '/management/login'}}/>
                </Switch>
            </div>
        );
    }
}