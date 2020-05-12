import {getQuery} from "./api";
import actions from "../redux/actions/ActionCreator";

export const getSettingListAllRedux = () => {
    return (dispatch) => {
        getQuery("/manage/setting/get_setting_all")
            .then((res) => {
                const data = res.data; // data 就是body
                const action = actions.getSettings(data);
                dispatch(action);
            })
    }
};