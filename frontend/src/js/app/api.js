import axios from 'axios';

// const URL = 'http://127.0.0.1:3333/api/v1';
// const URL = './test_data';
const URL = 'https://testing.rezero.pro/api/v1';

export default {
  getTest: function (testId) {
    return axios.get(URL + '/test/' + testId);
  },
  getContainer: function (containerId) {
    return axios.get(URL + '/container/' + containerId);
  },
  setContainerVoidZones: function (containerId, voidZones) {
    return axios.post(URL + '/container/' + containerId + '/add/voidzone', voidZones);
  },
  approveContainer: function (containerId) {
    return axios.patch(URL + '/container/' + containerId + '/approve');
  },
  getImageURL: function (imageId) {
    return URL + '/image/' + imageId;
  }
};