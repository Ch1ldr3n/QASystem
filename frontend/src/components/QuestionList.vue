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
      </template>
    </el-table-column>
    <el-table-column prop="date" label="日期" sortable min-width="20%" column-key="date" />
    <el-table-column prop="title" label="问题"/>
    <el-table-column prop="name" label="昵称" min-width="10%" />
    <el-table-column prop="price" label="金额" sortable min-width="10%" />
    <el-table-column prop="state" label="状态" sortable min-width="10%" />

    <el-table-column
      prop="tag"
      label="Tag"
      min-width="10%"
      :filters="[
        { text: '我提的问题', value: 'ask' },
        { text: '别人问我的问题', value: 'que' },
      ]"
      :filter-method="filterTag"
      filter-placement="bottom-end"
    >
      <template #default="scope">
        <el-tag
          :type="scope.row.tag === 'ask' ? 'warning' : 'success'"
          disable-transitions
          >{{ scope.row.tag }}</el-tag
        >
      </template>
    </el-table-column>
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
    filterTag(value, row) {
      return row.tag === value;
    },
  },
  created() {
    // TODO: filter question related to user
    fetch('/v1/question/list', {
      method: 'GET',
      // headers: {}
    })
      .then((resp) => {
        if (!resp.ok) { throw new Error('获取我的问题列表失败！'); }
        return resp.json();
      })
      .then((data) => {
        this.tableData = data.questionlist;
        console.log(data);
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
  },
};
</script>
