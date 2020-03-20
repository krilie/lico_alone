import React from 'react';
import './App.css';
import { Layout } from 'antd';

const {  Content, Footer } = Layout;

function App() {
  return (
    <div className="App">
        <Layout className="layout">
            <Content/>
            <Footer style={{ textAlign: 'center' }}>
                <a href="www.beian.miit.gov.cn">皖ICP备20003857号</a>
            </Footer>
        </Layout>
    </div>
  );
}

export default App;