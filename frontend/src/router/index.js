import { createRouter, createWebHistory } from 'vue-router';

import QuestionOne from '../pages/Questions/QuestionOne.vue';
import QuestionTwo from '../pages/Questions/QuestionTwo.vue';

import NotFound from '../pages/NotFound.vue';

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes: [
    { path: '/', redirect: '/question1' },
    { path: '/question1', component: QuestionOne },
    { path: '/question2', component: QuestionTwo },
    { path: '/:notFound(.*)', component: NotFound }
  ]
});

export default router;
