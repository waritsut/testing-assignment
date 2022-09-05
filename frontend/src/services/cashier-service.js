import API from './api';

export default {
  async getBalance() {
    const payload = {
      action: 'cashier-get-balance'
    };
    return API.post('/handle', payload);
  },

  async resetCashier() {
    const payload = {
      action: 'cashier-reset'
    };
    return API.post('/handle', payload);
  },

  async getTheChange(cash) {
    const payload = {
      action: 'cashier-cal-the-change',
      cashier: cash
    };
    return API.post('/handle', payload);
  }
};
