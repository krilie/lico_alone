import React from "react";
import "./Home.less"
import {connect} from 'react-redux'

class Home extends React.Component {
    render() {
        return (<div>
            <div>您有444555件事未完成</div>
        </div>);
    }
}

export default Home = connect((state) => ({...state}))(Home);