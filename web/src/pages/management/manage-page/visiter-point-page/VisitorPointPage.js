import React, {Component} from 'react';
import {Map, Marker} from 'react-amap';

class VisitorPointPage extends Component {

    constructor(props) {
        super(props);
        this.state = {mapKey: ""}
    }

    componentDidMount() {
        // todo: 请求mapKey
        this.setState({
            mapKey: "123"
        })
    }

    render() {
        const {mapKey} = this.state;
        return mapKey === ""
            ?
            <div>loading</div>
            :
            <div>
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