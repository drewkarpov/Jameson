import Vue from 'vue'
import VueRouter from 'vue-router'
import MainComponent from '../views/Main';
import NotFoundComponent from '../views/NotFound';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'default',
    component: MainComponent
  },
  {
    path: '/test/:testIdParam', // http://my-service.com/test/test-id
    name: 'project',
    component: MainComponent,
    props: true
  },
  {
    path: '*',
    component: NotFoundComponent
  }
];

const router = new VueRouter({
  // mode: 'history',
  scrollBehavior() {
    return {x: 0, y: 0}
  },
  routes
});

export default router;
