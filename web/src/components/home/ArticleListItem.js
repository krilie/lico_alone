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

    render() {
        const {title, description, pv, id, picture} = this.props
        const url = "/article_detail?id=" + id;
        return (
            <Card className="article-item-card" bodyStyle={{padding: "0", margin: "0"}} style={{minWidth: 300}}>
                <Layout className="article-layout">

                    <Layout.Sider width={100} className="article-layout-sider">
                        <img className="pic" src={picture} alt={"pic"}/>
                    </Layout.Sider>

                    <Layout.Content className="article-layout-content">
                        <Layout className="article-layout">

                            <Layout.Header style={{minHeight: "30px", maxHeight: "30px"}}
                                           className="article-layout-content-real">
                                <a href={url} rel="noopener noreferrer" target="_blank">
                                    <div style={{height: "unset", verticalAlign: "center"}}>
                                        {title}
                                    </div>
                                </a>
                            </Layout.Header>

                            <Layout.Content>
                                <a href={url} rel="noopener noreferrer" target="_blank"
                                   className="article-description-view">{description}</a>
                            </Layout.Content>

                            <Layout.Footer className="article-layout-footer">
                                <div style={{textAlign: "left", fontWeight: "600"}}>
                                    访问量:&nbsp;{pv}&nbsp;次
                                </div>
                            </Layout.Footer>

                        </Layout>
                    </Layout.Content>
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
