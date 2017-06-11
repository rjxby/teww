import * as types from '../actions/actionTypes';
import initialState from './initialState';
import {propIsExist} from '../api/utils';
//todo
export default function authReducer(state = initialState.user, action) {
    switch (action.type) {
        case types.CHECK_AUTHENTICATION:
            return {
                isAuthenticated: propIsExist(action.user),
                user: action.user
            };
        case types.ON_AUTHENTICATION_SUCCESS:
            return {
                isAuthenticated: propIsExist(action.user),
                user: action.user
            };
        case types.LOG_OUT_SUCCESS:
            return {
                isAuthenticated: propIsExist(action.user),
                user: action.user
            };
        default:
            return state;
    }
}