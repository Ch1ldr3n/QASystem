<template>
  <el-main>
    <el-container>
      <QuestionList/>
    </el-container>
  </el-main>
</template>

<script>
import QuestionList from './QuestionList.vue'
export default {
  name: "Explore",
  components: {
    QuestionList,
  }
};
</script>
