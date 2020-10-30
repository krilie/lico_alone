import React, {Component} from 'react';

export default class NotFoundBox extends Component {
    render() {
        const {msg} = this.props
        return (
            <div>
                {msg}未寻得此内容...
            </div>
        );
    }
}