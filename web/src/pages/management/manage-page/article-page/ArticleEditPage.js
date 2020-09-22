import React, {Component} from 'react';
import {manageCreateArticle, manageGetArticleById, manageUpdateArticle} from "../../../../api/ManageAritcleApi";
import {Button, Card, Col, Form, Input, InputNumber, message, Row} from "antd";
import "./ArticleEditPage.less"
import TextArea from "antd/lib/input/TextArea";
import AutoCompleteImageFile from "../../../../components/imageFile/AutoCompleteImageFile";
import ArrowLeftOutlined from "@ant-design/icons/lib/icons/ArrowLeftOutlined";
import CodeBlock from "../../../../components/mark_down/CodeBlock";
import ReactMarkdown from "react-markdown/with-html";
import SelectFileModal from "../../../../components/imageFile/SelectFileModal";
import "github-markdown-css"
import "highlight.js/styles/github.css"

class ArticleEditPage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            isCreate: false,
            showMarkDownView: true,
            currentContent: "",
            articleId: props.match.params.articleId,
            waiting: false,
            article: {
                id: "",
                created_at: "2020-05-27T20:23:20+08:00",
                updated_at: "2020-05-27T20:23:20+08:00",
                deleted_at: null,
                title: "",
                pv: 0,
                content: "",
                picture: "",
                description: "",
                sort: 0,
            },
            selectFileModal: {
                isShow: false,
                onOk: () => {
                    this.setState({selectFileModal: {...this.state.selectFileModal,isShow: false}})
                },
                onCancel: () => {
                    this.setState({selectFileModal: {...this.state.selectFileModal,isShow: false}})
                },
            }
        }
    }

    showSelectFileModal = () => {
        this.setState({selectFileModal: {...this.state.selectFileModal,isShow: true}})
    }

    componentDidMount() {

        let articleId = this.props.match.params.articleId
        if (articleId === undefined || articleId === {} || articleId === "") {
            this.setState({isCreate: true, articleId: ""})
        } else {
            manageGetArticleById(articleId).then(res => {
                this.setState({article: res.data.data, currentContent: res.data.data.content})
                this.formRef.current.setFieldsValue({...res.data.data})
            }).catch(err => {
                message.error(err.toString())
            })
            this.setState({isCreate: false, articleId: articleId})
        }

    }

    createArticle = () => {
        const data = {
            id: this.formRef.current.getFieldValue("id"),
            title: this.formRef.current.getFieldValue("title"),
            content: this.formRef.current.getFieldValue("content"),
            picture: this.formRef.current.getFieldValue("picture"),
            description: this.formRef.current.getFieldValue("description"),
            sort: this.formRef.current.getFieldValue("sort"),
        }
        this.setState({waiting: true})
        manageCreateArticle(data).then(res => {
            message.info(res.data.message)
            this.props.history.push("/management/manage/article")
        }).catch(err => {
            message.error(err.toString())
        }).finally(() => {
            this.setState({waiting: false})
        })
    }

    updateArticle = () => {
        const data = {
            id: this.formRef.current.getFieldValue("id"),
            title: this.formRef.current.getFieldValue("title"),
            content: this.formRef.current.getFieldValue("content"),
            picture: this.formRef.current.getFieldValue("picture"),
            description: this.formRef.current.getFieldValue("description"),
            sort: this.formRef.current.getFieldValue("sort"),
        }
        this.setState({waiting: true})
        manageUpdateArticle(data).then(res => {
            message.info(res.data.message)
        }).catch(err => {
            message.error(err.toString())
        }).finally(() => {
            this.setState({waiting: false})
        })
    }

    formRef = React.createRef();

    render() {
        const {isCreate, article, showMarkDownView, currentContent, selectFileModal} = this.state
        const titleVal = <div>
            <Row justify="start">
                <Col flex="30px"><ArrowLeftOutlined style={{color: "#eb2f96", fontSize: "20px"}}
                                                    onClick={() => this.props.history.goBack()}/></Col>
                <Col flex="auto"> {isCreate ? <div className="fit-content">创建</div> :
                    <div className="fit-content">修改: {article.id}</div>}</Col>
                <Col flex="auto">
                    <Button type={"primary"} onClick={() => this.showSelectFileModal()}>选择图片</Button>
                </Col>
            </Row>
        </div>

        const contentView = showMarkDownView ?
            <Row>
                <Col span={11}>
                    <Form.Item name="content_markdown">
                        <ReactMarkdown className="markdown-body markdown-content" renderers={{code: CodeBlock,}}
                                       escapeHtml={false} skipHtml={false} source={currentContent}/>
                    </Form.Item>
                </Col>
                <Col span={12}>
                    <Form.Item name="content">
                        <TextArea onChange={(event) => this.setState({currentContent: event.target.value})}
                                  autoSize={{maxRows: 60, minRows: 6}} defaultValue="" placeholder="请输入"/>
                    </Form.Item>
                </Col>
            </Row>
            : <Row>
                <Col span={24}>
                    <Form.Item name="content">
                        <TextArea autoSize={{maxRows: 60, minRows: 6}} defaultValue="" placeholder="请输入"/>
                    </Form.Item>
                </Col>
            </Row>

        // 查询form参数
        const createForm = <Form
            ref={this.formRef}
        >
            <Row>
                <Col span={10}>
                    <Form.Item label="标题" name="title">
                        <Input defaultValue="" placeholder="请输入"/>
                    </Form.Item>
                </Col>
                <Col span={8}>
                    <Form.Item label="图片" name="picture">
                        <AutoCompleteImageFile/>
                    </Form.Item>
                </Col>
                <Col span={6}>
                    <Form.Item label="排序" name="sort">
                        <InputNumber defaultValue={0} value={0}/>
                    </Form.Item>
                </Col>
            </Row>
            <Row>
                <Col span={24}>
                    <Form.Item label="描述" name="description">
                        <TextArea defaultValue="" placeholder="请输入"/>
                    </Form.Item>
                </Col>
            </Row>
            {contentView}
            <Row>
                <Col span={24}>
                    <div style={{textAlign: "left"}}>
                        <Form.Item style={{textAlign: "left"}}>
                            <Button type="primary" style={{textAlign: "right"}} onClick={() => {
                                isCreate ? this.createArticle() : this.updateArticle()
                            }}>{isCreate ? "创建" : "更新"}</Button>
                        </Form.Item>
                    </div>
                </Col>
            </Row>
        </Form>


        return (
            <Card className="carousel-card" bodyStyle={{padding: "10px"}} title={titleVal}>
                <Row><Col span={24} className="search-form">{createForm}</Col></Row>
                <SelectFileModal {...selectFileModal} />
            </Card>
        );
    }
}

export default ArticleEditPage;