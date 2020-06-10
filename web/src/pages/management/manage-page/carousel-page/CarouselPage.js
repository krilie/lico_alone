import React, {Component} from 'react';
import {manageGetCarouselList} from "../../../../api/ManageCarouselApi";
import {message} from "antd";

class CarouselPage extends Component {

    constructor(props) {
        super(props);
        this.state = {
            carousels: []
        }
    }

    componentDidMount() {
        this.loadCarouselData()
    }

    loadCarouselData = () => {
        manageGetCarouselList().then(res => {
            this.setState({
                carousels: res.data.data
            })
        }).catch(err => {
            message.error(err)
        }).finally(() => {

        })
    }

    render() {
        const {carousels} = this.state
        return (
            <div>
                {carousels.map(val => <div>val.id</div>)}
            </div>
        );
    }
}


export default CarouselPage;