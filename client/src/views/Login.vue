<template>
  <div class="login">
    <vs-row>
      <vs-col offset="5" w="6">
        <vs-input warn v-model="email" placeholder="Email" class="email-field">
          <template #icon>
            <i class="material-icons">email</i>
          </template>
        </vs-input>

        <vs-input
          warn
          class="password-field"
          type="password"
          v-model="password"
          placeholder="Password"
        >
          <template #icon>
            <i class="material-icons">lock</i>
          </template>
        </vs-input>

        <vs-button warn class="submit-button" @click="loginAction">
          Submit
          <template #animate>
            <i class="material-icons">login</i>
          </template>
        </vs-button>
      </vs-col>
    </vs-row>
    <!-- <Spinner v-if="loading" class="spinner" line-fg-color="#ffba00" size="medium" /> -->
  </div>
</template>

<script>
import store from "../store/index";
import Spinner from "vue-simple-spinner";

export default {
  name: "Login",
  store: store,
  components: {
    Spinner
  },
  data() {
    return {
      email: "",
      password: "",
      errorMessage: "",
      loading: false
    };
  },
  computed: {
    isAuth() {
      return this.$store.state.status;
    }
  },
  methods: {
    loginAction() {
      this.loading = true;
      if (this.email.length === 0 || this.password.length === 0) {
        this.errorMessage = "Please enter valid credentials";
        this.openNotification();
        this.loading = false;
        return;
      }
      let data = {
        email: this.email,
        password: this.password
      };
      this.$store.dispatch("postLogin", data);

      this.loading = false;
    },
    openNotification() {
      this.$vs.notification({
        progress: "auto",
        border: "warm",
        position: "top-left",
        title: "Error",
        text: this.errorMessage
      });
    }
  }
};
</script>

<style>
.login {
  margin-top: 40px;
}

.login .email-field {
  margin-bottom: 30px;
}

.login .submit-button {
  margin-top: 20px;
  width: 30%;
}

.login .mb {
  margin-bottom: 40px;
}

.login .spinner {
  margin-top: 30px;
  text-align: center;
  color: #ffba00;
}
</style>
