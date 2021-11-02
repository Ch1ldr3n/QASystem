<template>
  <el-main>
    <el-container style="margin-left: 50px; margin-top: 20px; margin-right: 50px;">
      <AdminQuestionList />
    </el-container>
  </el-main>
</template>

<script>
import AdminQuestionList from './AdminQuestionList.vue';

export default {
  name: 'AdminQuestion',
  components: {
    AdminQuestionList,
  },
  data() {
    return {};
  },
  methods: {},
};
</script>
