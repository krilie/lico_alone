import React from "react";
import "./Article.less"
import ArticleListPageRollView from "../../components/home/ArticleListPageRollView";

export default class Article extends React.Component {
    render() {
        return (
            <div style={{padding:"20px",maxWidth:"800px",textAlign:"center",margin:"auto"}}>
                <ArticleListPageRollView/>
            </div>
        );
    }
}