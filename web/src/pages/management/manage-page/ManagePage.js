import { Layout, Menu } from 'antd';
import { UploadOutlined, UserOutlined, VideoCameraOutlined } from '@ant-design/icons';
import React from "react";
import "./ManagePage.less"

const { Header, Content, Footer, Sider } = Layout;

export default class ManagePage extends React.Component {

    goToPage = path => {
        this.props.history.push(path);
    };

    render() {
        return (
            <Layout className="manage-layout">
                <Sider>
                    <Menu theme="white" mode="inline" defaultSelectedKeys={['4']}>
                        <Menu.Item key="1" icon={<UserOutlined />}>
                            nav 1
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
                <Layout>

                    334423423412341234

                </Layout>
            </Layout>
        );
    }
}