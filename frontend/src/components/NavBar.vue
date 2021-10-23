<template>
  <el-menu
    mode="horizontal"
    text-align="right"
  >
    <el-menu-item index="1">提问者列表</el-menu-item>
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
    <el-menu-item index="3" style="position: absolute; right: 70px">登录</el-menu-item>
    <el-menu-item index="4" style="position: absolute; right: 10px">注册</el-menu-item>
  </el-menu>
</template>

<script>
// import { defineComponent } from '@vue/composition-api'

export default {
    name:"NavBar",
    methods:{

    }
}
</script>

<style scoped>
</style>