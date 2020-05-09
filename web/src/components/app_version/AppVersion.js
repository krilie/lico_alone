import React from "react";
import "./AppVersion.less"
import {Row, Col} from "antd"
import {getVersion} from "../../api/common";

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

    render() {
        const {build_time, git_commit, go_version, version} = this.state;
        const buildTime = <Row>
            <Col flex="50px" className="text-right ellipsis text-size">构建时间</Col>
            <Col flex="auto" className="text-left ellipsis text-size">{build_time}</Col>
        </Row>
        const gitCommit = <Row>
            <Col flex="50px"  className="text-right ellipsis text-size">散列值</Col>
            <Col flex="auto"  className="text-left ellipsis text-size">{git_commit}</Col>
        </Row>
        const goVersion = <Row>
            <Col flex="50px" className="text-size text-right ellipsis">Go版本</Col>
            <Col flex="auto" className="text-size text-left ellipsis">{go_version}</Col>
        </Row>
        const appVersion = <Row>
            <Col flex="50px" className="text-size text-right ellipsis">App版本</Col>
            <Col flex="auto" className="text-size text-left ellipsis">{version}</Col>
        </Row>
        return (
            <div className="appVersion">
                <Row><Col span={12}>{buildTime}</Col><Col span={12}>{gitCommit}</Col></Row>
                <Row><Col span={12}>{goVersion}</Col><Col span={12}>{appVersion}</Col></Row>
            </div>
        );
    }
}