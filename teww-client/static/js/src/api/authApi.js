import {checkHttpStatus, parseJSON, getToken} from './utils';

class AuthApi {
  static onAuthentication(auth) {
    return fetch('/auth/login', {
      method: 'post',
      credentials: 'include',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
        body: JSON.stringify({username: auth.username, password: auth.password})
      })
      .then(checkHttpStatus)
      .then(parseJSON);
  }

  static logOut(auth) {
    return fetch('/auth/logout', {
      method: 'post',
      credentials: 'include',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': getToken()
      }
    }).then(checkHttpStatus);
  }
}

export default AuthApi;