<template>
  <el-menu
    mode="horizontal"
    text-align="right"
  >
    <el-menu-item index="1" v-if="loginDone===true">提问者列表</el-menu-item>
    <el-sub-menu index="2">
      <template #title>回答者列表</template>
      <el-menu-item index="2-1">item one</el-menu-item>
      <el-menu-item index="2-2">item two</el-menu-item>
      <el-menu-item index="2-3">item three</el-menu-item>
      <el-sub-menu index="2-4">
        <template #title>item four</template>
        <el-menu-item index="2-4-1">item one</el-menu-item>
        <el-menu-item index="2-4-2">item two</el-menu-item>
        <el-menu-item index="2-4-3">item three</el-menu-item>
      </el-sub-menu>
    </el-sub-menu>
    <el-menu-item index="3" style="position: absolute; right: 70px">
        <span v-on:click="signin.dialogVisible=true">登录</span>
    </el-menu-item>
    <el-menu-item index="4" style="position: absolute; right: 10px">
        <span v-on:click="register.dialogVisible=true">注册</span>
    </el-menu-item>
  </el-menu>

    <Signin v-model="signin.dialogVisible" />
    <Register v-model="register.dialogVisible" />
</template>

<script>
import Signin from './Signin.vue'
import Register from './Register.vue'
import {postsignin, postregister} from '@/utils/http.js'

export default {
    name:"NavBar",
    components: {
        Signin,
        Register,
    },
    props: {
        loginDone: {
            type: Boolean,
			default: () => false
        },
    },
    data() {
        return {
            signin: {
                dialogVisible: false,
        },
            register:{
                dialogVisible: false
        },
        }
    },
    methods:{
        signin_: function(username, password) {
            postsignin(username, password, this.token, this.translate);
            this.signin.dialogVisible = false;
        },
        register_: function(username, password){
            postregister(username, password, this.translate);
            this.register.dialogVisible = false;
        },
        quit: function () {
            this.signin.dialogVisible = false;
            this.register.dialogVisible = false;
        },
    }
}
</script>

<style scoped>
</style>