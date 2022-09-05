export default {
  setBalance(state, payload) {
    payload.cash[1000] = payload.cash['oneThousandNote'];
    delete payload.cash['oneThousandNote'];

    payload.cash[500] = payload.cash['fiveHundredNote'];
    delete payload.cash['fiveHundredNote'];

    payload.cash[100] = payload.cash['oneHundredNote'];
    delete payload.cash['oneHundredNote'];

    payload.cash[50] = payload.cash['fiftyNote'];
    delete payload.cash['fiftyNote'];

    payload.cash[20] = payload.cash['twentyNote'];
    delete payload.cash['twentyNote'];

    payload.cash[10] = payload.cash['tenCoin'];
    delete payload.cash['tenCoin'];

    payload.cash[5] = payload.cash['fiveCoin'];
    delete payload.cash['fiveCoin'];

    payload.cash[1] = payload.cash['oneCoin'];
    delete payload.cash['oneCoin'];

    payload.cash[0.25] = payload.cash['twentyFiveSatang'];
    delete payload.cash['twentyFiveSatang'];

    state.balance = payload.balance;
    state.cashInCashier = payload.cash;
  },

  setDrawer(state, payload) {
    payload.availableCash[1000] = payload.availableCash['oneThousandNote'];
    delete payload.availableCash['oneThousandNote'];

    payload.availableCash[500] = payload.availableCash['fiveHundredNote'];
    delete payload.availableCash['fiveHundredNote'];

    payload.availableCash[100] = payload.availableCash['oneHundredNote'];
    delete payload.availableCash['oneHundredNote'];

    payload.availableCash[50] = payload.availableCash['fiftyNote'];
    delete payload.availableCash['fiftyNote'];

    payload.availableCash[20] = payload.availableCash['twentyNote'];
    delete payload.availableCash['twentyNote'];

    payload.availableCash[10] = payload.availableCash['tenCoin'];
    delete payload.availableCash['tenCoin'];

    payload.availableCash[5] = payload.availableCash['fiveCoin'];
    delete payload.availableCash['fiveCoin'];

    payload.availableCash[1] = payload.availableCash['oneCoin'];
    delete payload.availableCash['oneCoin'];

    payload.availableCash[0.25] = payload.availableCash['twentyFiveSatang'];
    delete payload.availableCash['twentyFiveSatang'];

    state.balance = payload.balance;
    state.cashInCashier = payload.availableCash;
  },

  setTheChange(state, payload) {
    state.change = payload.change;
    state.changeCash = payload.changeCash;
  }
};
