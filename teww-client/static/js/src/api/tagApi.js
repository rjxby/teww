import {checkHttpStatus, parseJSON, getToken} from './utils';

class TagApi {
    static getAllTags() {
        return fetch('/api/tags', {
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
}

export default TagApi;

// export function loginUser(email, password, redirect="/") {
//     return function(dispatch) {
//         dispatch(loginUserRequest());
//         return fetch('http://localhost:3000/auth/getToken/', {
//             method: 'post',
//             credentials: 'include',
//             headers: {
//                 'Accept': 'application/json',
//                 'Content-Type': 'application/json'
//             },
//                 body: JSON.stringify({email: email, password: password})
//             })
//             .then(checkHttpStatus)
//             .then(parseJSON)
//             .then(response => {
//                 try {
//                     let decoded = jwtDecode(response.token);
//                     dispatch(loginUserSuccess(response.token));
//                     dispatch(pushState(null, redirect));
//                 } catch (e) {
//                     dispatch(loginUserFailure({
//                         response: {
//                             status: 403,
//                             statusText: 'Invalid token'
//                         }
//                     }));
//                 }
//             })
//             .catch(error => {
//                 dispatch(loginUserFailure(error));
//             })
//     }
// }