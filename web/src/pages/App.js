import React from 'react';
import 'antd/dist/antd.css';
import {Menu, Row, Col, Divider, BackTop, Affix} from 'antd';
import Logo from "../components/logo/Logo";
import {Route, Switch} from "react-router-dom";
import Article from "./article/Article";
import IcpLabel from "../components/icp/IcpLabel";
import {withRouter, Redirect} from "react-router-dom";
import "./App.less"
import Home from "./home/Home";
import RightCircleTwoTone from "@ant-design/icons/lib/icons/RightCircleTwoTone";
import Management from "./management/Management";
import RhythmStateHolder from "../utils/RhythmStateHolder";
import ArticleDetailPage from "./article/ArticleDetailPage";
import {postVisited} from "../api/ApiCommon";
import {GetCustomerTraceId} from "../utils/LocalStorageUtil";

class App extends React.Component {

    state = {
        current: '/home',
    };

    componentWillMount() {
        // if (this.props.location.pathname === '/') {
        //     this.props.location.pathname = "/home";
        // }
        const {pathname} = this.props.location;
        // this.props.history.push(pathname);
        this.setState({
            current: pathname,
        });

        postVisited(GetCustomerTraceId(), (res) => {
        })

    }

    handleClick = e => {
        this.props.history.push(e.key);
        this.setState({
            current: e.key,
        });
    };

    times = new RhythmStateHolder()
    toManagePage = () => {
        if (this.times.PushStateOrReset() === true) {
            this.handleClick({key: "/management"})
        } else {

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
            <div className="global-style">
                <Affix className="menu-affix" offsetTop={0}>
                    <Row className="fix-height-menu menu-affix" justify="start" align="middle">
                        <Col className="menu-logo" flex="150px">
                            <Logo/>
                        </Col>
                        <Col className="fix-height-menu" flex="auto">
                            <Menu className="Menu" onClick={this.handleClick}
                                  selectedKeys={[this.state.current]}
                                  mode="horizontal">
                                <Menu.Item key="/home">主页</Menu.Item>
                                <Menu.Item key="/article">文章</Menu.Item>
                                {/*<Menu.Item key="/photos">图片</Menu.Item>*/}
                                {/*<Menu.Item key="/share">动态</Menu.Item>*/}
                                <Menu.Item key="/management" onClick={() => this.toManagePage()}>
                                    <RightCircleTwoTone className="array-router"/>
                                </Menu.Item>
                            </Menu>
                        </Col>
                    </Row>
                    <Divider orientation="left" className="div-line"/>
                </Affix>

                <Switch>
                    <Route exact path="/home" component={Home}/>
                    {/*<Route exact path="/share" component={Share}/>*/}
                    {/*<Route exact path="/photos" component={Photos}/>*/}
                    <Route exact path="/article" component={Article}/>
                    <Route exact path="/article/:articleId" component={ArticleDetailPage}/>
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