import React, {Component} from 'react';
import {AutoComplete, Button, message, Spin} from "antd";
import {manageGetFilePage} from "../../api/ManageFileApi";
import {replaceForImageProxy} from "../../utils/ImageProxy";

class AutoCompleteImageFile extends Component {

    constructor(props) {
        super(props);
        this.state = {imgList: [], imgPage: 0, imgLoading: false}
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
        manageGetFilePage({page_num: imgPage + 1, page_size: 3}).then(res => {
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
        const {style = {}, ...otherProps} = this.props;
        const {imgList, imgLoading} = this.state

        const options =
            imgList.map(val => {
                return {
                    value: val.url,
                    label: <img src={replaceForImageProxy(val.url,"150x")} height="150px" width="auto" alt={"img"}/>
                }
            });
        options.push({
            value: "",
            label: imgLoading ? <Spin/> : <div>
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
        })

        return (
            <AutoComplete
                style={{width: "100%", ...{style}}}
                options={options}
                {...otherProps}
                onFocus={() => this.appendImgList(true)}
            >
                {this.props.children}
            </AutoComplete>
        );
    }
}

export default AutoCompleteImageFile;