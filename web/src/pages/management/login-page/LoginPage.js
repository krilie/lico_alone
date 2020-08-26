import React from "react";
import "./LoginPage.less"
import {SetUserToken} from "../../../utils/LocalStorageUtil";
import {Button, Col, Form, Input, message, Row} from "antd";
import {userLogin} from "../../../api/ManageUserApi";

/**
 * 1.检查token 如果token不为空或null 则跳转主管理页面
 * 2.执行用户登录操作 操作成功设置token并跳转到主管理页面
 */
export default class LoginPage extends React.Component {

    goToPage = path => this.props.history.push(path);

    setToken(token) {
        SetUserToken(token)
    }

    render() {
        const layout = {
            labelCol: {span: 8},
            wrapperCol: {span: 16},
        };
        const tailLayout = {
            wrapperCol: {offset: 12, span: 12},
        };

        const onFinish = values => {
            // 登录
            userLogin(values).then(res => {
                if (res.data.code === 2000) {
                    this.setToken(res.data.data.token)
                    this.goToPage("/management/manage")
                } else {
                    message.error("from login page:" + res.data.message + res.data.detail)
                }
            }).then(err => {
                if (err !== undefined)
                    message.error("from login page:" + err.toString())
            }).finally(() => {
            })
        };

        const onFinishFailed = errorInfo => console.log('Failed:', errorInfo);

        const form = <Form
            {...layout}
            name="basic"
            initialValues={{remember: true}}
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
        >
            <Form.Item
                label="手机号"
                name="phone"
                rules={[{required: true, message: '手机号'}]}
            >
                <Input/>
            </Form.Item>

            <Form.Item
                label="密码"
                name="password"
                rules={[{required: true, message: '请输入密码'}]}
            >
                <Input.Password/>
            </Form.Item>

            <Form.Item {...tailLayout}>
                <Button type="primary" htmlType="submit">
                    登录
                </Button>
            </Form.Item>
        </Form>

        return (
            <Row className="main-row">
                <Col span={2}/>
                <Col span={16}>{form}</Col>
                <Col span={6}/>
            </Row>
        );
    };
}
