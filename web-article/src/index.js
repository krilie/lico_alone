import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Article from './Article';
import {ToastContainer} from 'react-toastify';
import {Route ,BrowserRouter} from "react-router-dom";

ReactDOM.render(
    <BrowserRouter basename='/'>
        <React.StrictMode>
            <Route path={`/`} component={Article}/>
            <ToastContainer/>
        </React.StrictMode>
    </BrowserRouter>,
    document.getElementById('root')
);
