import React from "react";
import "./Article.less"
import {getArticleSampleList} from "../../../api/ApiCommon";
import {Button, Col, Input, List, message, Row} from "antd";

/**
 * -----------------------search key----
 * -------------------------------------
 * ------------------------------------20
 */
export default class Article extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            loading: false,
            nowPage: 1,
            pageSize: 20,
            articleList: [],
            moreButtonText: "加载更多...",
            searchKey: ""
        };
    }

    componentDidMount() {
        this.loadData(true)
    }

    componentWillUnmount = () => {
        this.setState = (state,callback)=>{};
    }

    updateSearchKey = (searchKey) => {
        this.setState({
            searchKey: searchKey
        })
    }

    loadData = (reset) => {
        let {nowPage, articleList, pageSize, searchKey} = this.state
        if (reset === true) {
            nowPage = 0;
            articleList = []
        }
        this.setState({loading: true})
        getArticleSampleList(nowPage + 1, pageSize, searchKey, (data) => {
            if (data.data.length <= 0) {
                message.info("没有更多了");
                this.setState({
                    moreButtonText: "到底了",
                    articleList: [...articleList, ...data.data],
                })
            } else {
                this.setState({
                    nowPage: nowPage + 1,
                    articleList: [...articleList, ...data.data],
                    loading: false,
                })
                if (data.data.length < pageSize) {
                    this.setState({moreButtonText: "到底了"})
                }
            }
        }, () => {
            this.setState({loading: false})
        })
    }

    render() {
        const {loading, articleList, moreButtonText} = this.state;
        const loadMore =
            (
                <div style={{
                    textAlign: 'center',
                    marginTop: 12,
                    height: 32,
                    lineHeight: '32px',
                }}>
                    <Button type="link" onClick={this.loadData}>{moreButtonText}</Button>
                </div>
            );

        return (
            <div className="article-page-view">
                <div style={{textAlign: "right", paddingTop: "10px"}}>
                    <Input name="search_key"
                           onChange={event => this.updateSearchKey(event.target.value)}
                           onPressEnter={() => this.loadData(true)}
                           style={{marginTop: "0px", margin: "0px", width: "250px"}}
                           placeholder="输入关键字并按回车进行搜索"/>
                    <div style={{marginTop: "10px"}}/>
                </div>
                <List
                    loading={loading}
                    itemLayout="horizontal"
                    loadMore={loadMore}
                    dataSource={articleList}
                    renderItem={(item, index) => {
                        const link = "/article_detail?id=" + item.id
                        return <div className="article-item-view">
                            <Row>
                                <Col span={1} style={{minWidth: "20px", color: "#3f6600", fontWeight: 700}}>
                                    <div>#{index}</div>
                                </Col>
                                <Col span={21}>
                                    <a href={link} style={{color: "#3f6600", fontWeight: 400, fontSize: "15px"}}>
                                        <Row><Col span={24}>{item.title} </Col></Row>
                                        <Row><Col span={24}>{item.description} </Col></Row>
                                    </a>
                                </Col>
                                <Col style={{textAlign: "right", color: "#3f6600", fontWeight: 400}} span={2}>
                                    <div>pv:&nbsp;{item.pv}</div>
                                </Col>
                            </Row>
                            <div style={{marginTop: "2px"}}/>
                        </div>
                    }
                    }
                />
            </div>
        );
    }

}