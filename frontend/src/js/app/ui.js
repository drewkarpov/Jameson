import Vue from 'vue';
import {BootstrapVue} from 'bootstrap-vue'

import Percentage from '../components/Percentage';

/** Register Global Components */
export default {
  register() {
    Vue.use(BootstrapVue);

    Vue.component('percentage', Percentage);
  }
}
