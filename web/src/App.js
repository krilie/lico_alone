import React from 'react';
import './App.css';
import {Layout} from 'antd';

const {Content, Footer} = Layout;

function App() {
    return (
        <div className="App" style={{height: '100%'}}>
            <Layout className="layout" style={{height: '100%'}}>
                <Content className="site-layout" style={{padding: '0 50px', marginTop: 64, height: '100%'}}>
                    <div className="site-layout-background" style={{padding: 24, minHeight: 380}}>
                        建设中
                    </div>
                </Content>
                <Footer style={{textAlign: 'center'}}>
                    <a href="http://www.beian.miit.gov.cn" target="_blank" rel="noopener noreferrer">皖ICP备20003857号</a>
                </Footer>
            </Layout>
        </div>
    );
}

export default App;