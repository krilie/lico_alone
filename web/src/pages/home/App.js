import React from 'react';
import {Menu, Row, Col, BackTop, Affix} from 'antd';
import Logo from "../../components/logo/Logo";
import {Redirect, Route, Switch} from "react-router-dom";
import Article from "./article/Article";
import IcpLabel from "../../components/icp/IcpLabel";
import {withRouter} from "react-router-dom";
import "./App.less"
import Home from "./index/Home";
import RightCircleTwoTone from "@ant-design/icons/lib/icons/RightCircleTwoTone";
import {postVisited} from "../../api/ApiCommon";
import {GetCustomerTraceId} from "../../utils/LocalStorageUtil";

// 每个文件夹一个单独页面
class App extends React.Component {

    constructor(props) {
        super(props);
        console.log(props)
        this.state = {current: `${this.props.location.pathname}`}
    }

    componentDidMount() {
        process.nextTick(() => {
            this.props.history.push(`${this.state.current}`);
            postVisited(GetCustomerTraceId(), (res) => {
            })
        })
    }

    handleClick = e => {
        this.props.history.push(e.key);
        this.setState({current: e.key,});
    };

    render() {
        return (
            <div className="global-style">
                <Affix className="menu-affix" offsetTop={0}>
                    <Row className="fix-height-menu-all menu-affix">
                        <Col className="menu-logo" sm={12} xs={0}><Logo/></Col>
                        <Col className="fix-height-menu" sm={12} xs={24} style={{margin: "0px 0px 0px 20px"}}>
                            <Menu className="Menu" onClick={this.handleClick}
                                  selectedKeys={[this.state.current]}
                                  mode="horizontal">
                                <Menu.Item key={`${this.props.match.path}/index`}>主页</Menu.Item>
                                <Menu.Item key={`${this.props.match.path}/article`}>文章</Menu.Item>
                                <Menu.Item key="/management">
                                    <a
                                        title={"management"}
                                        href={"/management"}
                                        target="_parent"
                                        rel="noopener noreferrer">
                                        <RightCircleTwoTone className="array-router"/>
                                    </a>
                                </Menu.Item>
                            </Menu>
                        </Col>
                    </Row>
                </Affix>

                <Switch>
                    <Route path={`${this.props.match.path}/index`} component={Home}/>
                    <Route path={`${this.props.match.path}/article`} component={Article}/>
                    <Redirect path={`${this.props.match.path}/`} to={{pathname: `${this.props.match.path}/index`}}/>
                </Switch>
                <IcpLabel/>
                <BackTop>
                    <div className="up-button">up</div>
                </BackTop>
            </div>
        );
    }
}

export default withRouter(App);