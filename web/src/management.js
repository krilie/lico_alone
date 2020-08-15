import React from 'react';
import ReactDOM from 'react-dom';
import ManagePage from "./pages/management/manage-page/ManagePage";
import {Provider} from "react-redux";
import store from "./redux/RuduxIndex";
import {BrowserRouter, Route} from "react-router-dom";

ReactDOM.render((
    <Provider store={store}>
        <BrowserRouter basename='/'>
            <Route path={`/management`} component={ManagePage}/>
        </BrowserRouter>
    </Provider>
), document.getElementById('management'));
