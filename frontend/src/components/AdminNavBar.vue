<template>
  <el-menu
    mode="horizontal"
    :router="true"
  >
    <el-menu-item
      index="0"
      :route="{ name: 'Review' }"
    >
      待审核问题
    </el-menu-item>
    <el-menu-item
      index="1"
      :route="{ name: 'AdminList' }"
    >
      管理员目录
    </el-menu-item>
    <el-sub-menu index="3">
      <template #title>
        更多
      </template>
      <el-menu-item
        index="4"
        :route="{ name: 'AdminPasswordChange' }"
      >
        修改密码
      </el-menu-item>
      <el-menu-item
        index="5"
        :route="{ name: 'SystemConfig' }"
      >
        系统参数
      </el-menu-item>
      <el-menu-item
        index="6"
        @click="logout"
      >
        退出登录
      </el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script>
export default {
  name: 'AdminNavBar',
  methods: {
    logout() {
      window.localStorage.removeItem('token');
      this.$message({
        message: '登出成功',
        type: 'success',
      });
      this.$router.push({
        name: 'Explore',
      });
    },
  },
};
</script>

<style scoped></style>
