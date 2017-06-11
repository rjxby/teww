import {combineReducers} from 'redux';
import items from './itemReducer';
import tags from './tagReducer';
import ajaxCallsInProgress from './ajaxStatusReducer';
import auth from './authReducer';
import * as types from '../actions/actionTypes';
import initialState from './initialState';

const rootReducer = combineReducers({items, tags, ajaxCallsInProgress, auth});

// const rootReducer = (state, action) => {
//     if (action.type === types.LOG_OUT_SUCCESS) {
//         state.tags = initialState.tags;
//         state.items = initialState.items; //todo
//     }
//     return appReducer(state, action);
// };

export default rootReducer;