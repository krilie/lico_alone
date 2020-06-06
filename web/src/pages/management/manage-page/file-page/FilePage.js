import React, {Component} from 'react';
import "./FilePage.less"
import {Card, Table, Tag} from "antd";

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







    render() {
        return (
            <Card bodyStyle={{padding:"10px"}}>

                <Table
                    columns={columns}

                    dataSource={data} />
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
            <div >
                <div>Invite {record.name}</div>
                <div>Delete</div>
            </div>
        ),
    },
];

const data = [
    {
        key: '1',
        name: 'John Brown',
        age: 32,
        address: 'New York No. 1 Lake Park',
        tags: ['nice', 'developer'],
    },
    {
        key: '2',
        name: 'Jim Green',
        age: 42,
        address: 'London No. 1 Lake Park',
        tags: ['loser'],
    },
    {
        key: '3',
        name: 'Joe Black',
        age: 32,
        address: 'Sidney No. 1 Lake Park',
        tags: ['cool', 'teacher'],
    },
];