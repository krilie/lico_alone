import initState from "./States";
import {GET_SETTING} from "../actions/ActionCreator";

function reducer(state = initState, action) {
    let newState;
    switch (action.type) {
        case GET_SETTING:
            newState = {...state, settings: action.payload};
            break;
        default:
            newState = state;
            break;
    }
    return newState;
}

export default reducer;