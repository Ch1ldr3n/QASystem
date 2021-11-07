<template>
  <el-container>
    <el-header>
      <NavBar :login-done="navbar.loginDone" />
    </el-header>
    <router-view />
  </el-container>
</template>

<script>
import NavBar from './components/NavBar.vue';

export default {
  name: 'App',
  components: {
    NavBar,
  },
  data() {
    return {
      navbar: {
        loginDone: false,
      },
    };
  },
};
</script>

<style></style>
