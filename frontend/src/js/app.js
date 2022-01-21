import Vue from 'vue';
import globalComponents from './app/ui';
import router from './app/router';

import App from './views/App';

/** Global Components */
globalComponents.register();

/** Application */
new Vue({
  router,
  render: h => h(App)
}).$mount('#jameson-app');