<template>
  <el-main>
    <el-container>
      <el-aside width="unset">
        <MyInfoCard/>
      </el-aside>
      <el-container style="margin-left: 50px; margin-top: 20px;">
        <QuestionList/>
      </el-container>
    </el-container>
  </el-main>
</template>

<script>
import QuestionList from './QuestionList.vue'
import MyInfoCard from './MyInfoCard.vue'

export default {
  name: "question",
  components: {
    QuestionList,
    MyInfoCard,
  },
  data() {
    return {
    }
  },
  methods:{
  },
};
</script>
