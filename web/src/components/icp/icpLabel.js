import React from "react";
import "./icplable.less"
import {getIcpInfo} from "../../api/common";

export default class IcpLabel extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            name:"",link:"",label:"",
        }
    }

    componentDidMount() {
        var data = getIcpInfo();
        this.setState({
            ...data
        })
    }

    render() {
        var {name,link,label} = this.state;
        return (
            <div className="icpLableDiv" >
                <a
                    title={label}
                    href={link}
                    target="_blank"
                    rel="noopener noreferrer">{name}</a>
            </div>
        );
    }
}