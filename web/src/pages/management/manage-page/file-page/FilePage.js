import React, {Component} from 'react';
import "./FilePage.less"
import {Button, Card, Col, DatePicker, Form, message, Modal, Row, Table, Upload} from "antd";
import {manageDeleteFile, manageGetFilePage, manageUpdateFile} from "../../../../api/ManageFileApi";
import UploadOutlined from "@ant-design/icons/lib/icons/UploadOutlined";
import CopyToBoard from "../../../../utils/CopyToBoard";
import {Input} from 'antd';
import {replaceForImageProxy} from "../../../../utils/ImageProxy";

class FilePage extends Component {
    // 表格列信息
    columns = [
        {title: 'ID', key: 'id', dataIndex: 'id'},
        {title: '创建时间', key: 'created_at', dataIndex: 'created_at', sorter: {multiple: 1,}},
        {
            title: '地址', key: 'url', dataIndex: 'url',
            render: text => <img src={replaceForImageProxy(text,"200x")} alt={"img"}/>
        },
        {title: '用户ID', key: 'user_id', dataIndex: 'user_id'},
        {
            title: '大小', key: 'size', dataIndex: 'size', sorter: {multiple: 2,},
            render: val => {
                const showVal = val / 1024 / 1024;
                return <div>{showVal.toFixed(3)} Mb</div>
            }
        },
        {
            title: '操作', key: 'operation',
            render: file =>
                <div className="table-file-operator">
                    <Row>
                        <Col style={{textAlign: "center", margin: "2px"}} span={24}>
                            <Button type={"primary"} onClick={() => CopyToBoard(file.url)}>复制地址</Button>
                        </Col>
                    </Row>
                    <Row>
                        <Col style={{textAlign: "center", margin: "2px"}} span={24}>
                            <Button type={"danger"} onClick={() => this.deleteFileItem(file.id)}>删除</Button>
                        </Col>
                    </Row>
                </div>
        }
    ];

    constructor(props) {
        super(props);
        this.state = {
            initPageNum: 1, initPageSize: 5, loading: true,
            files: {
                page_info: {total_count: 0, total_page: 0, page_num: 1, page_size: 5},
                data: []
            },
            uploadModal: {
                show: false,
            },
            isUploading: false
        }
    }

    // 上传对话框
    uploadFileModalSuccess = e => this.uploadFileModalSetShow(false) // 对话框 文件上传成功
    uploadFileModalCancel = e => this.uploadFileModalSetShow(false)  // 对话框 文件上传对话框取消
    uploadFileModalSetShow = (show) => this.setState({uploadModal: {show: show}}) // 对话框 设置显示或隐藏

    // 根据id删除文件
    deleteFileItem = (id) => {
        manageDeleteFile(id).then(res => {
            message.info("delete success")
            this.reloadFileItems()
        })
    }

