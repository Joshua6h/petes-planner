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
    profile: {},
    profileStatus: ''
  },
  getters: {
    isAuthenticated: state => !!state.token,
    authStatus: state => state.status,
    hasProfile: state => state.profileStatus == 'success',
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
    PROFILE_REQUEST: (state) => {
      state.profileStatus = 'loading'
    },
    PROFILE_SUCCESS: (state, profile) => {
      state.profile = profile
      state.profileStatus = 'success'
    },
    PROFILE_ERROR: (state) => {
      state.profile = {}
      state.profileStatus = 'error'
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
          commit('PROFILE_ERROR', {});
          commit('AUTH_SUCCESS', "");
          resolve(resp);
        })
        .catch(err => {
          commit('AUTH_ERROR');
          reject(err);
        })
      })
    },
    GET_PROFILE: ({commit}) => {
      return new Promise((resolve, reject) => {
        commit('PROFILE_REQUEST');
        axios.defaults.headers.common['Authorization'] = localStorage.getItem('user-token');
        axios({url: 'http://localhost:8080/getprofile', data: {}, method: 'POST'})
        .then(resp => {
          commit('PROFILE_SUCCESS', resp.data.data);
          resolve(resp);
        })
        .catch(err => {
          commit('PROFILE_ERROR');
          reject(err);
        })
      })
    }
  },
  modules: {
  }
})
