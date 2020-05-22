export const GET_SETTING = 'GET_SETTING';

let actions = {
    getSettings: function (settings) {
        return {type: GET_SETTING, payload: settings}
    },
};

export default actions;