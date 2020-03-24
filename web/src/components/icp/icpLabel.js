import React from "react";
import "./icplable.less"

export default class IcpLabel extends React.Component {
    render() {
        return (
            <div className="icpLableDiv" >
                <a
                    title="皖ICP备20003857号-1"
                    href="http://www.beian.miit.gov.cn"
                    target="_blank"
                    rel="noopener noreferrer">皖ICP备20003857号-1</a>
            </div>
        );
    }
}