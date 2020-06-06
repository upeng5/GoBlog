<template>
  <div class="add-post">
    <h1>Add Post</h1>
    <vs-row>
      <vs-col class="add-form">
        <div>
          <vs-input warn border v-model="title" placeholder="Title" class="title-field">
            <template #icon>
              <i class="material-icons">title</i>
            </template>
          </vs-input>

          <textarea class="content-field" cols="50" rows="10" v-model="content"></textarea>

          <vs-button warn class="submit-button" @click="createPost">
            Submit
            <template #animate>
              <i class="material-icons">add</i>
            </template>
          </vs-button>
        </div>
      </vs-col>
    </vs-row>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "AddPost",
  data() {
    return {
      title: "",
      content: ""
    };
  },
  computed: {
    userToken() {
      return this.$store.state.user.token;
    }
  },
  methods: {
    createPost() {
      axios
        .post(
          "http://localhost:3000/api/v1/posts",
          {
            title: this.title,
            content: this.content.trim()
          },
          { headers: { "x-auth-token": this.userToken } }
        )
        .then(response => {
          console.log(response);
          this.$router.push("/");
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>

<style>
.add-post {
}

.add-post .add-form {
  position: absolute;
  left: 35%;
  top: 50%;
}
.add-post .title-field {
  width: 30%;
}

.add-post .content-field {
  position: relative;
  left: -35%;
  margin-top: 50px;
  resize: none;
}

.add-post .submit-button {
  margin-top: 40px;
  width: 200px;
  position: relative;
  left: 8%;
}
</style>