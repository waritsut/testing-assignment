import seqNumService from '../../../services/seq-number-service.js';
import { notify } from '@kyvg/vue3-notification';

export default {
  async getAns(context) {
    try {
      const response = await seqNumService.getAns();
      const value = {
        x: response.data.data.xValue,
        y: response.data.data.yValue,
        z: response.data.data.zValue
      };
      context.commit('setAns', value);
    } catch (error) {
      notify({
        type: 'error',
        title: 'Error occurred',
        text: error.response.data.message
      });
    }
  },

  async getValueByIndex(context, index) {
    try {
      const response = await seqNumService.getValueByIndex(index);
      context.commit('setValueByIndex', response.data.data.value);
    } catch (error) {
      notify({
        type: 'error',
        title: 'Error occurred',
        text: error.response.data.message
      });
    }
  }
};
