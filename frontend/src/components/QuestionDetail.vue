<template>
  <div>
    <el-main>
      <el-container>
        <el-aside width="unset">
        </el-aside>
        <el-container style="margin-left: 50px; margin-top: 50px;">
        </el-container>
      </el-container>
    </el-main>
  </div>
</template>

<script>
import TIM from 'tim-js-sdk'
let options = {
  SDKAppID: 1400586942
};
let tim = TIM.create(options);
tim.setLogLevel(0);

export default {
  name: 'Detail',
  data() {
    return {
    }
  },
  methods: {
  },
}
</script>