<template>
  <div class="demo-pagination-block">
    <el-pagination
      v-model:currentPage="currentPage"
      :page-size="5"
      layout="prev, pager, next, jumper"
      :total="100"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>
<script>
export default ({
  name: 'Pagination',
  data() {
    return {
      currentPage: 1,
    };
  },
  methods: {
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`);
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
    //   this.data = [];
    },
    getlist() {
    },
  },
});
</script>
