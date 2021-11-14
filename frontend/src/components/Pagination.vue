<template>
  <div class="demo-pagination-block">
    <el-pagination
    :page-size="5"
    layout="prev, pager, next, jumper"
    v-model:currentPage="currentPage"
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
      currentPage: 1,
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
