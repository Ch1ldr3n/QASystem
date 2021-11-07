<template>
  <el-main>
    <el-container style="margin-left: 50px; margin-top: 20px; margin-right: 50px;">
      <AdminQuestionList />
    </el-container>
  </el-main>
</template>

<script>
import AdminQuestionList from './AdminQuestionList.vue';
import AdminNavBar from './AdminNavBar.vue';

export default {
  name: 'AdminQuestion',
  components: {
    AdminQuestionList,
    AdminNavBar,
  },
  data() {
    return {};
  },
  methods: {},
};
</script>
