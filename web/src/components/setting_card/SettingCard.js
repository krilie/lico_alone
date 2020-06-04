import React from "react";
import "./SettingCard.less"
import {Button, Card, Input, message} from "antd";
import {updateSettingItem} from "../../api/SettingApi";

const {TextArea} = Input

/**
 * 编辑json数据的类
 * 目标一 编辑一级json字段
 */
export default class SettingCard extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            value: this.props.data.value,
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
        const {value} = this.state
        updateSettingItem(name, value).then(data => {
            message.info("保存成功");
        }).catch(err => {
            console.log("保存操作：" + name + value, err.status)
        })
    }

    render() {
        const {name, create_time, value} = this.props.data
        const extra = <Button type="primary" style={{margin: '0px 0px 0px 0px'}} size="large"
                              onClick={this.onCommitSetting}>保存</Button>
        const title = <div style={{padding: "3px 3px 3px 3px", margin: "3px 3px 3px 3px"}}>项目:  {name} 创建时间: {create_time}</div>
        return <Card bodyStyle={{padding: "3px 3px 3px 3px", margin: "3px 3px 3px 3px"}}
                     headStyle={{padding: "0px 0px 0px 0px", margin: "0px 0px 0px 0px"}}
                     title={title}>
            <TextArea style={{fontSize: '20px'}} rows={2} onChange={(e) => this.upDataSetting(e)} defaultValue={value}/>
            {extra}
        </Card>
    }
}