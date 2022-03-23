import Vue from 'vue';
import VueKonva from 'vue-konva';
import {
  BootstrapVue,
  BIcon,
  BIconBoundingBoxCircles,
  BIconTrash,
  BIconCheck2Circle
} from 'bootstrap-vue'

import Percentage from '../components/Percentage';

/** Register Global Components */
export default {
  register() {
    Vue.use(BootstrapVue);
    Vue.use(VueKonva);

    Vue.component('BIcon', BIcon);
    Vue.component('BIconBoundingBoxCircles', BIconBoundingBoxCircles);
    Vue.component('BIconTrash', BIconTrash);
    Vue.component('BIconCheck2Circle', BIconCheck2Circle);

    Vue.component('percentage', Percentage);
  }
}
