import initState from "./States";
import {GET_DIR_FILE_ITEMS} from "../actions/ActionCreator";

function reducer(state = initState, action) {
    let newState;
    switch (action.type) {
        case GET_DIR_FILE_ITEMS:
            newState = {...state, fileItems: action.payload};
            break;
        default:
            newState = state;
            break;
    }
    return newState;
}

export default reducer;