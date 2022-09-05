import mutations from './mutations.js';
import actions from './actions';
import getters from './getters.js';

export default {
  namespaced: true,
  state() {
    return {
      isLoading: false,
      xValue: 0,
      yValue: 0,
      zValue: 0,
      valueByIndex: 0
    };
  },
  mutations,
  actions,
  getters
};
