<template>
  <el-container>
    <el-header>
      <NavBar :loginDone="navbar.loginDone"/>
    </el-header>
    <router-view/>
  </el-container>
</template>

<script>
import NavBar from './components/NavBar.vue'

export default {
  name: "App",
  data() {
    return {
      navbar: {
        loginDone: false,
      }
    }
  },
  components: {
    NavBar,
  },
}
</script>

<style></style>
