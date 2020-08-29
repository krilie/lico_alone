import React from "react";
import "./Home.less"
import {connect} from 'react-redux'
import {Col, Row} from "antd";
import SlidePictures from "../../../components/home/SlidePictures";
import ArticleListPageRollView from "../../../components/home/ArticleListPageRollView";
import AppVersion from "../../../components/app_version/AppVersion";

class Home extends React.Component {
    render() {
        return (<div className="home">
            <Row className="main-home-area">
                <Col sm={16} xs={24} className="article-area">
                    <ArticleListPageRollView/>
                </Col>
                <Col sm={8} xs={0} className="status-area">
                    <div><AppVersion/></div>
                    <div style={{margin:"3px"}}><a href="/home/article">->文章区</a></div>
                    <div className="slide-picture">
                        <SlidePictures className="slide-picture"/>
                    </div>
                </Col>
            </Row>
        </div>);
    }
}

export default Home = connect((state) => ({...state}))(Home);