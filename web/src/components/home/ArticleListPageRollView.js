import React from "react";
import {List, Button, message} from 'antd';
import "./ArticleListPageRollView.less"
import {getArticleSampleList} from "../../api/common";
import ArticleListItem from "./ArticleListItem";

const pageSize = 7;

class ArticleListPageRollView extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            initLoading: true,
            loading: false,
            nowPage: 1,
            articleList: [],
            moreButtonText: "加载更多..."
        };

    }

    componentDidMount() {
        this.getData(1)
        this.setState({
            initLoading: false,
        })
    }

    getData = (pageNum) => {
        this.setState({loading: true})
        getArticleSampleList(pageNum, pageSize, "", (data) => {
            if (data.data.length <= 0) {
                message.info("没有更多了");
                this.setState({
                    moreButtonText: "到底了"
                })
            } else {
                this.setState({
                    nowPage: pageNum,
                    articleList: [...this.state.articleList, ...data.data],
                    loading: false,
                })
                if (data.data.length < pageSize) {
                    this.setState({
                        moreButtonText: "到底了"
                    })
                }
            }
        }, () => {
            this.setState({
                loading: false,
            })
        })
    };

    onLoadMore = () => {
        const {nowPage} = this.state
        this.getData(nowPage + 1)
    };

    render() {
        const {initLoading, loading, articleList, moreButtonText} = this.state;
        const loadMore =
            !initLoading && !loading ? (
                <div
                    style={{
                        textAlign: 'center',
                        marginTop: 12,
                        height: 32,
                        lineHeight: '32px',
                    }}
                >
                    <Button type="link" onClick={this.onLoadMore}>{moreButtonText}</Button>
                </div>
            ) : null;

        return (
            <List
                className="demo-loadmore-list"
                loading={initLoading}
                itemLayout="horizontal"
                loadMore={loadMore}
                dataSource={articleList}
                renderItem={item => (
                    <ArticleListItem
                        id={item.title}
                        title={item.title}
                        create_time={item.title}
                        pv={item.title}
                        short_content={item.title}
                        picture="https://pic1.zhimg.com/80/v2-af6f3a9444c74d726c63ed5291f9e53d_720w.jpg"
                        description={item.title}/>
                )}
            />
        );
    }

}

export default ArticleListPageRollView;