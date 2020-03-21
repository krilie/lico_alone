import React from 'react';
import './App.css';
import Icplable from "./icp/icplable";
import {Layout} from 'antd';

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
                <Col span={12} />
            </Row>
            <Layout>
                <Header>Header</Header>
                <Content>Content</Content>
                <Footer><Icplable/></Footer>
            </Layout>
        </div>
    );
}

export default App;