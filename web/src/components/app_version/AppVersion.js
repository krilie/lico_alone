import React from "react";
import "./AppVersion.less"
import {Row,Col} from "antd"
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
        return (
            <div className="appVersion">
                <Row>
                    <Col span={12}>{build_time}</Col>
                    <Col span={12}>{git_commit}</Col>
                </Row>
                <Row>
                    <Col span={12}>{go_version}</Col>
                    <Col span={12}>{version}</Col>
                </Row>
            </div>
        );
    }
}