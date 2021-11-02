<template>
  <el-main>
    <el-container>
      <el-aside width="unset">
        <MyInfoCard />
      </el-aside>
      <el-container style="margin-left: 50px; margin-top: 20px">
        <AdminQuestionList />
      </el-container>
    </el-container>
  </el-main>
</template>

<script>
import AdminQuestionList from './AdminQuestionList.vue'
import MyInfoCard from './MyInfoCard.vue'

export default {
  name: 'question',
  components: {
    AdminQuestionList,
    MyInfoCard,
  },
  data () {
    return {
    }
  },
  methods: {
  },
};
</script>
