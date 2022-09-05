export default {
  getAnswerValue(state) {
    const value = {
      xValue: state.xValue,
      yValue: state.yValue,
      zValue: state.zValue
    };

    return value;
  },

  getvalueByIndex(state) {
    return state.valueByIndex;
  }
};
