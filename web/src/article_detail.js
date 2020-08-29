import React from 'react';
import 'antd/dist/antd.css';
import ReactDOM from 'react-dom';
import ArticleDetailPage from "./pages/article/ArticleDetailPage";
import {Provider} from "react-redux";
import store from "./redux/RuduxIndex";
import {BrowserRouter, Route} from "react-router-dom";

ReactDOM.render((
    <Provider store={store}>
        <BrowserRouter basename='/'>
            <Route path={`/article_detail`} component={ArticleDetailPage}/>
        </BrowserRouter>
    </Provider>
), document.getElementById('article_detail'));
