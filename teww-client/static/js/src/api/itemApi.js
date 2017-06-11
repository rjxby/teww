import {checkHttpStatus, parseJSON, getToken, propIsExist} from './utils';

class ItemApi {
    static getAllItems() {
        return fetch('/api/items', {
                method: 'get',
                credentials: 'include',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': getToken()
                }
            })
            .then(checkHttpStatus)
            .then(parseJSON);
    }

    static saveItem(item) {
        return fetch('/api/items', {
            method: propIsExist(item.id) ? 'put' : 'post',
            credentials: 'include',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': getToken()
            },
                body: JSON.stringify(item)
            })
            .then(checkHttpStatus)
            .then(parseJSON);
    }

    static removeItem(id) {
        return fetch('/api/items', {
            method: 'delete',
            credentials: 'include',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': getToken()
            },
                body: JSON.stringify({id: id})
            })
            .then(checkHttpStatus);
    }
}

export default ItemApi;
