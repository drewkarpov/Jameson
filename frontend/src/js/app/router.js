import Vue from 'vue'
import VueRouter from 'vue-router'
import MainComponent from '../views/Main';
import NotFoundComponent from '../views/NotFound';
import VoidZonesComponent from '../views/VoidZones';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    redirect: '/test/'
  },
  {
    path: '/test',
    component: MainComponent
  },
  {
    path: '/test/:testIdParam',
    name: 'test',
    component: MainComponent,
    props: true
  },
  {
    path: '/voidzones/',
    component: VoidZonesComponent
  },
  {
    path: '/voidzones/:containerIdParam',
    name: 'container',
    component: VoidZonesComponent,
    props: true
  },
  {
    path: '*',
    component: NotFoundComponent
  }
];

const router = new VueRouter({
  scrollBehavior() {
    return {x: 0, y: 0}
  },
  routes
});

export default router;
