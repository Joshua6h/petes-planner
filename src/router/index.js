import Vue from 'vue'
import VueRouter from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import AddEvents from '../views/AddEvents.vue'
import Calendar from '../views/Calendar.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/add-events',
    name: 'Add Events',
    component: AddEvents
  },
  {
    path: '/calendar',
    name: 'Calendar',
    component: Calendar
  }
]

const router = new VueRouter({
  routes
})

export default router
