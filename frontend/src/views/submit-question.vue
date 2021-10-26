<template>
  <div>
    <el-main>
      <el-container>
        <el-aside width="unset">
          <el-card class="box-card">
            <div class="card-header">
              <el-image :src="'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png'"/>
            </div>
            <div>
              <el-descriptions direction="vertical" column="1">
                <el-descriptions-item label="用户名">测试用户</el-descriptions-item>
                <el-descriptions-item label="手机号">18100000000</el-descriptions-item>
                <el-descriptions-item label="称号"><el-tag size="small">回答者</el-tag></el-descriptions-item>
                <el-descriptions-item label="余额">999999.99</el-descriptions-item>
              </el-descriptions>
            </div>
          </el-card>
        </el-aside>
        <el-container>
        <el-form label-width="120px">
          <el-form-item label="问题名称">
            <el-input v-model="qname"></el-input>
          </el-form-item>
          <el-form-item label="问题描述">
            <el-input type="textarea" v-model="qdesc"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">提交问题</el-button>
            <el-button>取消提问</el-button>
          </el-form-item>
        </el-form>
        </el-container>
      </el-container>
      </el-main>
  </div>
</template>

<script></script>

<style scoped></style>

<script>
export default {
  name: "pay",
  data() {
    return {
      qname: "",
      qdesc: "",
    };
  },
};
</script>

<style scoped>
.el-aside {
  height: 100%;
}
</style>
