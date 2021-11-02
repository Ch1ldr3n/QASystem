<template>
  <el-card class="box-card">
    <div class="card-header">
      <el-image
        :src="
          'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png'
        "
      />
    </div>
    <div>
      <el-descriptions :column="1">
        <el-descriptions-item label="用户名">{{
          username
        }}</el-descriptions-item>
        <el-descriptions-item label="称号"
          ><el-tag size="small">{{ answerer }}</el-tag></el-descriptions-item
        >
        <el-descriptions-item label="账户余额">{{}}</el-descriptions-item>
      </el-descriptions>
    </div>
  </el-card>
</template>

<script>
export default {
  data() {
    return {
      answerer: '',
      username: '',
    };
  },
  created() {
    fetch('/v1/user/info', {
      method: 'GET',
      headers: { authorization: window.localStorage.getItem('token') },
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('获取用户信息失败');
        }
        return resp.json();
      })
      .then((data) => {
        this.username = data.username;
        this.answerer = data.answerer ? '回答者' : '提问者';
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
  },
};
</script>
