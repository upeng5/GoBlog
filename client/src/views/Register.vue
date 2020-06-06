<template>
  <div class="register center">
    <vs-row>
      <vs-col offset="5" w="6">
        <vs-input warn v-model="username" placeholder="Username" class="username-field">
          <template #icon>
            <i class="material-icons">person</i>
          </template>
        </vs-input>

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

        <vs-button warn class="submit-button" @click="postRegister">
          Submit
          <template #animate>
            <i class="material-icons">login</i>
          </template>
        </vs-button>
        <!-- <Spinner size="medium" /> -->
      </vs-col>
    </vs-row>
  </div>
</template>

<script>
import Spinner from "vue-simple-spinner";

export default {
  name: "Register",
  components: {
    Spinner
  },
  data() {
    return {
      username: "",
      email: "",
      password: ""
    };
  },
  methods: {
    postRegister() {
      if (this.username.length === 0) {
        this.errorMessage = "Please enter a valid username";
        this.openNotification();
        this.loading = false;
        return;
      }
      if (this.email.length === 0) {
        this.errorMessage = "Please enter a valid email";
        this.openNotification();
        this.loading = false;
        return;
      }
      if (this.password.length === 0) {
        this.errorMessage = "Please enter a valid password";
        this.openNotification();
        this.loading = false;
        return;
      }

      this.$store.dispatch("postRegister", {
        username: this.username,
        email: this.email,
        password: this.password
      });
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
.register {
  margin-top: 40px;
}

.register .email-field {
  margin-bottom: 30px;
}

.register .username-field {
  margin-bottom: 30px;
}

.register .password-field {
}

.register .submit-button {
  margin-top: 20px;
  width: 30%;
}
</style>
