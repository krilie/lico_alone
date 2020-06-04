import React from "react";
import "./Article.less"
import {getArticleById} from "../../api/common";
import ReactMarkdown from "react-markdown";

export default class ArticleDetailPage extends React.Component {

    constructor(props) {
        super(props);
        this.setState({
            articleId: props.match.params.articleId,
            article:{}
        })
    }
    // id, created_at, updated_at, deleted_at, title, pv, content, picture, description, sort
    componentWillMount() {
        const articleId = this.props.match.params.articleId
        this.setState({
            articleId:articleId
        })
        getArticleById(articleId, (data) => {
            this.setState({
                article: data
            })
        })

    }

    render() {
        var input = "# This is a header\n\nAnd this is a paragraph";
        const {articleId, article} = this.state
        if (article === undefined){
            return <div> {articleId} wait...</div>
        }else{
            return (
                <div style={{padding: "20px", maxWidth: "1000px", textAlign: "center", margin: "auto"}}>
                    <ReactMarkdown className="markdown-content" escapeHtml={false} source={article.content} />
                    <ReactMarkdown className="markdown-content" escapeHtml={false} source={input} />
                </div>
            );
        }


    }
}