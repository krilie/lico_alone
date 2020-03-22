import React from 'react';
import './App.css';
import Icplable from "./components/icp/icplable";
import {Layout} from 'antd';
import Row from "antd/es/descriptions/Row";
import Col from "antd/es/grid/col";

const {Header, Footer, Content} = Layout;

function App() {
    return (
        <div className="App" style={{height: '100%'}}>
            <Row gutter={[8, 8]}>
                <Col span={12} />
                <Col span={12} />
            </Row>
            <Row gutter={[8, 8]}>
                <Col span={12} />
                <Col span={12} ><Icplable/></Col>
            </Row>
        </div>
    );
}

export default App;