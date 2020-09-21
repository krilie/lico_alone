import React, {Component} from 'react';
import {Map, Marker} from 'react-amap';
import {getAMapKey} from "../../../../api/ManageSettingApi";
import {getVisitorPoints} from "../../../../api/ManageStatisticApi";

class VisitorPointPage extends Component {

    constructor(props) {
        super(props);
        this.state = {mapKey: "", marker: []}
    }

    componentDidMount() {
        getAMapKey().then(res => {
            this.setState({
                mapKey: res.data.data.a_map_key,
            })
        })
        getVisitorPoints().then(res => {
            this.setState({
                marker: res.data.data,
            })
        })


    }

    render() {
        const {mapKey, marker} = this.state;
        const Markers = marker.map(m => <Marker label={<div>{m.city}</div>} autoRotation={true} position={[m.lon, m.lat]}/>)
        return mapKey === ""
            ?
            <div>loading...</div>
            :
            <div style={{height: "100%", width: "100%"}}>
                <Map
                    amapkey={mapKey}
                    plugins={["ToolBar", 'Scale', 'OverView', 'MapType', 'Geolocation']}
                    center={["116.397128", "39.916527"]}
                    zoom={15}>
                    {Markers}
                </Map>
            </div>
            ;
    }
}

export default VisitorPointPage;