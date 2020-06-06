import Vue from "vue";
import Vuex from "vuex";
import router from "../router/index";
import jwtDecode from "jwt-decode";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    status: false,
    user: {
      token: null,
      user_name: null,
    },
  },
  getters: {
    getUser(state) {
      return state.user;
    },
    getStatus(state) {
      return state.isAuthorized;
    },
  },
  mutations: {
    postLogin(state, { email, password }) {
      fetch("http://localhost:3000/api/v1/users", {
        method: "POST",
        mode: "cors",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: email,
          password: password,
        }),
      }).then((response) => {
        response
          .json()
          .then((jsonData) => {
            let decoded = jwtDecode(jsonData);
            console.log(decoded);
            state.user.token = jsonData;
            state.user.user_name = decoded.name;
            state.status = true;
            router.push("/");
          })
          .catch((err) => {
            console.log(err);
          });
      });
    },
    logout(state) {
      state.isAuthorized = false;
      state.user.token = null;
      state.user.user_name = null;
    },
    postRegister(state, { username, email, password }) {
      axios
        .post("http://localhost:3000/api/v1/users/register", {
          username,
          email,
          password,
        })
        .then((response) => {
          let data = response.data;
          let decoded = jwtDecode(data.token);
          console.log(decoded);
          state.user.token = data.token;
          state.user.user_name = decoded.name;
          state.status = true;
          router.push("/");
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
  actions: {
    postLogin(context, { email, password }) {
      context.commit("postLogin", { email, password });
    },
    logout(context) {
      context.commit("logout");
    },
    postRegister(context, { username, email, password }) {
      context.commit("postRegister", { username, email, password });
    },
  },
  modules: {},
});
