import {Layout, Menu} from 'antd';
import React from "react";
import "./ManagePage.less"
import {Route, Switch} from "react-router";
import SettingPage from "./setting-page/SettingPage";
import FilePage from "./file-page/FilePage";
import CarouselPage from "./carousel-page/CarouselPage";
import ArticleHomePage from "./article-page/ArticleHomePage";
import {Redirect} from "react-router-dom";
import VisitorPointPage from "./visiter-point-page/VisitorPointPage";

const {Sider} = Layout;

export default class ManagePage extends React.Component {

    constructor(props) {
        console.log(props)
        super(props);
        const {pathname} = this.props.location;
        this.state = {
            currentPage: pathname
        }
    }

    goToPage = path => this.props.history.push(path)

    handleClick = e => {
        console.log(e.key)
        this.props.history.push(e.key);
        this.setState({
            currentPage: e.key,
        });
    };

    render() {
        return (
            <Layout>
                <Sider className="sider-layout-background" width={"60px"}>
                    <Menu onClick={this.handleClick}
                          selectedKeys={[this.state.currentPage]}
                          className="sider-layout-background"
                          mode="inline"
                          defaultSelectedKeys={['4']}>
                        <Menu.Item className="v-center" key = {`${this.props.match.path}/setting`}>设置</Menu.Item>
                        <Menu.Item className="v-center" key = {`${this.props.match.path}/files`}>文件</Menu.Item>
                        <Menu.Item className="v-center" key = {`${this.props.match.path}/carousel`}>轮播图</Menu.Item>
                        <Menu.Item className="v-center" key = {`${this.props.match.path}/article`}>文章</Menu.Item>
                        <Menu.Item className="v-center" key = {`${this.props.match.path}/visitor_point`}>访问</Menu.Item>
                    </Menu>
                </Sider>
                <Layout className="manage-layout">
                    <Switch>
                        <Route  path={`${this.props.match.path}/setting`} component={SettingPage}/>
                        <Route  path={`${this.props.match.path}/files`} component={FilePage}/>
                        <Route  path={`${this.props.match.path}/carousel`} component={CarouselPage}/>
                        <Route  path={`${this.props.match.path}/article`} component={ArticleHomePage}/>
                        <Route  path={`${this.props.match.path}/visitor_point`} component={VisitorPointPage}/>
                        <Redirect path={`${this.props.match.path}/`}
                                  to={{pathname: `${this.props.match.path}/setting`}}/>
                    </Switch>
                </Layout>
            </Layout>
        );
    }
}