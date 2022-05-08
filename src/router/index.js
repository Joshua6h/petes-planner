import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '@/store'

Vue.use(VueRouter)

const ifAuthenticated = (to, from, next) => {
  if(store.getters.isAuthenticated) {
    next();
    return;
  }
  next('/login');
}

const ifNotAuthenticated = (to, from, next) => {
  if(!store.getters.isAuthenticated) {
    next();
    return;
  }
  next('/');
}

const routes = [
  {
    path: '/',
    name: 'About',
    component: () => import('@/views/About.vue')
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Dashboard.vue'),
    beforeEnter: ifAuthenticated
  },
  {
    path: '/add-events',
    name: 'Add Events',
    component: () => import('@/views/AddEvents.vue'),
    beforeEnter: ifAuthenticated
  },
  {
    path: '/calendar',
    name: 'Calendar',
    component: () => import('@/views/Calendar.vue'),
    beforeEnter: ifAuthenticated
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    beforeEnter: ifNotAuthenticated
  },
  {
    path: '/sign-up',
    name: 'SignUp',
    component: () => import('@/views/SignUp.vue'),
    beforeEnter: ifNotAuthenticated
  },
  {
    path: '/reset-password',
    name: 'ResetPassword',
    component: () => import('@/views/ResetPassword.vue'),
    beforeEnter: ifNotAuthenticated
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/Profile.vue'),
    beforeEnter: ifAuthenticated
  }
]

const router = new VueRouter({
  routes
})

export default router
