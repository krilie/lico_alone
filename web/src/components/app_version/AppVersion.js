import React from "react";
import "./AppVersion.less"
import {Row, Col} from "antd"
import {getVersion} from "../../api/ApiCommon";
import CopyToBoard from "../../utils/CopyToBoard";

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
    }

    render() {
        const {build_time, git_commit, go_version, version} = this.state;
        const buildTime = <Row>
            <Col flex="auto" className="text-left ellipsis text-size-small" title={build_time}>构建时间:&nbsp;&nbsp;{build_time}</Col>
        </Row>
        const gitCommit = <Row>
            <Col flex="auto"  className="text-left ellipsis text-size" title={git_commit}>
                <div style={{cursor:"pointer"}} onClick={()=>this.copyText(git_commit)}>散列值:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{git_commit}</div>
            </Col>
        </Row>
        const goVersion = <Row>
            <Col flex="auto" className="text-size text-left ellipsis" title={go_version}>Go版本:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{go_version}</Col>
        </Row>
        const appVersion = <Row>
            <Col flex="auto" className="text-size-small text-left ellipsis" title={version}>App版本:&nbsp;&nbsp;&nbsp;{version}</Col>
        </Row>
        return (
            <div className="appVersion">
                <Row><Col span={24}>{buildTime}</Col></Row>
                <Row><Col span={24}>{appVersion}</Col></Row>
                <Row><Col span={24}>{gitCommit}</Col></Row>
                <Row><Col span={24}>{goVersion}</Col></Row>
            </div>
        );
    }
}