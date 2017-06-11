import { hashHistory } from 'react-router';
import * as types from './actionTypes';
import {logoutSuccess} from './authActions';
import tagApi from '../api/tagApi';
import {clearToken} from '../api/utils';
import {beginAjaxCall, ajaxCallError} from './ajaxStatusActions';

export function loadTagsSuccess(tags) {
    return {type: types.LOAD_TAGS_SUCCESS, tags};
}

export function loadTags() {
    return function (dispatch, getState) {
        dispatch(beginAjaxCall());
        return tagApi
            .getAllTags()
            .then(tags => {
                dispatch(loadTagsSuccess(tags));
            })
            .catch(error => {
                if (error.response.status === 401) {
                    clearToken();
                    let user = {};
                    dispatch(logoutSuccess(user));
                    hashHistory.push('/');
                } else {
                    dispatch(ajaxCallError(error));
                    throw error;
                }
            });
    };
}
