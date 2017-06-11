import * as types from './actionTypes';
import authApi from '../api/authApi';
import {clearToken} from '../api/utils';
import {beginAjaxCall, ajaxCallError} from './ajaxStatusActions';

export function checkAuthentication(user) {
    return {type: types.CHECK_AUTHENTICATION, user};
}

export function onAuthenticationSuccess(user) {
    return {type: types.ON_AUTHENTICATION_SUCCESS, user};
}

export function logoutSuccess(user) {
    return {type: types.LOG_OUT_SUCCESS, user};
}

export function checkAuth() {
    return function (dispatch) {
        let user = {};
        if (sessionStorage.getItem('token')) {
            user.token = sessionStorage.getItem('token');
        }
        dispatch(checkAuthentication(user));
    };
}

export function onAuthentication(auth) {
    return function (dispatch, getState) {
        dispatch(beginAjaxCall());
        return authApi
            .onAuthentication(auth)
            .then(authenticatedUser => {
                if(Object.keys(authenticatedUser).length === 0){
                    throw "Authorization fails";
                }
                sessionStorage.setItem('token', authenticatedUser.token);
                sessionStorage.setItem('tokenType', authenticatedUser.tokenType);
                sessionStorage.setItem('expiresIn', authenticatedUser.expiresIn);
                dispatch(onAuthenticationSuccess(authenticatedUser));
            })
            .catch(error => {
                dispatch(ajaxCallError(error));
                throw error;
            });
    };
}

export function logOut() {
    return function (dispatch, getState) {
        dispatch(beginAjaxCall());
        return authApi
            .logOut()
            .then(logoutUser => {
                clearToken();
                let user = {};
                dispatch(logoutSuccess(user));
            })
            .catch(error => {
                dispatch(ajaxCallError(error));
                throw error;
            });
    };
}
