import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from "react-redux";
import App from "./pages/App";
import {HashRouter, Route, Switch} from 'react-router-dom'
import store from "./redux/RuduxIndex";

ReactDOM.render((
    <Provider store={store}>
        <HashRouter basename='/'>
            <Route path={`/`} component={App}/>
        </HashRouter>
    </Provider>
), document.getElementById('root'));
