import delay from './delay';

// This file mocks a web API by working with the hard-coded data below.
// It uses setTimeout to simulate the delay of an AJAX call.
// All calls return promises.
const items = [
  {
    id: "1",
    dateStart: "04/03/2017",
    dateEnd: "24/03/2017",
    length: "20",
    description: "test",
    tagId: "cory-house"
  },
  {
    id: "2",
    dateStart: "04/03/2017",
    dateEnd: "24/03/2017",
    length: "20",
    description: "test 2",
    tagId: "scott-allen"
  },
  {
    id: "3",
    dateStart: "04/03/2017",
    dateEnd: "24/03/2017",
    length: "20",
    description: "test 3",
    tagId: "dan-wahlin"
  },
  {
    id: "4",
    dateStart: "04/03/2017",
    dateEnd: "24/03/2017",
    length: "20",
    description: "test 4",
    tagId: "scott-allen"
  }
];

class ItemApi {
  static getAllItems() {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        resolve(Object.assign([], items));
      }, delay);
    });
  }

  static saveItem(item) {
    item = Object.assign({}, item); // to avoid manipulating object passed in.
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        // Simulate server-side validation
        // const minItemTitleLength = 1;
        // if (item.title.length < minItemTitleLength) {
        //   reject(`Title must be at least ${minItemTitleLength} characters.`);
        // }

        if (item.id) {
          const existingItemIndex = items.findIndex(a => a.id == item.id);
          items.splice(existingItemIndex, 1, item);
        } else {
          //Just simulating creation here.
          //Cloning so copy returned is passed by value rather than by reference.
          item.id = items.length + 1;
          items.push(item);
        }

        resolve(item);
      }, delay);
    });
  }

  static deleteItem(itemId) {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        const indexOfItemToDelete = items.findIndex(item => {
          item.id == itemId;
        });
        items.splice(indexOfItemToDelete, 1);
        resolve();
      }, delay);
    });
  }
}

export default ItemApi;