import React from "react";
import "./Article.less"
import {getArticleById} from "../../api/common";
import ReactMarkdown from "react-markdown";
import CodeBlock from "../../components/mark_down/CodeBlock";

const toc = require('remark-toc')

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
                    <ReactMarkdown className="markdown-content"
                                   skipHtml={true}
                                   renderers={{code: CodeBlock}}
                                   plugins={[toc]}
                                   escapeHtml={false}
                                   source={article.content}
                    />
                </div>
            );
        }
    }
}
