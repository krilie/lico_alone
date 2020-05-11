import {Layout, Menu} from 'antd';
import { UploadOutlined, UserOutlined, VideoCameraOutlined } from '@ant-design/icons';
import React from "react";
import "./ManagePage.less"
import { Route, Switch} from "react-router";
import SettingPage from "./setting-page/SettingPage";

const { Header, Content, Footer, Sider } = Layout;

export default class ManagePage extends React.Component {

    componentWillMount() {
        const {pathname} = this.props.location;
        this.props.history.push(pathname);
    }

    constructor(props) {
        super(props);
        const {pathname} = this.props.location;
        this.state = {
            currentPage:pathname
        }
    }

    goToPage = path => this.props.history.push(path)

    handleClick = e => {
        this.props.history.push(e.key);
        this.setState({
            currentPage: e.key,
        });
    };

    render() {
        return (
            <Layout>
                <Sider className="sider-layout-background">
                    <Menu onClick={this.handleClick} selectedKeys={[this.state.currentPage]}  className="sider-layout-background" mode="inline" defaultSelectedKeys={['4']}>
                        <Menu.Item key="/management/manage/setting" icon={<UserOutlined />}>
                            设置
                        </Menu.Item>
                        <Menu.Item key="2" icon={<VideoCameraOutlined />}>
                            nav 2
                        </Menu.Item>
                        <Menu.Item key="3" icon={<UploadOutlined />}>
                            nav 3
                        </Menu.Item>
                        <Menu.Item key="4" icon={<UserOutlined />}>
                            nav 4
                        </Menu.Item>
                    </Menu>
                </Sider>
                <Layout  className="manage-layout">
                    <Switch>
                        <Route exact path="/management/manage/setting" component={SettingPage}/>
                    </Switch>
                </Layout>
            </Layout>
        );
    }
}