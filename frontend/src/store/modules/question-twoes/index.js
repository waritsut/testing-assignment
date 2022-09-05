import mutations from './mutations.js';
import actions from './actions';
import getters from './getters.js';

export default {
  namespaced: true,
  state() {
    return {
      isLoading: false,
      balance: 0,
      cashInCashier: {
        oneThousandNote: 0,
        fiveHundredNote: 0,
        oneHundredNote: 0,
        fiftyNote: 0,
        twentyNote: 0,
        tenCoin: 0,
        fiveCoin: 0,
        oneCoin: 0,
        twentyFiveSatang: 0
      },
      change: 0,
      changeCash: {
        oneThousandNote: 0,
        fiveHundredNote: 0,
        oneHundredNote: 0,
        fiftyNote: 0,
        twentyNote: 0,
        tenCoin: 0,
        fiveCoin: 0,
        oneCoin: 0,
        twentyFiveSatang: 0
      }
    };
  },
  mutations,
  actions,
  getters
};
