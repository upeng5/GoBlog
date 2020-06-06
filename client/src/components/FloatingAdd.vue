<template>
  <div class="floating-add">
    <vs-button circle icon floating color="#ffba00" class="f-button" @click="loadAddPostPage">
      <i class="material-icons">add</i>
    </vs-button>
  </div>
</template>

<script>
export default {
  name: "FloatingAdd",
  methods: {
    loadAddPostPage() {
      this.$router.push("/add-post");
    }
  }
};
</script>

<style>
.floating-add .f-button {
  position: fixed;
  width: 50px;
  height: 50px;
  top: 85%;
  left: 92%;
}
</style>
