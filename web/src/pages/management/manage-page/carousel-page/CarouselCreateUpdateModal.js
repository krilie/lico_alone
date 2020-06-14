import React, {Component} from 'react';
import PropTypes from 'prop-types';
import { Form, Input, Modal} from "antd";

class CarouselCreateUpdateModal extends Component {

    formRef = React.createRef();

    onDialogOk(isCreate) {
        const {success} = this.props
        success();
    }

    render() {
        let {data, isShow, cancel,isCreate} = this.props
        const titleText = isCreate ? "创建" : "修改"
        const layout = {
            labelCol: {span: 4},
            wrapperCol: {span: 20}
        };

        const formShow = isCreate ?
            <Form
                {...layout}
                ref={this.formRef}
            >
                <Form.Item label="is_on_show" name="is_on_show">
                    <Input defaultValue="" placeholder="请输入"/>
                </Form.Item>
                <Form.Item label="message" name="message">
                    <Input defaultValue="" placeholder="请输入"/>
                </Form.Item>
                <Form.Item label="url" name="url">
                    <Input defaultValue="" placeholder="请输入"/>
                </Form.Item>
            </Form>
            :
            <Form
                {...layout}
                ref={this.formRef}
            >
                <Form.Item label="id" name="id">
                    <Input contentEditable={"false"} defaultValue={data.id} placeholder="请输入"/>
                </Form.Item>
                <Form.Item label="is_on_show" name="is_on_show">
                    <Input defaultValue={data.is_on_show} placeholder="请输入"/>
                </Form.Item>
                <Form.Item label="message" name="message">
                    <Input defaultValue={data.message} placeholder="请输入"/>
                </Form.Item>
                <Form.Item label="url" name="url">
                    <Input defaultValue={data.url} placeholder="请输入"/>
                </Form.Item>,
            </Form>

        return (
            <Modal title={titleText}
                   visible={isShow}
                   onOk={()=>this.onDialogOk(isCreate)}
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