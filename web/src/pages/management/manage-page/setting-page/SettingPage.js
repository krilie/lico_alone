import React from "react";
import "./SettingPage.less"
import {connect} from "react-redux";
import {getSettingListAllRedux} from "../../../../api/ManageSettingApi";
import store from "../../../../redux/RuduxIndex"
import SettingCard from "../../../../components/setting_card/SettingCard";
import {Card} from "antd";

class SettingPage extends React.Component {

    goToPage = path => this.props.history.push(path)

    constructor(props) {
        super(props);
        store.dispatch(getSettingListAllRedux())
    }

    render() {
        const {settings} = this.props
        return (
            <Card className="setting-card" title={<div>配置</div>}>
                {settings.map(val => <SettingCard key={val.name} data={val}/>)}
            </Card>
        );
    }
}

export default SettingPage = connect((state) => ({...state}))(SettingPage);
