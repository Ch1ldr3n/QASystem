<template>
  <el-table
    ref="filterTable"
    :data="tableData"
    :default-sort="{ prop: 'date', order: 'descending' }"
    style="width: 100%"
  >
    <el-table-column type="expand">
      <template #default="props">
        <p>{{ props.row.content }}</p>
        <el-button @click="submit">
          查看详情
        </el-button>
      </template>
    </el-table-column>
    <el-table-column prop="title" label="问题" />
    <el-table-column prop="qusername" label="提问者" min-width="10%" />
    <el-table-column prop="ausername" label="回答者" min-width="10%" />
  </el-table>
</template>

<script>
export default {
  data() {
    return {
      tableData: [],
    };
  },
  methods: {
    submit() {
      this.$message({
        message: '尚未实现！',
      });
    },
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
        message: '登录后查看问题详情！',
      });
    },
  },
};
</script>
