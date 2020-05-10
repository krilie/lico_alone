import React from "react";
import "./Management.less"
import {Redirect, Route, Switch} from "react-router-dom";
import Home from "../home/Home";
import Share from "../share/Share";
import Photos from "../photos/Photos";
import Article from "../article/Article";
import LoginPage from "./login-page/LoginPage";
import ManagePage from "./manage-page/ManagePage";

export default class Management extends React.Component {

    componentWillMount() {
        const {pathname} = this.props.location;
        this.props.history.push(pathname);
    }

    goToPage = path => {
        this.props.history.push(path);
    };

    render() {
        return (
            <div>
                <Switch>
                    <Route exact path="/management/login" component={LoginPage}/>
                    <Route exact path="/management/manage" component={ManagePage}/>
                    <Redirect path="/management/login" to={{pathname: '/home'}}/>
                </Switch>
            </div>
        );
    }
}