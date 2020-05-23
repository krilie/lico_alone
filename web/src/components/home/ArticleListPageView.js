import React from "react";
import "./ArticleListView.less"
import PropTypes from "prop-types";
import ArticleListItem from "./ArticleListItem";

class ArticleListPageView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            loading: false,
            articles: []
        };

    }

    componentDidMount() {
        // 请求获取数据

    }

    render() {
        const {articles} = this.state
        const articlesItems = articles.map(val => <ArticleListItem {...val}/>);
        return ({articlesItems});
    }
}

ArticleListPageView.stateTypes = {
    page_index: PropTypes.number.isRequired,
    page_size: PropTypes.number.isRequired,
    // articles: PropTypes.arrayOf(PropTypes.shape({
    //     id: PropTypes.string.isRequired,
    //     title: PropTypes.string.isRequired,
    //     create_time: PropTypes.string.isRequired,
    //     pv: PropTypes.number.isRequired,
    //     short_content: PropTypes.string.isRequired,
    //     picture: PropTypes.string.isRequired,
    // }))
};
ArticleListPageView.defaultProps = {
    page_index: 1,
    page_size: 10,
}

export default ArticleListPageView;