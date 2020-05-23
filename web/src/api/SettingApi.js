import {getQuery, postForm} from "./api";
import actions from "../redux/actions/ActionCreator";
// 获取所有配置列表=>redux
export const getSettingListAllRedux = () => {
    return (dispatch) => {
        getQuery("/manage/setting/get_setting_all")
            .then((res) => {
                const data = res.data.data; // data 就是body
                const action = actions.getSettings(data);
                dispatch(action);
            })
    }
};

// 更新配置 已经检查2000成功
export const updateSettingItem = (name, value) => {
    return postForm("/manage/setting/update_config", {name, value});
}

