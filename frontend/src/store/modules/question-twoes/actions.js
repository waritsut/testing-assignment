import cashierService from '../../../services/cashier-service';
import API from '../../../services/api';

import { notify } from '@kyvg/vue3-notification';

export default {
  async getBalance(context) {
    try {
      const response = await cashierService.getBalance();
      context.commit('setBalance', response.data.data);
    } catch (error) {
      notify({
        type: 'error',
        title: 'Error occurred',
        text: error.response.data.message
      });
    }
  },

  async getTheChange(context, payload) {
    try {
      const response = await cashierService.getTheChange(payload);
      context.commit('setDrawer', response.data.data);
      context.commit('setTheChange', response.data.data);
      notify({
        type: 'info',
        title: 'Success',
        text: 'success'
      });
    } catch (error) {
      notify({
        type: 'error',
        title: 'Error occurred',
        text: error.response.data.message
      });
    }
  },

  async resetCashier(context) {
    const payload = {
      action: 'cashier-reset'
    };
    const response = API.post('/handle', payload)
      .then((response) => {
        if (response.status == 200) {
          const payload = {
            action: 'cashier-get-balance'
          };
          const res = API.post('/handle', payload).then((res) => {
            if (res.status == 200) {
              context.commit('setBalance', res.data.data);
              notify({
                type: 'info',
                title: 'Success',
                text: 'success'
              });
            } else {
              throw new Error(`Cloud not reset the cashier`);
            }
          });
          console.log(res);
        } else {
          throw new Error(`Cloud not reset the cashier`);
        }
      })
      .catch((error) => {
        notify({
          type: 'error',
          title: 'Cloud not reset the cashier',
          text: error.response.data.message
        });
        return;
      });

    console.log(response);
  }
};
