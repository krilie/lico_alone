import React from "react";
import "./SettingPage.less"
import {connect} from "react-redux";
import {getSettingListAllRedux} from "../../../../api/SettingApi";
import store from "../../../../redux/RuduxIndex"
import SettingCard from "../../../../components/setting_card/SettingCard";

class SettingPage extends React.Component {

    componentWillMount() {
        // const {pathname} = this.props.location;
        // this.props.history.push(pathname);
    }

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
                    <SettingCard key={val.name} data={val}/>
                )}
            </div>
        );
    }
}

export default SettingPage = connect((state) => ({...state}))(SettingPage);
