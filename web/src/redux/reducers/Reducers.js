import initState from "./States";
import {ADD_TODO, TOGGLE_COMPLETE} from "../actions/ActionItems";

function reducer(state = initState, action) {
    let newState;
    switch (action.type) {
        case ADD_TODO:
            newState = {
                todos: [
                    ...state.todos,
                    action.payload
                ]
            };
            break;
        case TOGGLE_COMPLETE:
            newState = {
                todos: state.todos.map(item => {
                    if (item.id === action.payload) {
                        item.isComplete = false;
                    }
                    return item;
                })
            };
            break;
        default:
            newState = state;
            break;
    }
    return newState;
}

export default reducer;