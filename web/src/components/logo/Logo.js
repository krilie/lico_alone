import React from "react";
import imageLogo from "../../images/mttc.png"
import "./Logo.less"

export default class Logo extends React.Component {
    render() {
        return (
            <a href={"/"}><img src={ imageLogo }  alt={"Logo"}/></a>
        );
    }
}