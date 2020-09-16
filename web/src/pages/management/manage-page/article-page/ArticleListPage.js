import React, {Component} from 'react';
import {manageDeleteArticle, manageQueryArticle} from "../../../../api/ManageAritcleApi";
import {Button, Card, Col, Form, Input, message, Row, Table} from "antd";
import "./ArticleListPage.less"
import {replaceForImageProxy} from "../../../../utils/ImageProxy";

class ArticleListPage extends Component {

    // 构造
    constructor(props) {
        super(props);
        this.state = {
            loading: false,
            articles: {
                page_info: {total_count: 0, total_page: 0, page_num: 1, page_size: 5},
                data: []
            },
        }
    }

    // 初始加载数据
    componentWillMount() {
        process.nextTick(() => {
            const {page_num, page_size} = this.state.articles.page_info
            this.onLoadPageData(page_num, page_size)
        })
    }

    // 加载文章列表
    loadArticleList = (pageNum, pageSize, searchKey) => {
        this.setState({loading: true})
        manageQueryArticle(searchKey, pageNum, pageSize).then(res => {
            this.setState({
                articles: res.data.data.data
            })
        }).catch(err => {
            message.error(err.toString())
        }).finally(() => {
            this.setState({loading: false})
        })
    }

    // 分页修改当前页大小 回调
    onLoadPageData = (page_num, page_size) => {
        // 获取所有可能的参数
        let searchKey = ""
        const ResetToUndefined = (val) => (val === "" || val === null) ? undefined : val
        searchKey = ResetToUndefined(this.formRef.current.getFieldValue("search_key"))
        // todo: sorter 参数
        this.loadArticleList(page_num, page_size, searchKey);
    }

    // 表格参数变化
    onTableParamChange = (pagination, filters, sorter, extra) => {
        this.sorter = sorter
        const {current, pageSize} = pagination
        this.onLoadPageData(current, pageSize)
        console.log(sorter)
    }

    // form表单按钮搜索
    formOnSearch = () => {
        const {page_num, page_size} = this.state.articles.page_info
        this.onLoadPageData(page_num, page_size)
    }
    reloadArticleList = this.formOnSearch

    deleteArticleById = (id) => {
        this.setState({loading: true})
        manageDeleteArticle(id).then(res => {
            message.info("delete ok")
            this.reloadArticleList();
        }).catch(err => {
            message.error(err.toString())
        }).finally(() => {
            this.setState({loading: false})
        })
    }

    // form
    formRef = React.createRef();

    /**
     * "id": "string",
     * "created_at": "string",
     * "updated_at": "string"
     * "deleted_at": "string",
     * "picture": "string",
     * "title": "string",
     * "content": "string",
     * "description": "string",
     * "pv": 0,
     * "sort": 0,
     */
    columns = [
        {title: 'ID', key: 'id', dataIndex: 'id'},
        {
            title: '图片', key: 'picture', dataIndex: 'picture',
            render: text => <img src={replaceForImageProxy(text,"200x")} alt={"img"}/>
        },
        {title: '标题', key: 'title', dataIndex: 'title', render: val => <div>{val}</div>},
        {
            title: '描述',
            key: 'description',
            dataIndex: 'description',
            render: val => <div className="article-layout-content-real">{val}</div>
        },
        {
            title: '内容',
            key: 'content',
            dataIndex: 'content',
            render: val => <div className="article-layout-content-real">{val}</div>
        },
        {title: 'PV', key: 'pv', dataIndex: 'pv', render: val => <div>{val}</div>},
        {title: '排序', key: 'sort', dataIndex: 'sort', render: val => <div>{val}</div>},
        {
            title: '操作', key: 'operation',
            render: article =>
                <Row>
                    <Button style={{margin: "3px"}} type={"primary"} onClick={() => {
                        this.props.history.push('/management/manage/article/edit/' + article.id)
                    }}>编辑</Button>
                    <Button style={{margin: "3px"}} type={"danger"}
                            onClick={() => this.deleteArticleById(article.id)}>删除</Button>
                </Row>
        }
    ];

    render() {
        const {page_num, page_size, total_count} = this.state.articles.page_info
        const {data} = this.state.articles
        const {loading} = this.state

        // 查询form参数
        const searchForm = <Form
            layout={"inline"}
            ref={this.formRef}
        >
            <Form.Item label="键名" name="search_key">
                <Input defaultValue="" placeholder="请输入"/>
            </Form.Item>
            <Form.Item>
                <Button type="primary" onClick={this.formOnSearch}>搜索</Button>
            </Form.Item>
        </Form>

        return (
            <Card className="carousel-card" bodyStyle={{padding: "10px"}} title={<div>
                <Row>
                    <Col className="v-center" style={{textAlign: "left"}} span={12}>
                        <div>Article</div>
                    </Col>
                    <Col style={{textAlign: "right"}} span={12}>
                        <Button style={{margin: "3px"}} type={"primary"}
                                onClick={() => {
                                    this.props.history.push('/management/manage/article/create')
                                }}>添加</Button>
                        <Button style={{margin: "3px"}} type={"primary"}
                                onClick={() => this.reloadArticleList()}>刷新</Button>
                    </Col>
                </Row>
            </div>}>

                <div style={{margin: "3px"}}>
                    <div className="search-form">{searchForm}</div>
                </div>
                <div className="table">
                    <Table
                        bordered
                        scroll={{x: "400px", scrollToFirstRowOnChange: true}}
                        pagination={{
                            showSizeChanger: true,
                            onShowSizeChange: this.onLoadPageData,
                            current: page_num,
                            pageSize: page_size,
                            defaultCurrent: 1,
                            defaultPageSize: 7,
                            position: "bottom",
                            total: total_count,
                            onChange: this.onLoadPageData,
                        }}
                        rowKey={record => record.id}
                        loading={loading}
                        columns={this.columns}
                        onChange={this.onTableParamChange}
                        dataSource={data}/>
                </div>
            </Card>
        );
    }
}


export default ArticleListPage;