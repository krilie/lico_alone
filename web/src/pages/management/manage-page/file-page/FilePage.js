import React, {Component} from 'react';
import "./FilePage.less"
import {Button, Card, Col, message, Modal, Row, Table, Upload} from "antd";
import {manageDeleteFile, manageGetFilePage} from "../../../../api/ManageFileApi";
import UploadOutlined from "@ant-design/icons/lib/icons/UploadOutlined";
import {GetUserToken} from "../../../../utils/LocalStorageUtil";
import {apiBaseUrl} from "../../../../api/ApiBaseUrl";
import CopyToBoard from "../../../../utils/CopyToBoard";

class FilePage extends Component {

    columns = [
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
            render: text => <img src={text+"?imageView2/2/w/200/h/100"} alt={"img"}/>
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
        {
            title: '操作',
            key: 'operation',
            render: file =>
                <div className="table-file-operator">
                    <Row>
                        <Col style={{textAlign:"center",margin:"2px"}} span={24}>
                            <Button onClick={() => CopyToBoard(file.url)}>复制地址</Button>
                        </Col>
                    </Row>
                    <Row>
                        <Col style={{textAlign:"center",margin:"2px"}} span={24}>
                            <Button onClick={() => this.deleteFileItem(file.id)}>删除</Button>
                        </Col>
                    </Row>
                </div>
        }
    ];

    constructor(props) {
        super(props);
        this.state = {
            loading: true,
            files: {
                page_info: {total_count: 0, total_page: 0, page_num: 1, page_size: 2},
                data: []
            },
            uploadModal: {
                show: false,
            }
        }
    }

    uploadFileModalSuccess = e => {
        this.uploadFileModalSetShow(false)
    }
    uploadFileModalCancel = e => {
        this.uploadFileModalSetShow(false)
    }
    uploadFileModalSetShow = (show) => {
        this.setState({
            uploadModal: {
                show: show
            }
        })
    }

    deleteFileItem = (id) => {
        manageDeleteFile(id).then(res => {
            message.info("delete success")
            this.reloadFileItems()
        })
    }

    // 加载数据
    loadFileItems = (page_num, page_size) => {
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

    reloadFileItems = () => {
        const {page_num, page_size} = this.state.files.page_info
        this.loadFileItems(page_num, page_size)
    }

    uploadFileProps = {
        name: 'file',
        action: `${apiBaseUrl}/api/manage/file/upload`,
        headers: {
            authorization: GetUserToken()
        },
        defaultFileList: false,
        showUploadList: false,
        onChange: (info) => {
            if (info.file.status !== 'uploading') {
                console.log(info.file, info.fileList);
            }
            if (info.file.status === 'done') {
                message.success(`${info.file.name} file uploaded successfully`);
                this.reloadFileItems()
            } else if (info.file.status === 'error') {
                message.error(`${info.file.name} file upload failed.`);
            }
        },
        progress: {
            strokeColor: {
                '0%': '#108ee9',
                '100%': '#87d068',
            },
            strokeWidth: 3,
            format: percent => `${parseFloat(percent.toFixed(2))}%`,
        },
    };


    componentWillMount() {
        this.loadFileItems(1, 2)
    }

    // 分页修改当前页大小 回调
    onLoadPageData = (page_num, page_size) => {
        console.log(page_num, page_size);
        this.loadFileItems(page_num, page_size);
    }

    render() {
        const {data} = this.state.files
        const {page_num} = this.state.files.page_info
        const {page_size} = this.state.files.page_info
        const {loading} = this.state
        return (
            <Card bodyStyle={{padding: "10px"}}>
                <Button type={"primary"} onClick={() => this.uploadFileModalSetShow(true)}>添加</Button>
                <Upload{...this.uploadFileProps}>
                    <Button>
                        <UploadOutlined/> 上传文件
                    </Button>
                </Upload>
                <div className="table">
                    <Table
                        bordered
                        pagination={{
                            current: page_num,
                            pageSize: page_size,
                            defaultCurrent: 1,
                            defaultPageSize: 10,
                            position: "buttom"
                        }}
                        loading={loading}
                        columns={this.columns}
                        dataSource={data}/>
                </div>

                <Modal
                    title="Basic Modal"
                    visible={this.state.uploadModal.show}
                    onOk={this.uploadFileModalSuccess}
                    onCancel={this.uploadFileModalCancel}
                >
                    <p>Some contents...</p>
                    <p>Some contents...</p>
                    <p>Some contents...</p>
                </Modal>

            </Card>
        );
    }
}

export default FilePage;