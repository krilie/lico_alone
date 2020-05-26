import React from "react";
import "./Home.less"
import {connect} from 'react-redux'

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
        return (<div>
            <div>您有444555件事未完成</div>
        </div>);
    }
}

export default Home = connect((state) => ({...state}))(Home);