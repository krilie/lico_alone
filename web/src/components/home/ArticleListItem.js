import React from "react";
import PropTypes from 'prop-types';
import "./ArticleListItem.less.less"
import {Col, Row} from "antd";

class ArticleListItem extends React.Component {
    constructor(props) {
        super(props);

        this.state = {};

    }

    render() {
        const {title, create_time, pv, short_content, picture} = this.props
        return (
            <Row>
                <Col span={4}>{picture}</Col>
                <Col span={20}>
                    <Row><Col span={20}>{title}</Col><Col span={4}>{create_time}</Col></Row>
                    <Row><Col span={24}>{short_content}</Col></Row>
                    <Row><Col span={18}/><Col span={6}>pv:{pv}</Col></Row>
                </Col>
            </Row>
        );
    }
}

ArticleListItem.propTypes = {
    id:  PropTypes.string.isRequired,
    title: PropTypes.string.isRequired,
    create_time: PropTypes.string.isRequired,
    pv: PropTypes.number.isRequired,
    short_content: PropTypes.string.isRequired,
    picture: PropTypes.string.isRequired,
};

export default ArticleListItem;