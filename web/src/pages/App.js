import React from 'react';
import 'antd/dist/antd.css';
import {Menu, Row,Col,Divider} from 'antd';
import Logo from "../components/logo/Logo";
import {Route, Switch} from "react-router-dom";
import Share from "./share/Share";
import Photos from "./photos/Photos";
import Article from "./article/Article";
import IcpLabel from "../components/icp/IcpLabel";
import AppVersion from "../components/app_version/AppVersion";
import {withRouter, Redirect} from "react-router-dom";
import "./App.less"
import Home from "./home/Home";

class App extends React.Component {

    state = {
        current: '/home',
    };

    componentWillMount() {
        if (this.props.location.pathname === '/') {
            this.props.location.pathname = "/home";
        }
        const {pathname} = this.props.location;
        this.props.history.push(pathname);
        this.setState({
            current: pathname,
        });
    }

    handleClick = e => {
        this.props.history.push(e.key);
        this.setState({
            current: e.key,
        });
    };

    render() {
        return (
            <div>
                <Row>
                    <Col span={3}><Logo/></Col>
                    <Col span={11}>
                        <Menu className="Menu" onClick={this.handleClick} selectedKeys={[this.state.current]} mode="horizontal">
                            <Menu.Item key="/home">主页</Menu.Item>
                            <Menu.Item key="/article">博文</Menu.Item>
                            <Menu.Item key="/photos">图片</Menu.Item>
                            <Menu.Item key="/share">分享</Menu.Item>
                        </Menu>
                    </Col>
                    <Col span={8}><AppVersion/></Col>
                    <Col span={2}/>
                </Row>
                <Divider orientation="left" className="div-line"/>
                <Switch>
                    <Route exact path="/home" component={Home}/>
                    <Route exact path="/share" component={Share}/>
                    <Route exact path="/photos" component={Photos}/>
                    <Route exact path="/article" component={Article}/>
                    <Redirect path="/" to={{pathname: '/home'}}/>
                </Switch>
                <IcpLabel/>
            </div>
        );
    }
}

export default withRouter(App);