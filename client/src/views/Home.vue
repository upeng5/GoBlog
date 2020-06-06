<template>
  <div class="home">
    <FloatingAdd />
    <div v-for="(post, index) in posts" :key="index">
      <Post :post="post" />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import store from "../store/index";
import FloatingAdd from "../components/FloatingAdd";
import Post from "../components/Post";
import axios from "axios";

export default {
  name: "Home",
  components: {
    FloatingAdd,
    Post
  },
  data() {
    return {
      posts: [],
      loading: true
    };
  },
  store: store,
  methods: {
    getPosts() {
      axios
        .get("http://localhost:3000/api/v1/posts", {
          headers: {
            "x-auth-token": this.userToken
          }
        })
        .then(res => {
          if (res.status === 200) {
            console.log(res.data);
            this.posts = res.data;
          }
          console.log(res);
        })
        .catch(err => {
          console.log(err);
        });
    }
  },
  computed: {
    userToken() {
      return this.$store.state.user.token;
    }
  },
  created() {},
  mounted() {
    this.getPosts();
    this.loading = false;
  }
};
</script>

<style>
.home {
  text-align: center;
}

.home .spinner {
  position: absolute;
  top: -50%;
  left: -50%;
}
</style>
