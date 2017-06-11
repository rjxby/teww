import * as types from './actionTypes';
import itemApi from '../api/itemApi'; 
import {beginAjaxCall, ajaxCallError} from './ajaxStatusActions';

export function loadItemsSuccess(items){
    return {type: types.LOAD_ITEMS_SUCCESS, items};
}

export function createItemSuccess(item){
    return {type: types.CREATE_ITEM_SUCCESS, item};
}

export function updateItemSuccess(item){
    return {type: types.UPDATE_ITEM_SUCCESS, item};
}

export function removedItemSuccess(id){
    return {type: types.REMOVE_ITEM_SUCCESS, id};
}

export function loadItems(){
    return function(dispatch){
        dispatch(beginAjaxCall());
        return itemApi.getAllItems().then(items => {
            dispatch(loadItemsSuccess(items));
        }).catch(error => {
            dispatch(ajaxCallError(error));
            throw error;
        });
    };
}

export function saveItem(item){
    return function(dispatch, getState){
        dispatch(beginAjaxCall());
        return itemApi.saveItem(item).then(savedItem => {
            item.id ? dispatch(updateItemSuccess(savedItem)) :
                dispatch(createItemSuccess(savedItem));
        }).catch(error => {
            dispatch(ajaxCallError(error));
            throw error;
        });
    };
}

export function removeItem(id){
    return function(dispatch, getState){
        dispatch(beginAjaxCall());
        return itemApi.removeItem(id).then(removedItem => 
            dispatch(removedItemSuccess(id))).catch(error => {
                dispatch(ajaxCallError(error));
                throw error;
        });
    };
}