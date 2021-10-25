<template>
  <div>
    <el-container>
      <el-aside width="400px" style="background-color: rgb(238, 241, 246)">
        <el-image
          style="width: 100; height: 100"
          :src="require('../assets/logo.png')"
          :fit="fill"
        >
        </el-image>
        <h1>name</h1>
        <p>descriptiondescriptiondescriptiondescription</p>
        <p>
          descriptiondescriptiondesc
        </p>
        <p>descriptiondescriptiondescription</p>
        <p>descriptiondescriptiondescription</p>
        <p>descriptiondescriptiondescription</p>
        <p>descriptiondescriptiondescription</p>
        <p>descriptiondescriptiondescription</p>
        <p>descriptiondescriptiondescription</p>
      </el-aside>
      <el-main>
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
      </el-main>
    </el-container>
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
