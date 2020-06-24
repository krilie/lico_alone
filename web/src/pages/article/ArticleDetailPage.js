import React from "react";
import "./Article.less"
import {getArticleById} from "../../api/ApiCommon";
import ReactMarkdown from "react-markdown/with-html";
import "github-markdown-css"
import "highlight.js/styles/github.css"
import CodeBlock from "../../components/mark_down/CodeBlock";

export default class ArticleDetailPage extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            articleId: props.match.params.articleId,
            article: {}
        }
    }

    // id, created_at, updated_at, deleted_at, title, pv, content, picture, description, sort
    componentWillMount() {
        const articleId = this.props.match.params.articleId
        this.setState({
            articleId: articleId
        })
        getArticleById(articleId, (data) => {
            this.setState({
                article: data
            })
        })

    }

    render() {
        const {articleId, article} = this.state
        if (article === undefined) {
            return <div> {articleId} wait...</div>
        } else {
            return (
                <div style={{ padding: "20px", maxWidth: "1000px", textAlign: "center", margin: "auto"}}>

                    <ReactMarkdown className="markdown-body markdown-content"
                                   renderers={{
                                       code: CodeBlock,
                                       // heading: HeadingBlock
                                   }}
                                   escapeHtml={false}
                                   skipHtml={false}
                                   source={article.content}
                    />
                    <div style={{textAlign:"right",color:"#ff9900"}}>
                        标题:&nbsp;{article.title}&nbsp;&nbsp;创建时间:&nbsp;{article.created_at}&nbsp;共访问{article.pv}次
                    </div>
                </div>
            );
        }
    }
}
