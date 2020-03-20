import React from 'react';
import './App.css';
import Icplable from "./icp/icplable";
import {Layout} from 'antd';

const {Header, Footer, Content} = Layout;

function App() {
    return (
        <div className="App" style={{height: '100%'}}>
            <Layout>
                <Header>Header</Header>
                <Content>Content</Content>
                <Footer><Icplable/></Footer>
            </Layout>
        </div>
    );
}

export default App;