import React, {Component} from 'react';
import {Route, Switch} from "react-router";
import ArticleListPage from "./ArticleListPage";
import ArticleEditPage from "./ArticleEditPage";
import {Redirect} from "react-router-dom";

class ArticleHomePage extends Component {

    render() {
        return (
            <div>
                <Switch>
                    <Route exact path={`${this.props.match.path}/list`} component={ArticleListPage}/>
                    <Route exact path={`${this.props.match.path}/edit/:articleId`} component={ArticleEditPage}/>
                    <Route exact path={`${this.props.match.path}/create`} component={ArticleEditPage}/>
                    <Redirect path={`${this.props.match.path}/`} to={{pathname: `${this.props.match.path}/list`}}/>
                </Switch>
            </div>
        );
    }
}

export default ArticleHomePage;