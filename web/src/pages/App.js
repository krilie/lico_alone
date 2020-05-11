import React from 'react';
import 'antd/dist/antd.css';
import {Menu, Row, Col, Divider, BackTop} from 'antd';
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
import RightCircleTwoTone from "@ant-design/icons/lib/icons/RightCircleTwoTone";
import Management from "./management/Management";
import RhythmStateHolder from "../utils/RhythmStateHolder";

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

    times = new RhythmStateHolder()
    toManagePage = ()=>{
        if (this.times.PushStateOrReset() === true) {
            this.handleClick({key: "/management"})
        }else{

        }
    }




    render() {

       const style = {
            height: 40,
            width: 40,
            lineHeight: '40px',
            borderRadius: 4,
            backgroundColor: '#1088e9',
            color: '#fff',
            textAlign: 'center',
            fontSize: 14,
        };

        return (
            <div>
                <Row className="fix-height-menu" justify="start" align="middle">
                    <Col className="menu-logo" flex="0 1 150px">
                        <Logo/>
                    </Col>
                    <Col className="fix-height-menu" flex="1 1 250px">
                        <Menu className="Menu" onClick={this.handleClick} selectedKeys={[this.state.current]} mode="horizontal">
                            <Menu.Item key="/home">主页</Menu.Item>
                            <Menu.Item key="/article">博文</Menu.Item>
                            <Menu.Item key="/photos">图片</Menu.Item>
                            <Menu.Item key="/share">分享</Menu.Item>
                        </Menu>
                    </Col>
                    <Col flex="0 1 310px"> <AppVersion/> </Col>
                    <Col flex="10px"/>
                    <Col flex="0 1 20px"> <RightCircleTwoTone onClick={()=>this.toManagePage()} /></Col>
                </Row>
                <Divider orientation="left" className="div-line"/>
                <Switch>
                    <Route exact path="/home" component={Home}/>
                    <Route exact path="/share" component={Share}/>
                    <Route exact path="/photos" component={Photos}/>
                    <Route exact path="/article" component={Article}/>
                    <Route exact={false} path="/management" component={Management}/>
                    <Redirect path="/" to={{pathname: '/home'}}/>
                </Switch>
                <IcpLabel/>
                <BackTop>
                    <div style={style}>UP</div>
                </BackTop>
            </div>
        );
    }
}

export default withRouter(App);