import React, {Component} from 'react';
import "./FilePage.less"
import {Card, message, Pagination, Table, Tag} from "antd";
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
        manageGetFilePage().then(res => {
            debugger
            this.setState({
                files: {...res.data}
            })
        }).catch(err => {
            message.warning(err)
        }).finally(() => {
            this.setState({
                loading: false
            })
        })
    }

    componentWillMount() {
        this.loadData(1,10)
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
                <Table
                    pagination={pagination}
                    loading={loading}
                    columns={columns}
                    dataSource={data}/>
            </Card>
        );
    }
}

export default FilePage;


const columns = [
    {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
        render: text => <div>{text}</div>,
    },
    {
        title: 'Age',
        dataIndex: 'age',
        key: 'age',
    },
    {
        title: 'Address',
        dataIndex: 'address',
        key: 'address',
    },
    {
        title: 'Tags',
        key: 'tags',
        dataIndex: 'tags',
        render: tags => (
            <>
                {tags.map(tag => {
                    let color = tag.length > 5 ? 'geekblue' : 'green';
                    if (tag === 'loser') {
                        color = 'volcano';
                    }
                    return (
                        <Tag color={color} key={tag}>
                            {tag.toUpperCase()}
                        </Tag>
                    );
                })}
            </>
        ),
    },
    {
        title: 'Action',
        key: 'action',
        render: (text, record) => (
            <div>
                <div>Invite {record.name}</div>
                <div>Delete</div>
            </div>
        ),
    },
];