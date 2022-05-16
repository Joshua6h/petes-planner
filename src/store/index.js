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
    profileStatus: '',
    events: [],
    eventsStatus: '',
    addEventsStatus: '',
    friends: [],
    getFriendsStatus: ''
  },
  getters: {
    isAuthenticated: state => !!state.token,
    authStatus: state => state.status,
    hasProfile: state => state.profileStatus == 'success',
    username: state => state.profile['username'],
    firstname: state => state.profile['firstname'],
    lastname: state => state.profile['lastname'],
    email: state => state.profile['email'],
    events: state => state.events,
    friends: state => state.friends
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
    },
    REQUEST: (state) => {
      state.status = 'requesting'
    },
    SUCCESS: (state) => {
      state.status = 'success'
    },
    ERROR: (state) => {
      state.status = 'error'
    },
    GET_EVENTS_REQUEST: (state) => {
      state.status = 'loading'
    },
    GET_EVENTS_SUCCESS: (state, events) => {
      state.eventsStatus = 'success'
      state.events = events
    },
    GET_EVENTS_ERROR: (state) => {
      state.status = 'error'
      state.profile = {}
    },
    ADD_EVENTS_REQUEST: (state) => {
      state.addEventsStatus = 'loading'
    },
    ADD_EVENTS_SUCCESS: (state) => {
      state.addEventsStatus = 'success'
    },
    ADD_EVENTS_ERROR: (state) => {
      state.addEventsStatus = 'error'
    },
    GET_FRIENDS_REQUEST: (state) => {
      state.getFriendsStatus = 'loading'
    },
    GET_FRIENDS_SUCCESS: (state, friends) => {
      state.getFriendsStatus = 'success'
      state.friends = friends
    },
    GET_FRIENDS_FAILURE: (state) => {
      state.getFriendsStatus = 'success'
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
    },
    CREATE_USER: ({commit}, userInfo) => {
      return new Promise((resolve, reject) => {
        commit('REQUEST');
        axios.defaults.headers.common['Authorization'] = localStorage.getItem('user-token');
        axios({url: 'http://localhost:8080/adduser', data: userInfo, method: 'POST'})
        .then(resp => {
          commit('SUCCESS');
          resolve(resp);
        })
        .catch(err => {
          commit('ERROR');
          reject(err);
        })
      })
    },
    SEND_EMAIL: ({commit}, email) => {
      return new Promise((resolve, reject) => {
        commit('REQUEST');
        axios.defaults.headers.common['Authorization'] = localStorage.getItem('user-token');
        axios({url: 'http://localhost:8080/sendemail', data: email, method: 'POST'})
        .then(resp => {
          commit('SUCCESS');
          resolve(resp);
        })
        .catch(err => {
          commit('ERROR');
          reject(err);
        })
      })
    },
    GET_EVENTS: ({commit}) => {
      return new Promise((resolve, reject) => {
        commit('GET_EVENTS_REQUEST');
        axios.defaults.headers.common['Authorization'] = localStorage.getItem('user-token');
        axios({url: 'http://localhost:8080/getevents', data: {}, method: 'POST'})
        .then(resp => {
          commit('GET_EVENTS_SUCCESS', resp.data);
          resolve(resp);
        })
        .catch(err => {
          commit('GET_EVENTS_ERROR');
          reject(err);
        })
      })
    },
    ADD_EVENT: ({commit}, event) => {
      return new Promise((resolve, reject) => {
        commit('ADD_EVENTS_REQUEST');
        console.log(event)
        axios.defaults.headers.common['Authorization'] = localStorage.getItem('user-token');
        console.log(localStorage.getItem('user-token'))
        axios({url: 'http://localhost:8080/addevent', data: event, method: 'POST'})
        .then(resp => {
          commit('ADD_EVENTS_SUCCESS');
          resolve(resp);
          this.$store.dispatch('GET_EVENTS');
        })
        .catch(err => {
          commit('ADD_EVENTS_ERROR');
          reject(err);
        })
      })
    },
    GET_FRIENDS: ({commit}) => {
      return new Promise((resolve, reject) => {
        commit('GET_FRIENDS_REQUEST');
        axios.defaults.headers.common['Authorization'] = localStorage.getItem('user-token');
        axios({url: 'http://localhost:8080/getfriends', data: {}, method: 'POST'})
        .then(resp => {
          commit('GET_FRIENDS_SUCCESS', resp.data);
          resolve(resp);
        })
        .catch(err => {
          commit('GET_FRIENDS_ERROR');
          reject(err);
        })
      })
    },
  },
  
  modules: {
  }
})
