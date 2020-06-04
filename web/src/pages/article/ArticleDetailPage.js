import React from "react";
import "./Article.less"

export default class ArticleDetailPage extends React.Component {

    constructor(props) {
        super(props);
        this.setState({
            articleId: props.match.params.articleId
        })
    }

    componentWillMount() {
        this.setState({
            articleId: this.props.match.params.articleId
        })
    }

    render() {
        const {articleId} = this.state
        return (
            <div style={{padding: "20px", maxWidth: "1000px", textAlign: "center", margin: "auto"}}>
                articleId {articleId}
            </div>
        );
    }
}