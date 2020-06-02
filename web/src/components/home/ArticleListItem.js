import React from "react";
import PropTypes from 'prop-types';
import "./ArticleListItem.less"
import {Col, Row} from "antd";

/**
 * --------------------------------------
 * =================title=================
 * | ******  ---------------------------|
 * | *图片*  ------description----------|
 * | ******  ---------------------pv:6--|
 * --------------------------------------
 */
class ArticleListItem extends React.Component {
    constructor(props) {
        super(props);

        this.state = {};

    }

    render() {
        const {title, description, create_time, pv, short_content, picture} = this.props
        return (
            <div className="article-item">
                <Row>
                    <Col span={4} className="article-picture" >
                        {picture}
                    </Col>
                    <Col span={20}>
                        <div>1{title}</div>
                        <div>2{description}</div>
                        <div>3{create_time}</div>
                        <div>4{short_content}</div>
                        <div>5{pv}</div>
                    </Col>
                </Row>
            </div>
        );
    }
}

ArticleListItem.propTypes = {
    id: PropTypes.string.isRequired,
    title: PropTypes.string.isRequired,
    create_time: PropTypes.string.isRequired,
    pv: PropTypes.number.isRequired,
    short_content: PropTypes.string.isRequired,
    picture: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
};

export default ArticleListItem;