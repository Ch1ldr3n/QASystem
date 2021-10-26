<template>
  <el-menu mode="horizontal" :router="true">
    <el-menu-item index="1">我要提问</el-menu-item>
    <el-menu-item index="2" :route="{ name: 'issues' }">我的问题</el-menu-item>
    <el-sub-menu index="3">
        <template #title>更多</template>
        <el-menu-item index="4">我的资料</el-menu-item>
        <el-menu-item index="5">退出登录</el-menu-item>
     </el-sub-menu>
  </el-menu>
</template>

<script>
export default {
    name:"NavBar"
}
</script>

<style scoped>
</style>