import React from 'react';
import './App.css';
import {Row,Col} from 'antd';

function App() {
    return (
        <div className="App" style={{height: '100%'}}>
            <Row span={1}>
                <Col span={24}>
                    建设中
                </Col>
            </Row>
            <Row span={22}>
                <Col span={24}>col</Col>
            </Row>
            <Row span={1}>
                <Col span={24}>
                    <a href="http://www.beian.miit.gov.cn" target="_blank" rel="noopener noreferrer">皖ICP备20003857号</a>
                </Col>
            </Row>
        </div>
    );
}

export default App;