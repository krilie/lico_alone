import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {Form, Input, message, Modal, Select, Spin} from "antd";
import {manageCreateCarousel, manageUpdateCarousel} from "../../../../api/ManageCarouselApi";
import TextArea from "antd/lib/input/TextArea";
import AutoCompleteImageFile from "../../../../components/imageFile/AutoCompleteImageFile";

class CarouselCreateUpdateModal extends Component {
    // 构造函数
    constructor(props) {
        super(props);
        this.state = {waiting: false}
    }

    // 表单
    formRef = React.createRef();

    // 在Dialog点击ok时
    onDialogOk = (isCreate) => {
        if (isCreate) {
            this.createCarousel()
        } else {
            this.updateCarousel()
        }
    }

    // 更新时
    componentDidUpdate(prevProps, prevState, snapshot) {
        process.nextTick(() => {
            let {isShow, isCreate, data} = this.props
            if (isShow) {
                if (this.formRef.current !== null) {
                    if (isCreate) {
                        this.formRef.current.setFieldsValue({...data})
                    } else {
                        this.formRef.current.setFieldsValue({...data})
                    }
                }
            }
        })
    }

    // 更新Carousel
    updateCarousel = () => {
        const {success} = this.props
        const data = {
            id: this.formRef.current.getFieldValue("id"),
            is_on_show: this.formRef.current.getFieldValue("is_on_show"),
            message: this.formRef.current.getFieldValue("message"),
            url: this.formRef.current.getFieldValue("url"),
        }
        this.setState({waiting: true})
        manageUpdateCarousel(data).then(res => {
            message.info("update ok")
            success();
        }).finally(() => {
            this.setState({waiting: false})
        })

    }

    // 创建
    createCarousel = () => {
        const {success} = this.props
        const data = {
            is_on_show: this.formRef.current.getFieldValue("is_on_show"),
            message: this.formRef.current.getFieldValue("message"),
            url: this.formRef.current.getFieldValue("url"),
        }
        this.setState({waiting: true})
        manageCreateCarousel(data).then(res => {
            message.info("create ok")
            success();
        }).finally(() => {
            this.setState({waiting: false})
        })
    }

    render() {
        const {isShow, cancel, isCreate} = this.props
        const {waiting} = this.state
        const titleText = isCreate ? "创建" : "修改"

        const layout = {
            labelCol: {span: 3},
            wrapperCol: {span: 21}
        };

        const formShow =
            <Form
                {...layout}
                ref={this.formRef}
            >
                <div hidden={isCreate}>
                    <Form.Item label="键" name="id">
                        <Input disabled={true} defaultValue="" placeholder="请输入"/>
                    </Form.Item>
                </div>
                <Form.Item label="显示" name="is_on_show">
                    <Select
                        defaultValue={true}
                        placeholder="Select a option and change input text above"
                    >
                        <Select.Option value={true}>true</Select.Option>
                        <Select.Option value={false}>false</Select.Option>
                    </Select>
                </Form.Item>
                <Form.Item label="信息" name="message">
                    <TextArea aria-multiline={"true"} defaultValue="" placeholder="请输入"/>
                </Form.Item>
                <Form.Item label="图址" name="url">
                    <AutoCompleteImageFile/>
                </Form.Item>
            </Form>

        return (
            <Modal
                title={titleText}
                visible={isShow}
                maskClosable={false}
                onOk={() => this.onDialogOk(isCreate)}
                onCancel={() => cancel()}>
                {waiting ? <Spin/> : formShow}
            </Modal>
        );

    }
}

CarouselCreateUpdateModal.propTypes = {
    data: PropTypes.object,
    isShow: PropTypes.bool,
    success: PropTypes.func.isRequired,
    cancel: PropTypes.func.isRequired,
};

export default CarouselCreateUpdateModal;