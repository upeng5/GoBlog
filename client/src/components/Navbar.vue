<template>
  <div class="navbar">
    <div class="container">
      <router-link to="/" class="navbar-brand">
        <i class="material-icons book">book</i> GoBlog
      </router-link>

      <div class="nav-menu">
        <ul v-if="status">
          <li>
            <router-link to="/">
              <vs-button shadow>
                Home
                <template #animate>
                  <i class="material-icons">home</i>
                </template>
              </vs-button>
            </router-link>
          </li>
          <li>
            <vs-button danger shadow @click="logout">Logout</vs-button>
          </li>
        </ul>
        <!-- isAuth -->
        <ul v-else>
          <li>
            <router-link to="/login">
              <vs-button shadow>Login</vs-button>
            </router-link>
          </li>
          <li>
            <router-link to="/register">
              <vs-button shadow>Register</vs-button>
            </router-link>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Navbar",
  computed: {
    status() {
      return this.$store.state.status;
    }
  },
  methods: {
    logout() {
      this.$store.dispatch("logout");
      this.$store.state.status = false;
      this.$router.push("/login");
    }
  }
};
</script>

<style>
.navbar {
  background: #ffba00;
  padding: 30px;
  margin: 0;
}

.navbar .container {
  width: 80%;
}

.navbar .navbar-brand {
  font-size: 20px;
  color: #000;
  text-decoration: none;
  float: left;
  margin-top: -13px;
  font-weight: bold;
  letter-spacing: 1px;
}

.navbar .book {
  margin-right: -5px;
  font-size: 20px;
}

.navbar .nav-menu {
  float: right;
  margin-top: -37px;
}

.navbar .nav-menu ul {
  margin-right: -250px;
  list-style: none;
}

.navbar .nav-menu li {
  float: left;
  margin-left: 6px;
}
.navbar .nav-menu li a {
  text-decoration: none;

  color: #ffba00 !important;
}
</style>
