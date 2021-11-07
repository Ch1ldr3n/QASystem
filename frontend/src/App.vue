<template>
  <el-container>
    <el-header>
      <NavBar :login-done="navbar.loginDone" />

      <!-- TODO: add admin navbar -->
    </el-header>
    <router-view />
  </el-container>
</template>

<script>
import NavBar from './components/NavBar.vue';
// import AdminNavBar from './components/AdminNavBar.vue';

export default {
  name: 'App',
  components: {
    NavBar,
    // AdminNavBar,
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
