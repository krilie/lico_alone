import React from "react";
import "./SettingPage.less"
import {connect} from "react-redux";
import {getSettingListAllRedux} from "../../../../api/SettingApi";
import store from "../../../../redux/RuduxIndex"
import {Input} from "antd";

const {TextArea} = Input

class SettingPage extends React.Component {

    constructor(props) {
        super(props);
        store.dispatch(getSettingListAllRedux())
    }

    goToPage = path => {
        this.props.history.push(path);
    };

    handleUpdate = (name,value)=>{

    }

    render() {
        const {settings} = this.props
        return (
            <div className="setting-height">
                {settings.map(val =>
                    <div>
                        <div>{val.name}</div>
                        <TextArea rows={4} onChange={(e)=>console.log(e.value)} defaultValue={val.value}/>
                    </div>)}
            </div>
        );
    }
}

export default SettingPage = connect((state) => ({...state}))(SettingPage);
