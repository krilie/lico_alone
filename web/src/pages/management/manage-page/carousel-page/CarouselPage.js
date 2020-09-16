import React, {Component} from 'react';
import {manageDeleteCarouselById, manageGetCarouselList} from "../../../../api/ManageCarouselApi";
import {Button, Card, Col, message, Row, Table} from "antd";
import "./CarouselPage.less"
import "../../../../utils/common.less"
import CarouselCreateUpdateModal from "./CarouselCreateUpdateModal";
import {replaceForImageProxy} from "../../../../utils/ImageProxy";

class CarouselPage extends Component {

    goToPage = path => this.props.history.push(path);

    constructor(props) {
        super(props);
        this.state = {
            carousels: [],
            loading: false,
            createUpdateDialog: {
                isShow: false,
                isCreate: true,
                success: this.createUpdateDialogSuccessClick,
                cancel: this.createUpdateDialogCancelClick,
            }
        }
    }

    showCreateCarouselDialog = () => {
        const {createUpdateDialog} = this.state
        this.setState({
            createUpdateDialog: {
                ...createUpdateDialog,
                data: {id:"",is_on_show: false,message:"",url:""},
                isCreate: true,
                isShow: true,
            }
        })
    }

    showUpdateCarouselDialog = (item) => {
        const {createUpdateDialog} = this.state
        this.setState({
            createUpdateDialog: {
                ...createUpdateDialog,
                data: {...item},
                isCreate: false,
                isShow: true,
            }
        })
    }


    createUpdateDialogSuccessClick = () => {
        const {createUpdateDialog} = this.state
        this.loadCarouselData()
        this.setState({
            createUpdateDialog: {
                ...createUpdateDialog,
                isShow: false,
            }
        })
    }
    createUpdateDialogCancelClick = () => {
        const {createUpdateDialog} = this.state
        this.setState({
            createUpdateDialog: {
                ...createUpdateDialog,
                isShow: false,
            }
        })
    }

    componentDidMount() {
        this.loadCarouselData()
    }

    // 加载数据
    loadCarouselData = () => {
        this.setState({loading: true})
        manageGetCarouselList().then(res => {
            this.setState({
                carousels: res.data.data
            })
        }).catch(err => {
            message.error(err)
        }).finally(() => {
            this.setState({loading: false})
        })
    }

    deleteCarouselById = (id) => {
        this.setState({loading: true})
        manageDeleteCarouselById(id).then(res => {
            this.loadCarouselData()
        }).finally(() => {
            this.setState({loading: false})
        })
    }

    columns = [
        {title: 'ID', key: 'id', dataIndex: 'id'},
        {
            title: 'is_on_show',
            key: 'is_on_show',
            dataIndex: 'is_on_show',
            render: val => <div>{val === true ? "true" : "false"}</div>
        },
        {
            title: 'message', key: 'message', dataIndex: 'message',
            render: val => <div className="multi-line">{val}</div>
        },
        {
            title: 'url', key: 'url', dataIndex: 'url',
            render: text => <img src={replaceForImageProxy(text,"150x")} alt={"img"}/>
        },
        {
            title: '操作', key: 'operation',
            render: item =>
                <div className="table-file-operator">
                    <Row>
                        <Button style={{margin: "3px"}} type={"primary"}
                                onClick={() => this.showUpdateCarouselDialog(item)}>修改</Button>
                        <Button style={{margin: "3px"}} type={"danger"}
                                onClick={() => this.deleteCarouselById(item.id)}>删除</Button>
                    </Row>
                </div>
        }
    ]

    render() {
        const {carousels, loading, createUpdateDialog} = this.state

        const titleView = <div>
            <Row>
                <Col className="v-center" style={{textAlign: "left"}} span={12}>
                    <div>轮播图</div>
                </Col>
                <Col style={{textAlign: "right"}} span={12}>
                    <Button style={{margin: "3px"}} type={"primary"}
                            onClick={() => this.showCreateCarouselDialog()}>添加</Button>
                    <Button style={{margin: "3px"}} type={"primary"} onClick={() => this.loadCarouselData()}>刷新</Button>
                </Col>
            </Row>
        </div>

        return (
            <Card className="carousel-card carousel-table" title={titleView}>
                <CarouselCreateUpdateModal {...createUpdateDialog} />
                <Table
                    bordered
                    scroll={{x: "400px", scrollToFirstRowOnChange: true}}
                    pagination={false}
                    rowKey={record => record.id}
                    loading={loading}
                    columns={this.columns}
                    dataSource={carousels}/>
            </Card>
        );
    }
}

export default CarouselPage;