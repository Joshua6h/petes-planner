import axios from 'axios'
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    primary: "#FCEDDA",
    accent: "#EE4E34",
    token: localStorage.getItem('user-token') || '',
    status: ''
  },
  getters: {
    isAuthenticated: state => !!state.token,
    authStatus: state => state.status
  },
  mutations: {
    AUTH_REQUEST: (state) => {
      state.status = 'authenticating'
    },
    AUTH_SUCCESS: (state, token) => {
      state.status = 'success'
      state.token = token
    },
    AUTH_ERROR: (state) => {
      state.status = 'error'
    }
  },
  actions: {
    AUTH_REQUEST: ({commit}, user) => {
      return new Promise((resolve, reject) => {
        commit('AUTH_REQUEST');
        axios({})
      })
    }
  },
  modules: {
  }
})
