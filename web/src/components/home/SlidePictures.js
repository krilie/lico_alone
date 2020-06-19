import React from "react";
import "./SlidePictures.less"
import {Carousel} from "antd";

class SlidePictures extends React.Component {
    constructor(props) {
        super(props);

        this.state = {};
    }

    render() {
        return (
            <Carousel autoplay dotPosition='bottom'>
                <div>
                    <h3>1</h3>
                </div>
                <div>
                    <h3>2</h3>
                </div>
                <div>
                    <h3>3</h3>
                </div>
                <div>
                    <h3>4</h3>
                </div>
            </Carousel>
        );
    }
}

SlidePictures.propTypes = {};

export default SlidePictures;