<template>
  <el-container>
    <el-header>
      <el-menu mode="horizontal">
        <el-menu-item index="1">付费问答</el-menu-item>
      </el-menu>
    </el-header>
    <el-main>
      <el-row justify="center">
        <el-image
        style="height: 500px"
        :src="image"
        :fit="fit"
        ></el-image>
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
              计算机、电子、金融、医药<br>
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
              领域前沿学者<br>
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
              第三方支付平台<br>
              可信结算
            </div>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
    <el-footer>
    </el-footer>
  </el-container>
</template>

<script>
const image = require('../assets/fufeiwenda_pic1_tu.png')

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data() {
    return {
      image
    }
  }
}
</script>

<style scoped>
</style>
