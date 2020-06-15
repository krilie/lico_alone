import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {Form, Input, Modal, Select} from "antd";

class CarouselCreateUpdateModal extends Component {

    formRef = React.createRef();

    onDialogOk(isCreate) {
        const {success} = this.props
        success();
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
       process.nextTick(()=>{
           let {isShow, isCreate, data} = this.props
           if (isShow) {
               if (this.formRef.current !== null) {
                   if (isCreate) {
                       this.formRef.current.setFieldsValue({
                           ...data
                       })
                   } else {
                       this.formRef.current.setFieldsValue({
                           ...data
                       })
                   }
               }
           }
       })
    }

    render() {
        const {isShow, cancel, isCreate} = this.props
        const titleText = isCreate ? "创建" : "修改"

        const layout = {
            labelCol: {span: 4},
            wrapperCol: {span: 20}
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
                    <Input defaultValue="" placeholder="请输入"/>
                </Form.Item>
                <Form.Item label="图址" name="url">
                    <Input defaultValue="" placeholder="请输入"/>
                </Form.Item>
            </Form>

        return (
            <Modal
                title={titleText}
                visible={isShow}
                onOk={() => this.onDialogOk(isCreate)}
                onCancel={() => cancel()}>
                {formShow}
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