export default {
  setAns(state, payload) {
    state.xValue = payload.x;
    state.yValue = payload.y;
    state.zValue = payload.z;
  },

  setValueByIndex(state, payload) {
    state.valueByIndex = payload;
  }
};
