import React from "react";
import "./SettingPage.less"
import {connect} from "react-redux";
import {getSettingListAllRedux} from "../../../../api/SettingApi";
import store from "../../../../redux/RuduxIndex"
import JsonView from "../../../../components/json_view/JsonView";
import openNotification from "../../../../utils/MessageBoard";
//import {Col, Row} from "antd";

class SettingPage extends React.Component {

    constructor(props) {
        super(props);
        store.dispatch(getSettingListAllRedux())
    }

    goToPage = path => {
        this.props.history.push(path);
    };

    render() {
        const {settings} = this.props
        return (
            <div className="setting-height">
                {settings.map(val =>
                    <div>
                        <div >{val.name}</div>
                        <div >{val.create_time}</div>
                        <div >{val.value}</div>
                        <br/>
                        <JsonView data={val} onDataOk={(data) => openNotification(data)}/>
                    </div>)}
            </div>
        );
    }
}

export default SettingPage = connect((state) => ({...state}))(SettingPage);
