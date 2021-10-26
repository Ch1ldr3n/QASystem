<template>
  <el-menu mode="horizontal" router="true">
    <el-menu-item index="1">我要提问</el-menu-item>
    <el-menu-item index="2" :route="{ name: 'issues' }">我的问题</el-menu-item>
    <el-menu-item index="3">
        <span v-on:click="signin.dialogVisible=true">登录</span>
    </el-menu-item>
    <el-menu-item index="4">
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