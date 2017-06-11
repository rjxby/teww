export function checkHttpStatus(response) {
    if (response.status >= 200 && response.status < 300) {
        return response;
    } else {
        let error = new Error(response.statusText);
        error.response = response;
        throw error;
    }
}

export function parseJSON(response) {
    return response.json();
}

export function clearToken() {
    sessionStorage.removeItem('token');
    sessionStorage.removeItem('tokenType');
    sessionStorage.removeItem('expiresIn');
}

export function getToken() {
    let token = sessionStorage.getItem('token');
    let tokenType = sessionStorage.getItem('tokenType');
    if (token === undefined || tokenType === undefined) {
        throw "Token not set";
    }
    return tokenType + " " + token;
}

export function propIsExist(prop) {
    return Object
        .keys(prop)
        .length !== 0;
}