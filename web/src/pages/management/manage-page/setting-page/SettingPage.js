import React from "react";
import "./SettingPage.less"

export default class SettingPage extends React.Component {

    goToPage = path => {
        this.props.history.push(path);
    };

    render() {
        return (
            <div>
                setting
            </div>
        );
    }
}