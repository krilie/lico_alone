import {createStore} from 'redux';
import reducer from './reducers/Reducers';

let store = createStore(reducer);//传入reducer
export default store;