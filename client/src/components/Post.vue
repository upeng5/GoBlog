<template>
  <div class="post-component">
    <div class="post">
      <vs-row>
        <vs-col w="2" class="avatar-col">
          <vs-avatar warn circle size="60">
            <template #text>{{ post.username }}</template>
          </vs-avatar>
        </vs-col>
        <vs-col w="10">
          <h3 class="post-title">{{ post.title }}</h3>
          <p>{{ post.content }}</p>
        </vs-col>

        <vs-button
          v-if="user.user_name === post.username"
          circle
          icon
          floating
          color="#ffba00"
          class="f-button"
          @click="deletePost"
        >
          <i class="material-icons">delete</i>
        </vs-button>
      </vs-row>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Post",
  props: ["post"],
  computed: {
    user() {
      return this.$store.state.user;
    }
  },
  methods: {
    deletePost() {
      const postId = this.post._id;
      console.log(postId);
      axios
        .delete(`http://localhost:3000/api/v1/posts/${postId}`, {
          headers: { "x-auth-token": this.user.token }
        })
        .then(response => {
          console.log("Delete Request:", response);
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>

<style>
.post-component .post {
  max-width: 500px;
  padding: 10px;
  background: #fff;
  margin: 40px auto;
  border-radius: 25px;
  -webkit-box-shadow: 0px 0px 28px -14px rgba(0, 0, 0, 0.75);
  -moz-box-shadow: 0px 0px 28px -14px rgba(0, 0, 0, 0.75);
  box-shadow: 0px 0px 28px -14px rgba(0, 0, 0, 0.75);
  border: solid 2px #ffba00;
  transition: all 450ms;
}

.post-component .post:hover {
  transform: scale(1.1);
}

.post-component .avatar-col {
  border-right: solid 1px #eee;
}

.post .red {
  background: red;
}
</style>