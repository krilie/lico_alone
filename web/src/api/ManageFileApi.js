import {getQuery} from "./ApiManage";
import actions from "../redux/actions/ActionCreator";

export const getFilePage = (path) => {
    return (dispatch) => {
        getQuery("/file/get_file_list", {path: path})
            .then((res) => {
                const data = res.data; // data 就是body
                const action = actions.getDirFileItems({dirPath: path, items: data});
                dispatch(action);
            })
    }
};

