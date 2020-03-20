import React from 'react';
import './App.css';
import { Layout } from 'antd';

const {  Content, Footer } = Layout;

function App() {
  return (
    <div className="App">
        <Layout className="layout">
            <Content style={{ padding: '0 50px' }}>
                <div className="site-layout-content">建设中</div>
            </Content>
            <Footer style={{ textAlign: 'center' }}>
                <a href="http://www.beian.miit.gov.cn" target="_blank"  rel="noopener noreferrer">皖ICP备20003857号</a>
            </Footer>
        </Layout>
    </div>
  );
}

export default App;