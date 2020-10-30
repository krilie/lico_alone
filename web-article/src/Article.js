import './Article.css';
import React from 'react';
import {checkIsNotFound, checkResDataWithToast, getArticleById} from "./api/ApiCommon";
import IcpLabel from "./components/IcpLabel";
import ReactMarkdown from "react-markdown";
import "highlight.js/styles/github.css"
import NotFoundBox from "./components/NotFoundBox";
import CodeBlock from "./components/CodeBlock";

export default class Article extends React.Component {

    constructor(props) {
        super(props);
        const searchParams = new URLSearchParams(props.location.search);
        let id = searchParams.get("id")
        this.state = {
            articleId: id,
            article: id === undefined ? undefined : {}
        }
    }

    // id, created_at, updated_at, deleted_at, title, pv, content, picture, description, sort
    componentWillMount() {
        const {articleId} = this.state
        if (articleId === undefined)
            return
        getArticleById(articleId).then(res => {
            if (checkIsNotFound(res)) {
                this.setState({
                    article: undefined
                })
            } else {
                let data = checkResDataWithToast(res)
                if (data) {
                    this.setState({
                        article: data
                    })
                    document.title = data.title // 设置标题
                }
            }
        })
    }

    buildContent = () => {
        const {articleId, article} = this.state
        if (article === {}) {
            return <div>Loading...</div>
        } else if (article === undefined) {
            return <NotFoundBox msg={articleId}/>
        } else {
            return (
                <div>
                    <ReactMarkdown
                        renderers={{code: CodeBlock,}}
                        escapeHtml={false}
                        skipHtml={false}
                        source={article.content}
                    />
                    <div className="article-foot-info">
                        {articleId}-{article.title}&nbsp;{article.created_at}&nbsp;共访问{article.pv}次
                    </div>
                </div>
            );
        }
    }

    render() {
        return <div className="App">
            <div className="Article">
                {this.buildContent()}
            </div>
            <IcpLabel/>
        </div>
    }
}
