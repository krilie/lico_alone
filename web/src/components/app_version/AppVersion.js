import React from "react";
import "./AppVersion.less"
import {Row, Col,Button} from "antd"
import {getVersion} from "../../api/common";
import CopyToBoard from "../../utils/CopyToBoard";
import openNotification from "../../utils/MessageBoard";

export default class AppVersion extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            build_time: "", git_commit: "", go_version: "", version: ""
        }
    }

    // {"build_time":"2020-05-08 22:26:28","git_commit":"a22dd43a7b8ed831a0908e0ea97aab1bbd9a3181","go_version":"go version go1.14.2 linux/amd64","version":"v2.2.3"}
    componentDidMount() {
        getVersion(data => {
            this.setState({
                ...data
            })
        });
    }

    copyText = (text)=>{
        CopyToBoard(text)
        openNotification("copy success")
    }

    render() {
        const {build_time, git_commit, go_version, version} = this.state;
        const buildTime = <Row>
            <Col flex="auto" className="text-left ellipsis text-size-small" title={build_time}>构建时间: {build_time}</Col>
        </Row>
        const gitCommit = <Row>
            <Col flex="auto"  className="text-left ellipsis text-size" title={git_commit}>
                <div style={{cursor:"pointer"}} onClick={()=>this.copyText(git_commit)}> 散列值 : {git_commit}</div>
            </Col>
        </Row>
        const goVersion = <Row>
            <Col flex="auto" className="text-size text-left ellipsis" title={go_version}>Go版本 : {go_version}</Col>
        </Row>
        const appVersion = <Row>
            <Col flex="auto" className="text-size-small text-left ellipsis" title={version}>App版本 : {version}</Col>
        </Row>
        return (
            <div className="appVersion">
                <Row><Col span={8}>{buildTime}</Col><Col span={16}>{gitCommit}</Col></Row>
                <Row><Col span={8}>{appVersion}</Col><Col span={16}>{goVersion}</Col></Row>
            </div>
        );
    }
}