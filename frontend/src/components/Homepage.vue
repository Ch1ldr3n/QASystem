<template>
  <div>
    <el-container>
      <el-header style="text-align: right">
        <NavBar/>
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

  </div>
</template>

<script>
import NavBar from './NavBar.vue'
import { useRouter } from "vue-router";

const image = require('../assets/fufeiwenda_pic1_tu.png')

export default {
  setup() {
    const router = useRouter();
    const toUser = () => {
      console.log("USER");
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
    NavBar,
  },
  props: {
    msg: String,
  },
  data() {
    return {
      image,
      tolen: ""
    }
  },
  methods:{
    translate: function(response) {
      if (response.status == 200) {
        window.localStorage.setItem("token", response.data.token);
      }
    },
  },
};
</script>

<style scoped>
/* .el-header {
        text-align: right;
        width: 100%;
        height: 100px;      
    } */
</style>
