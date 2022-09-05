export default {
  getBalance(state) {
    const value = {
      balance: state.balance,
      cashInCashier: state.cashInCashier
    };
    return value;
  },

  getTheChange(state) {
    const value = {
      change: state.change,
      changeCash: state.changeCash
    };
    return value;
  }
};
