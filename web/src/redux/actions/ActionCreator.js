import {ADD_TODO, TOGGLE_COMPLETE} from "./ActionItems";

let actions = {
    addTodo: function (payload) {
        return {type: ADD_TODO, payload}
    },
    toggleComplete: function (payload) {
        return {type: TOGGLE_COMPLETE, payload}
    }
};
export default actions;