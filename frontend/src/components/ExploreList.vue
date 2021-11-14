<template>
<el-main>
  <el-container>
  <el-table
    ref="filterTable"
    :data="tableData.slice((currentPage-1)*pageSize,currentPage*pageSize)"
    :default-sort="{ prop: 'date', order: 'descending' }"
    style="width: 100%"
    stripe
  >
    <el-table-column type="expand">
      <template #default="props">
        <p>{{ props.row.content }}</p>
        <el-button @click="submit">
          查看详情
        </el-button>
      </template>
    </el-table-column>
    <el-table-column
      prop="title"
      label="问题"
    />
    <el-table-column
      prop="qusername"
      label="提问者"
      min-width="10%"
    />
    <el-table-column
      prop="ausername"
      label="回答者"
      min-width="10%"
    />
  </el-table>
  </el-container>
  <el-container style="text-align: center;">
    <el-pagination
    :page-size="10"
    layout="prev, pager, next, jumper"
    v-model:currentPage="currentPage"
    :total="1000"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
    style="margin:0 auto"
    >
    </el-pagination>
  </el-container>
</el-main>
</template>

<script>
export default {
  data() {
    return {
      tableData: [],
      currentPage: 1,
      pageSize: 10,
    };
  },
  created() {
    // TODO: ask backend to return more related information
    fetch('/v1/question/list', {
      method: 'GET',
      headers: {},
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('获取知识广场信息失败！');
        }
        return resp.json();
      })
      .then((data) => {
        this.tableData = data.questionlist;
        console.log(data);
      })
      .catch((error) => {
        this.$message({ message: error, type: 'error' });
      });
  },
  methods: {
    submit() {
      this.$message({
        message: '尚未实现！',
      });
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`);
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
    },
  },
};
</script>
