import delay from './delay';

// This file mocks a web API by working with the hard-coded data below.
// It uses setTimeout to simulate the delay of an AJAX call.
// All calls return promises.
const users = [
  {
    id: "123456",
    username: "test@test.com",
    password: "123",
    fullname: "test-test",
    token: "123"
  }
];

class AuthApi {

  static onAuthentication(auth) {
    auth = Object.assign({}, auth); // to avoid manipulating object passed in.
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        // Simulate server-side validation
        const minUsernameLength = 3;
        if (auth.username.length < minUsernameLength) {
          reject(`Username must be at least ${minUsernameLength} characters.`);
        }

        const existingUser = users.find(user => user.username === auth.username && user.password === auth.password);
        if(existingUser === null){
          reject(`Username not exist ${auth.username}.`);
        }

        resolve(existingUser);
      }, delay);
    });
  }

  // static deleteItem(itemId) {
  //   return new Promise((resolve, reject) => {
  //     setTimeout(() => {
  //       const indexOfItemToDelete = items.findIndex(item => {
  //         item.id == itemId;
  //       });
  //       items.splice(indexOfItemToDelete, 1);
  //       resolve();
  //     }, delay);
  //   });
  // }
}

export default AuthApi;