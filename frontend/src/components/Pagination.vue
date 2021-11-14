<template>
  <div class="demo-pagination-block">
    <span class="demonstration">前往</span>
    <el-pagination
    :page-size="5"
    layout="prev, pager, next, jumper"
    v-model:currentPage="currentPage1"
    :total="100"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
    >
    </el-pagination>
  </div>
</template>
<script>
export default ({
  name: 'Pagination',
  data() {
    return {
      currentPage1: 1,
      currentPage2: 5,
      currentPage3: 5,
      cuurentPage4: 4,
    };
  },
  methods: {
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`);
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
    },
  },
});
</script>
