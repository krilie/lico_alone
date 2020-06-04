import React from "react";
import PropTypes from 'prop-types';
import "./ArticleListItem.less"
import {Card, Layout, message} from "antd";

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

    goDetailPage = (articleId) => {
        message.info("go to article" + articleId)
    }

    render() {
        const {title, description, create_time, pv, id, picture} = this.props
        return (
            <Card className="article-item-card"
                  bodyStyle={{padding: "0 0 0 0", margin: "0 0 0 0"}}
                  style={{minWidth: 400}}>
                <Layout className="article-layout">
                    <Layout.Header onClick={() => this.goDetailPage(id)} className="article-layout-header">
                        <div style={{height: "48px", verticalAlign: "center"}}>{title}</div>
                    </Layout.Header>
                    <Layout className="article-layout">
                        <Layout.Sider width={100} className="article-layout-sider">
                            <img className="pic" src={picture} alt={"pic"}/>
                        </Layout.Sider>
                        <Layout.Content className="article-layout-content">{description}</Layout.Content>
                    </Layout>
                    <Layout.Footer className="article-layout-footer">create_time:{create_time} pv:{pv}</Layout.Footer>
                </Layout>
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