import React from "react";
import "./SettingCard.less"
import {Button, Card, Col, Input, message, Row} from "antd";
import {updateSettingItem} from "../../api/ManageSettingApi";

const {TextArea} = Input

/**
 * 编辑json数据的类
 * 目标一 编辑一级json字段
 */
export default class SettingCard extends React.Component {

    constructor(props) {
        super(props);
        try {
            this.state = {
                value: JSON.stringify(JSON.parse(this.props.data.value), null, 2),
            }
        } catch (e) {
            this.state = {
                value: this.props.data.value,
            }
        }
    }

    // 更新配置项
    upDataSetting = (e) => {
        console.log("update value " + e.target.value)
        this.setState({
            value: e.target.value
        })
    }

    onCommitSetting = () => {
        const {name} = this.props.data
        let {value} = this.state
        try {
            value = JSON.stringify(JSON.parse(value))
        } catch (e) {
            value = this.state.value
        }
        updateSettingItem(name, value).then(data => {
            message.info("保存成功");
        }).catch(err => {
            console.log("保存操作：" + name + value, err.status)
        })
    }

    render() {
        const {name, create_time} = this.props.data
        const {value} = this.state
        const extra = <Button type="primary"
                              className="extra-area-button" size="large"
                              onClick={this.onCommitSetting}>保存</Button>
        const title =
            <Row className="setting-card-title">
                <Col span={12}>{name}</Col>
                <Col span={12} style={{textAlign: "right"}}>创建时间:{create_time}</Col>
            </Row>
        return <Card className="setting-card-style" title={title}>
            <TextArea className="setting-text-area"
                      rows={3}
                      onChange={(e) => this.upDataSetting(e)}
                      defaultValue={value}/>
            <div className="text-right">{extra}</div>
        </Card>
    }
}