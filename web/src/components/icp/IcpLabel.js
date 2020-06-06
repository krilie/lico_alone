import React from "react";
import "./Icplable.less"
import {getIcpInfo} from "../../api/ApiCommon";

export default class IcpLabel extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            name:"",link:"",label:"",
        }
    }

    componentDidMount() {
        getIcpInfo(data=>{
            this.setState({
                ...data
            })
        });
    }

    render() {
        const {name,link,label} = this.state;
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