    // 基层 加载文件的分页数据
    // page_num page_size key_name_like bucket_name_like url_like user_id biz_type content_type created_at_begin created_at_end
    loadFileItems = (page_num, page_size, others) => {
        this.setState({
            loading: true
        })
        manageGetFilePage({page_num, page_size, ...others}).then(res => {
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

    // 重新加载列表
    reloadFileItems = () => {
        const {page_num, page_size} = this.state.files.page_info
        this.onLoadPageData(page_num, page_size)
    }

    uploadFile = (file) => {
        this.setState({isUploading: true})
        manageUpdateFile(file).then(res => {
            this.reloadFileItems()
        }).finally(() => {
            this.setState({isUploading: false})
        })
    }

    // 上传文件属性
    uploadFileProps = {
        defaultFileList: false,
        showUploadList: false,
        beforeUpload: file => {
            this.uploadFile(file)
            console.log(file)
            return false;
        },
    };

    // 初始加载数据
    componentWillMount() {
        process.nextTick(() => {
            const {initPageNum, initPageSize} = this.state
            this.onLoadPageData(initPageNum, initPageSize)
        })
    }

    goToPage = path => this.props.history.push(path);

    // 分页修改当前页大小 回调
    onLoadPageData = (page_num, page_size) => {
        console.log(page_num, page_size);

        // 获取所有可能的参数
        let params = {
            key_name_like: undefined,
            bucket_name_like: undefined,
            url_like: undefined,
            user_id: undefined,
            biz_type: undefined,
            content_type: undefined,
            created_at_begin: undefined,
            created_at_end: undefined
        }
        const ResetToUndefined = (val) => (val === "" || val === null) ? undefined : val

        params.key_name_like = ResetToUndefined(this.formRef.current.getFieldValue("key_name_like"))
        params.bucket_name_like = ResetToUndefined(this.formRef.current.getFieldValue("bucket_name_like"))
        params.url_like = ResetToUndefined(this.formRef.current.getFieldValue("url_like"))
        params.user_id = ResetToUndefined(this.formRef.current.getFieldValue("user_id"))
        params.biz_type = ResetToUndefined(this.formRef.current.getFieldValue("biz_type"))
        params.content_type = ResetToUndefined(this.formRef.current.getFieldValue("content_type"))
        params.created_at_begin = ResetToUndefined(this.formRef.current.getFieldValue("created_at_begin"))
        if (params.created_at_begin !== undefined) params.created_at_begin = params.created_at_begin.format() // rfc3339
        params.created_at_end = ResetToUndefined(this.formRef.current.getFieldValue("created_at_end"))
        if (params.created_at_end !== undefined) params.created_at_end = params.created_at_end.format() // rfc3339

        // todo: sorter 参数

        this.loadFileItems(page_num, page_size, params);
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
        const {page_num, page_size} = this.state.files.page_info
        this.onLoadPageData(page_num, page_size)
    }

    formRef = React.createRef();
    sorter = null

    render() {
        const {data} = this.state.files
        const {page_num, page_size, total_count} = this.state.files.page_info
        const {loading, isUploading} = this.state
        const {initPageNum, initPageSize} = this.state

        // 查询form参数
        const searchForm = <Form
            layout={"inline"}
            ref={this.formRef}
        >
            <Form.Item label="键名" name="key_name_like">
                <Input defaultValue="" placeholder="请输入"/>
            </Form.Item>
            <Form.Item label="桶名" name="bucket_name_like">
                <Input defaultValue="" placeholder="请输入"/>
            </Form.Item>
            <Form.Item label="Url" name="url_like">
                <Input defaultValue="" placeholder="请输入"/>
            </Form.Item>
            <Form.Item label="用户Id" name="user_id">
                <Input defaultValue="" placeholder="请输入"/>
            </Form.Item>
            <Form.Item label="业务类型" name="biz_type">
                <Input defaultValue="" placeholder="请输入"/>
            </Form.Item>
            <Form.Item label="文件类型" name="content_type">
                <Input defaultValue="" placeholder="请输入"/>
            </Form.Item>
            <Form.Item label="开始创建时间" name="created_at_begin">
                <DatePicker
                    format="YYYY-MM-DD HH:mm:ss"
                    disabledDate={false}
                    disabledTime={false}
                    showTime={{defaultValue: null}}
                />
            </Form.Item>
            <Form.Item label="结束创建时间" name="created_at_end">
                <DatePicker
                    format="YYYY-MM-DD HH:mm:ss"
                    disabledDate={false}
                    disabledTime={false}
                    showTime={{defaultValue: null}}
                />
            </Form.Item>
            <Form.Item>
                <Button type="primary" onClick={this.formOnSearch}>搜索</Button>
            </Form.Item>
        </Form>

        return (
            <Card bodyStyle={{padding: "10px"}} title={<div>文件</div>}>

                <div style={{margin: "3px"}}>
                    <div className="search-form">{searchForm}</div>
                    <Button style={{margin: "3px"}}
                            type={"primary"}
                            onClick={() => this.uploadFileModalSetShow(true)}>添加</Button>
                    <Upload {...this.uploadFileProps}>
                        <Button disabled={isUploading}><UploadOutlined/>{isUploading ? "上传中..." : "点击上传"}</Button>
                    </Upload>
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
                            defaultCurrent: initPageNum,
                            defaultPageSize: initPageSize,
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

                <Modal title="Basic Modal" visible={this.state.uploadModal.show} onOk={this.uploadFileModalSuccess}
                       onCancel={this.uploadFileModalCancel}>
                    <p>Some contents...</p>
                </Modal>

            </Card>
        );
    }
}

export default FilePage;