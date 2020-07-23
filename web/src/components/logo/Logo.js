import React from "react";
import imageLogo from "../../images/mttc.png"
import "./Logo.less"

export default class Logo extends React.Component {
    render() {
        return (
            <div className="logo">
                <a className="logo" href={"/"}>
                    <img className="logo" src={imageLogo} alt={"Logo"}/>
                </a>
            </div>
        );
    }
}