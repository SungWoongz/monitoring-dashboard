import {createApp} from 'vue';
import App from './App.vue';
import router from './router';
import {setupStore} from './store'; // setupStore 가져오기
import './style.css';

const app = createApp(App);

// Pinia 설정
setupStore(app);

app.use(router);

app.mount('#app');