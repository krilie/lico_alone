import React from "react";
import "./SettingPage.less"
import {connect} from "react-redux";

class SettingPage extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            settings: []
        }

    }

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

export default SettingPage = connect((state) => ({...state}))(SettingPage);
