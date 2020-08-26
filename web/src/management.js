import React from 'react';
import 'antd/dist/antd.css';
import ReactDOM from 'react-dom';
import {Provider} from "react-redux";
import store from "./redux/RuduxIndex";
import {BrowserRouter, Route} from "react-router-dom";
import Management from "./pages/management/Management";

ReactDOM.render((
    <Provider store={store}>
        <BrowserRouter basename='/'>
            <Route path={`/management`} component={Management}/>
        </BrowserRouter>
    </Provider>
), document.getElementById('management'));
