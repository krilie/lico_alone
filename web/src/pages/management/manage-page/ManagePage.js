import {Layout, Menu} from 'antd';
import React from "react";
import "./ManagePage.less"
import {Route, Switch} from "react-router";
import SettingPage from "./setting-page/SettingPage";
import FilePage from "./file-page/FilePage";
import CarouselPage from "./carousel-page/CarouselPage";
import ArticleHomePage from "./article-page/ArticleHomePage";

const {Sider} = Layout;

export default class ManagePage extends React.Component {

    componentWillMount() {
    }

    constructor(props) {
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
                        <Menu.Item className="v-center" key="/manage/setting">设置</Menu.Item>
                        <Menu.Item className="v-center" key="/manage/files">文件</Menu.Item>
                        <Menu.Item className="v-center" key="/manage/carousel">轮播图</Menu.Item>
                        <Menu.Item className="v-center" key="/manage/article">Article</Menu.Item>
                    </Menu>
                </Sider>
                <Layout className="manage-layout">
                    <Switch>
                        <Route exact path="/manage/setting" component={SettingPage}/>
                        <Route exact path="/manage/files" component={FilePage}/>
                        <Route exact path="/manage/carousel" component={CarouselPage}/>
                        <Route path="/manage/article" component={ArticleHomePage}/>
                    </Switch>
                </Layout>
            </Layout>
        );
    }
}