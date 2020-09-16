import React, {Component} from 'react';
import {Button, message, Modal, Spin} from "antd";
import {manageGetFilePage} from "../../api/ManageFileApi";
import copyToBoard from "../../utils/CopyToBoard";
import {replaceForImageProxy} from "../../utils/ImageProxy";

class SelectFileModal extends Component {

    // 构造函数
    constructor(props) {
        super(props);
        this.state = {imgList: [], imgPage: 0, imgLoading: false}
    }

    componentDidMount() {
        this.appendImgList(true)
    }

    // 追加图片列表 用于选择
    appendImgList = (reload) => {
        let {imgPage, imgList} = this.state
        if (reload === true) {
            imgPage = 0;
            imgList = [];
        }
        this.setState({
            imgLoading: true
        })
        manageGetFilePage({page_num: imgPage + 1, page_size: 5}).then(res => {
            const {data} = res.data.data
            this.setState({
                imgList: [...imgList, ...data], // res.data = responseBody res.data.data=pageData res.data.data.data=dataList
                imgPage: data.length > 0 ? imgPage + 1 : imgPage
            })
        }).catch(err => {
            message.warning(err.str)
        }).finally(() => {
            this.setState({
                imgLoading: false
            })
        })
    }

    render() {
        const {isShow, onOk, onCancel} = this.props
        const {imgList, imgLoading} = this.state

        const options =
            imgList.map(val => {
                return (<div style={{height: "70px", width: "130px", cursor: "pointer", margin: "5px", padding: "0"}}
                             onClick={() => {
                                 copyToBoard(val.url)
                                 message.info("copy ok")
                                 onOk();
                             }}>
                    <img src={replaceForImageProxy(val.url,"70x")} height="70px" width="auto" alt={"img"}/>
                </div>)
            });
        options.push(
            imgLoading ? <Spin/> : <div>
                <Button type={"link"}
                        onClick={(event) => {
                            this.appendImgList(false);
                            if (event.stopPropagation)
                                event.stopPropagation();
                            else
                                event.cancelBubble = true;
                        }}>更多...</Button>
                <Button type={"link"}
                        onClick={(event) => {
                            this.appendImgList(true);
                            if (event.stopPropagation)
                                event.stopPropagation();
                            else
                                event.cancelBubble = true;
                        }}>重置</Button>
            </div>
        )

        return (
            <Modal
                bodyStyle={{margin: "3px", padding: "3px"}}
                title="图片 点击复制链接"
                width={"700px"}
                visible={isShow}
                maskClosable={false}
                onOk={() => onOk()}
                onCancel={() => onCancel()}>
                <div className="modal-content">
                    {options}
                </div>
            </Modal>
        );
    }
}

export default SelectFileModal;