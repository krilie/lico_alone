import React from "react";
import "./SettingPage.less"
import {connect} from "react-redux";
import {getSettingListAllRedux} from "../../../../api/SettingApi";
import store from "../../../../redux/RuduxIndex"
import JsonView from "../../../../components/json_view/JsonView";
import openNotification from "../../../../utils/MessageBoard";

class SettingPage extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            settings: []
        }
        store.dispatch(getSettingListAllRedux())
    }

    goToPage = path => {
        this.props.history.push(path);
    };

    render() {
        const {settings} = this.state
        const sets = settings.map(val=>
            <div>
                <JsonView data={val} onDataOk={(data) => openNotification(data)}/>
            </div>
        )
        return (
            <div>
                {sets}
            </div>
        );
    }
}

export default SettingPage = connect((state) => ({...state}))(SettingPage);
