<template>
  <div>
    <el-container>
      <el-main>
        <el-row justify="center">
          <el-col :span="16">
            <el-alert
              title="扫码支付"
              type="info"
              description="请扫描右方二维码进行支付"
              show-icon
            />
          </el-col>
        </el-row>

        <el-row :gutter="2" class="paybox" justify="center">
          <el-col :span="12">
            <el-image
              style="width: 50%; height: 100%"
              :src="require('../assets/pay/pay-demo.png')"
              :fit="contain"
            >
            </el-image>
          </el-col>
          <el-col :span="12">
            <el-image
              style="width: 100; height: 100"
              :src="require('../assets/pay/pay-qrscan.png')"
              :fit="fill"
            >
            </el-image>
            <p>支付已完成，点击这里跳转</p>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "pay",
};
</script>

<style scoped>
.el-row {
  margin: 50px auto;
}
.pay-box {
  height: 500px;
}
</style>
