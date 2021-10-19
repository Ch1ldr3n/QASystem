<template>
  <el-container>
    <el-header>
      <el-menu mode="horizontal">
        <el-menu-item index="2">付费问答</el-menu-item>

        <el-row
          v-if="!online"
          style="position: absolute; right: 10px; top: 5px"
        >
          <el-button plain v-on:click="signin.dialogVisible = true"
            >登陆</el-button
          >
          <el-button plain v-on:click="register.dialogVisible = true"
            >注册</el-button
          >
        </el-row>
        <el-row v-else style="position: absolute; right: 10px; top: 5px">
          <el-button @click="toUser" plain>个人信息</el-button>
          <el-button plain>注销</el-button>
        </el-row>
      </el-menu>
    </el-header>
    <el-main>
      <el-row justify="center">
        <el-image style="height: 500px" :src="image" :fit="fit"></el-image>
      </el-row>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-card class="box-card">
            <template #header>
              <div class="card-header">
                <span>海量题目</span>
              </div>
            </template>
            <div>
              计算机、电子、金融、医药<br />
              面面俱到
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card class="box-card">
            <template #header>
              <div class="card-header">
                <span>专业认知</span>
              </div>
            </template>
            <div>
              领域前沿学者<br />
              客观解答
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card class="box-card">
            <template #header>
              <div class="card-header">
                <span>安全可信</span>
              </div>
            </template>
            <div>
              第三方支付平台<br />
              可信结算
            </div>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
    <el-footer> </el-footer>
  </el-container>

  <Signin v-model="signin.dialogVisible" />
  <Register v-model="register.dialogVisible" />
</template>

<script>
import Signin from "./Signin.vue";
import Register from "./Register.vue";
import { useRouter } from "vue-router";
const image = require("../assets/fufeiwenda_pic1_tu.png");

export default {
  setup() {
    const router = useRouter();
    const toUser = () => {
      router.push({
        name: "user",
      });
    };
    return {
      toUser,
    };
  },
  name: "Homepage",
  components: {
    Signin,
    Register,
  },
  props: {
    msg: String,
  },
  data() {
    return {
      image,
      signin: {
        dialogVisible: false,
      },
      register: {
        dialogVisible: false,
      },
      online: true,
    };
  },
  methods: {
    register_: function () {
      this.register.dialogVisible = false;
    },
    signin_: function () {
      //username, password) {
      this.signin.dialogVisible = false;
    },
    quit: function () {
      this.signin.dialogVisible = false;
      this.register.dialogVisible = false;
    },
  },
};
</script>

<style scoped>
</style>
