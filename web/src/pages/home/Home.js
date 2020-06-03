import React from "react";
import "./Home.less"
import {connect} from 'react-redux'
import {Col, Row} from "antd";
import SlidePictures from "../../components/home/SlidePictures";
import ArticleListPageRollView from "../../components/home/ArticleListPageRollView";
import AppVersion from "../../components/app_version/AppVersion";

/**
 * ================================================
 * ------------------------------------------------
 * |      图片+文字     滚动图                     |
 * |                                              |
 * -------------------------------------------------
 *                                      |
 *                                      |     状
 *       文章列表区                      |     态
 *                                      |     区
 *                                      |     广
 *                                      |     告
 *                                      |     区
 *                                      |
 *==================================================
 */
class Home extends React.Component {
    render() {
        return (<div className="home">
            <Row >
                <Col span={24} className="slide-picture">
                    <SlidePictures/>
                </Col>
            </Row>
            <Row className="main-home-area">
                <Col span={16} className="article-area"><ArticleListPageRollView/></Col>
                <Col span={8} className="status-area">
                    <AppVersion/>
                </Col>
            </Row>
        </div>);
    }
}

export default Home = connect((state) => ({...state}))(Home);