import React, {Component} from 'react';
import "./FilePage.less"
import {Button, Card, message, Pagination, Table} from "antd";
import {manageGetFilePage} from "../../../../api/ManageFileApi";

class FilePage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            loading: true,
            files: {
                page_info: {total_count: 0, total_page: 0, page_num: 0, page_size: 0},
                data: []
            }
        }
    }

    // 加载数据
    loadData = (page_num, page_size) => {
        this.setState({
            loading: true
        })
        manageGetFilePage({page_num, page_size}).then(res => {
            this.setState({
                files: {...res.data.data}
            })
        }).catch(err => {
            message.warning(err.str)
        }).finally(() => {
            this.setState({
                loading: false
            })
        })
    }

    componentWillMount() {
        this.loadData(1, 10)
    }

    // 分页修改当前页大小 回调
    onLoadPageData = (page_num, page_size) => {
        console.log(page_num, page_size);
        this.loadData(page_num, page_size);
    }

    render() {
        const {data} = this.state.files
        const {page_num, total_count} = this.state.files.page_info
        // const { page_size, total_page} = this.state.files.page_info
        const {loading} = this.state
        const pagination =
            <Pagination
                showSizeChanger
                onShowSizeChange={this.onLoadPageData}
                onChange={this.onLoadPageData}
                defaultCurrent={page_num}
                defaultPageSize={7}
                total={total_count}/>
        return (
            <Card bodyStyle={{padding: "10px"}}>
                <Button type={"primary"}>添加</Button>
                <div className="table">
                    <Table
                        bordered
                        pagination={pagination}
                        loading={loading}
                        columns={columns}
                        dataSource={data}/>
                </div>
            </Card>
        );
    }
}

export default FilePage;


const columns = [
    {
        title: 'id',
        key: 'id',
        dataIndex: 'id'
    },
    {
        title: '创建时间',
        key: 'created_at',
        dataIndex: 'created_at',
    },
    {
        title: '地址',
        key: 'url',
        dataIndex: 'url',
        render: text=><img height={"100px"} src={text} alt={"img"}/>
    },
    {
        title: '用户ID',
        key: 'user_id',
        dataIndex: 'user_id'
    },
    {
        title: '大小',
        key: 'size',
        dataIndex: 'size'
    },
];