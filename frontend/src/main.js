import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import BaseCard from './components/ui/BaseCard.vue';
import BaseBadge from './components/ui/BaseBadge.vue';
import Notifications from '@kyvg/vue3-notification';

const app = createApp(App);

app.use(router);
app.use(store);
app.use(Notifications);

app.component('base-card', BaseCard);
app.component('base-badge', BaseBadge);

app.mount('#app');
