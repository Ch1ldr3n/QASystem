<template>
  <el-main>
    <el-container>
      <el-aside width="unset">
        <MyInfoCard/>
      </el-aside>
      <el-container style="margin-left: 50px; margin-top: 20px;">
        <AnswererList/>
      </el-container>
    </el-container>
  </el-main>
</template>

<script>
import AnswererList from './AnswererList.vue'
import MyInfoCard from './MyInfoCard.vue'

export default {
  name: "answerer",
  components: {
    AnswererList,
    MyInfoCard,
  },
  data() {
    return
  },
  methods:{
  }
};
</script>
