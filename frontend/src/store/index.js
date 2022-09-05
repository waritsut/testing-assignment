import { createStore } from 'vuex';

import questionOneModule from './modules/question-ones/index.js';
import questionTwoModule from './modules/question-twoes/index.js';

export default createStore({
  modules: {
    questionOnes: questionOneModule,
    questionTwoes: questionTwoModule
  }
});
