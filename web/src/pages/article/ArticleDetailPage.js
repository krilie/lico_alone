import React from "react";
import "./ArticleDetailPage.less"
import {getArticleById} from "../../api/ApiCommon";
import ReactMarkdown from "react-markdown/with-html";
import "github-markdown-css"
import "highlight.js/styles/github.css"
import CodeBlock from "../../components/mark_down/CodeBlock";
import IcpLabel from "../../components/icp/IcpLabel";

export default class ArticleDetailPage extends React.Component {

    constructor(props) {
        super(props);
        console.log(props);
        const searchParams = new URLSearchParams(props.location.search);
        this.state = {
            articleId: searchParams.get("id"),
            article: {}
        }
    }

    // id, created_at, updated_at, deleted_at, title, pv, content, picture, description, sort
    componentWillMount() {
        const {articleId} = this.state
        getArticleById(articleId, (data) => {
            this.setState({
                article: data
            })
            document.title = data.title // 设置标题
        })
    }

    render() {
        const {articleId, article} = this.state
        if (article === undefined) {
            return <div> {articleId} wait...</div>
        } else {
            return (
                <div className="article-detail"
                     style={{padding: "20px", maxWidth: "1000px", textAlign: "center", margin: "auto"}}>
                    <ReactMarkdown className="markdown-body markdown-content"
                                   renderers={{code: CodeBlock,}}
                                   escapeHtml={false}
                                   skipHtml={false}
                                   source={article.content}
                    />
                    <div className="article-foot-info">
                        {article.title}&nbsp;{article.created_at}&nbsp;共访问{article.pv}次
                    </div>
                    <IcpLabel/>
                </div>
            );
        }
    }
}
