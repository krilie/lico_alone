import React from "react";
import "./JsonView.less"
import {Button} from "antd";

/**
 * 编辑json数据的类
 * 目标一 编辑一级json字段
 */
export default class JsonView extends React.Component {
    constructor(props) {
        super(props); // data onDataChange
    }

    render() {
        const {data, onDataChange} = this.props
        if (data instanceof Object) {
            return <div>
                <div>
                    {data.toJSON().map((key, value) =>
                        <div>
                            <div>theKey {key}</div>
                            <div>theValue {value}</div>
                        </div>)}
                </div>
                <Button type="primary" onClick={() => onDataChange(data)}>确认</Button>
            </div>
        }
    }
}