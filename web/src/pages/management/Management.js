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
            this.goToPage(`${this.props.match.path}/login`);
    }

    goToPage = path => {
        this.props.history.push(path);
    };

    render() {
        return (
            <div>
                <div className="management-view">网站管理页面</div>
                <Switch>
                    <Route exact path={`${this.props.match.path}/login`} component={LoginPage}/>
                    <Route path={`${this.props.match.path}/manage`} component={ManagePage}/>
                    <Redirect path={`${this.props.match.path}/`} to={{pathname: `${this.props.match.path}/manage`}}/>
                </Switch>
            </div>
        );
    }
}