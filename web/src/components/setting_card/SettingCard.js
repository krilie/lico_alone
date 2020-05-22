import React from "react";
import "./SettingCard.less"
import {Button, Card, Input} from "antd";

const {TextArea} = Input

/**
 * 编辑json数据的类
 * 目标一 编辑一级json字段
 */
export default class SettingCard extends React.Component {
    // 更新配置项
    upDataSetting = (data) => {
        console.log(data.name + "  " + data.value)
        this.setState({
            ...data
        })
    }

    onCommitSetting = () => {
        const {name, value} = this.state
        console.log("保存操作：" + name + " " + value)
    }

    render() {
        const {name, create_time, value} = this.props.data
        const extra = <div>
            <p>{create_time}</p>
            <Button type="primary" size="small" onChange={this.onCommitSetting}>保存</Button>
        </div>
        return <Card title={name} extra={extra} style={{width: 300}}>
            <TextArea rows={4} onChange={(e) => this.upDataSetting({name, value: e.value})} defaultValue={value}/>
        </Card>
    }
}