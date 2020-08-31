import React from "react";
import "./SlidePictures.less"
import {GetCarouselPicData} from "../../api/ApiCommon";
import CodeBlock from "../mark_down/CodeBlock";
import ReactMarkdown from "react-markdown";
import "github-markdown-css"
import "highlight.js/styles/github.css"
import "react-image-gallery/styles/css/image-gallery.css"
import ImageGallery from 'react-image-gallery';

class SlidePictures extends React.Component {
    constructor(props) {
        super(props);
        // id: "2b995ee9-f2e8-4dfc-b997-748b79f247a3"
        // created_at: "2020-06-20T15:26:07+08:00"
        // updated_at: "2020-06-20T15:26:07+08:00"
        // deleted_at: null
        // message: "顯示的"
        // url: "http://oss.lizo.top/static/1273910222259228672b7a47c273de0783708ea5eb52b42c35d.jpg"
        // is_on_show: true
        this.state = {data: []};
    }

    componentDidMount() {
        this.loadCarouselData();
    }

    componentWillUnmount = () => {
        this.setState = (state, callback) => {
        };
    }

    loadCarouselData = () => {
        GetCarouselPicData(data => {
            this.setState({
                data: data
            })
        })
    }

    render() {
        const {data} = this.state

        const images = data.map(val => ({
            original: val.url,
            thumbnail: val.url + "?imageView2/2/w/200/h/100",
            renderItem: () => <div key={val.id} style={{height:"250px",width:"500px"}} className="div-relative carousels">
                <img src={val.url + "?imageView2/2/w/500/h/250"} alt={"img"}/>
                <div className="div-text-area">
                    <ReactMarkdown
                        className="markdown-content-carousel-view markdown-body"
                        renderers={{code: CodeBlock,}}
                        escapeHtml={false}
                        skipHtml={false}
                        source={val.message}
                    />
                </div>
            </div>
        }));
        return (
            <div className="image-gallery">
                <ImageGallery
                    autoPlay={true}
                    showFullscreenButton={true}
                    showPlayButton={false}
                    showThumbnails={false}
                    thumbnailPosition="bottom"
                    items={images}/>
            </div>
        );
    }
}

SlidePictures.propTypes = {};

export default SlidePictures;