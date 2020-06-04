import React from "react";
import PropTypes from 'prop-types';
import "./ArticleListItem.less"
import {Card} from "antd";

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
            <Card className="article-item-card" bodyStyle={{padding:"0 0 0 0",margin:"0 0 0 0"}} style={{ minWidth: 400 }}>

                <p>{title}</p>
                <p>{description}</p>
                <p>{create_time}</p>
                <p>{pv}</p>
                <p>{short_content}</p>
                <p>{picture}</p>
            </Card>
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