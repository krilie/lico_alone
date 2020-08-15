import React from "react";
import PropTypes from 'prop-types';
import "./ArticleListItem.less"
import {Card, Layout} from "antd";
import {withRouter} from "react-router-dom";

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
        this.props.history.push("/article/" + articleId);
    }

    render() {
        const {title, description, pv, id, picture} = this.props
        const url = "/article_detail/" + id;
        return (
            <Card className="article-item-card"
                  bodyStyle={{padding: "0 0 0 0", margin: "0 0 0 0"}}
                  style={{minWidth: 400}}>
                <Layout className="article-layout">
                    <Layout.Header className="article-layout-header">
                        <a href={url}>
                            <div style={{height: "unset", verticalAlign: "center"}}>
                                {title}
                            </div>
                        </a>
                    </Layout.Header>
                    <Layout className="article-layout">
                        <Layout.Sider width={100} className="article-layout-sider">
                            <img className="pic" src={picture} alt={"pic"}/>
                        </Layout.Sider>
                        <Layout.Content className="article-layout-content">
                            <Layout className="article-layout">
                                <Layout.Header style={{minHeight: "30px",maxHeight:"30px"}} className="article-layout-content-real">
                                    <a href={url} className="article-description-view">{description}</a>
                                </Layout.Header>
                                <Layout.Footer className="article-layout-footer">
                                    <div style={{textAlign: "left", fontWeight: "600"}}>
                                        访问量:&nbsp;{pv}&nbsp;次
                                    </div>
                                </Layout.Footer>
                            </Layout>
                        </Layout.Content>
                    </Layout>
                </Layout>
            </Card>
        );
    }
}

ArticleListItem.propTypes = {
    id: PropTypes.string.isRequired,
    title: PropTypes.string.isRequired,
    pv: PropTypes.number.isRequired,
    picture: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
};

export default withRouter(ArticleListItem)
