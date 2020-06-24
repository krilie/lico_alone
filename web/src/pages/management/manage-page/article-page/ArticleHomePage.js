import React, {Component} from 'react';
import {Route, Switch} from "react-router";
import ArticleListPage from "./ArticleListPage";
import ArticleEditPage from "./ArticleEditPage";
import {Redirect} from "react-router-dom";
import {GetUserToken} from "../../../../utils/LocalStorageUtil";

class ArticleHomePage extends Component {

    componentWillMount() {
        const token = GetUserToken();
        if (token === "")
            this.goToPage("/management/login");
    }

    render() {
        return (
            <div>
                <Switch>
                    <Route exact path="/management/manage/article/list" component={ArticleListPage}/>
                    <Route exact path="/management/manage/article/edit/:articleId" component={ArticleEditPage}/>
                    <Route exact path="/management/manage/article/create" component={ArticleEditPage}/>
                    <Redirect path="/management/manage/article" to={{pathname: '/management/manage/article/list'}}/>
                </Switch>
            </div>
        );
    }
}

export default ArticleHomePage;