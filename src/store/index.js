import axios from 'axios'
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    primary: "#FCEDDA",
    accent: "#EE4E34",
    token: localStorage.getItem('user-token') || '',
    status: '',
    profile: {}
  },
  getters: {
    isAuthenticated: state => !!state.token,
    authStatus: state => state.status,
    username: state => state.profile['username'],
    firstname: state => state.profile['firstname'],
    lastname: state => state.profile['lastname'],
    email: state => state.profile['email']
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
    },
    SET_PROFILE: (state, data) => {
      state.profile = data
    }
  },
  actions: {
    AUTH_REQUEST: ({commit}, user) => {
      return new Promise((resolve, reject) => {
        commit('AUTH_REQUEST');
        axios({url: 'http://localhost:8080/login', data: user, method: 'POST'})
        .then(resp => {
          const token = resp.data.userguid;
          localStorage.setItem('user-token', token);
          axios.defaults.headers.common['Authorization'] = token;
          commit('AUTH_SUCCESS', token);
          resolve(resp);
        })
        .catch(err => {
          localStorage.removeItem('user-token');
          commit('AUTH_ERROR');
          reject(err);
        })
      })
    },
    LOGOUT: ({commit, state}) => {
      return new Promise((resolve, reject) => {
        commit('AUTH_REQUEST');
        axios.defaults.headers.common['Authorization'] = state.token;
        axios({url: 'http://localhost:8080/logout', data: {}, method: 'POST'})
        .then(resp => {
          localStorage.removeItem('user-token');
          axios.defaults.headers.common['Authorization'] = "";
          commit('AUTH_SUCCESS', "");
          resolve(resp);
        })
        .catch(err => {
          commit('AUTH_ERROR');
          reject(err);
        })
      })
    },
    GET_PROFILE: ({commit, state}) => {
      return new Promise((resolve, reject) => {
        axios.defaults.headers.common['Authorization'] = state.token;
        axios({url: 'http://localhost:8080/getprofile', data: {}, method: 'POST'})
        .then(resp => {
          commit('SET_PROFILE', resp.data.data);
          resolve(resp);
        })
        .catch(err => {
          commit('SET_PROFILE', {});
          reject(err);
        })
      })
    }
  },
  modules: {
  }
})
