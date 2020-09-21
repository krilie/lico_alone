import React, {Component} from 'react';
import {Map, Marker} from 'react-amap';
import {getAMapKey} from "../../../../api/ManageSettingApi";

class VisitorPointPage extends Component {

    constructor(props) {
        super(props);
        this.state = {mapKey: ""}
    }

    componentDidMount() {
        getAMapKey().then(res => {
            this.setState({
                mapKey: res.data.data.a_map_key,
            })
        })
    }

    render() {
        const {mapKey} = this.state;
        return mapKey === ""
            ?
            <div>loading...</div>
            :
            <div style={{height:"100%",width:"100%"}}>
                <Map
                    amapkey={mapKey}
                    plugins={["ToolBar", 'Scale']}
                    center={["116.397128", "39.916527"]}
                    zoom={15}>
                    <Marker position={["116.397128", "39.916527"]}/>
                </Map>
            </div>
            ;
    }
}

export default VisitorPointPage;