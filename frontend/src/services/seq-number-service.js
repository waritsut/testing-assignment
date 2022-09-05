import API from './api';

export default {
  async getAns() {
    const payload = {
      action: 'seq-num'
    };
    return API.post('/handle', payload);
  },

  async getValueByIndex(index) {
    const payload = {
      action: 'seq-num-by-index',
      seqNum: {
        index: index
      }
    };
    return API.post('/handle', payload);
  }
};
