export const GET_DIR_FILE_ITEMS = 'GET_DIR_FILE_ITEMS';

let actions = {
    getDirFileItems: function (DirData) {
        return {type: GET_DIR_FILE_ITEMS, payload: DirData}
    },
};

export default actions;