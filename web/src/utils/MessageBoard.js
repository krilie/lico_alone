import { notification } from 'antd';

export default function openNotification(text) {
    notification.info({
        message: text,
        // description: text,
        placement: "bottomLeft"
    })
}