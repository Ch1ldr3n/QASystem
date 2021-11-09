<template>
  <el-container>
    <el-table
      ref="filterTable"
      :data="tableData"
      :default-sort="{ prop: 'date', order: 'descending' }"
      style="width: 100%"
      border
      stripe
    >
      <el-table-column type="expand">
        <template #default="props">
          <p>{{ props.row.content }}</p>
          <el-button
            type="primary"
          >
            审核通过
          </el-button>
          <el-button>
            拒绝通过
          </el-button>
        </template>
      </el-table-column>
      <el-table-column
        prop="title"
        label="问题"
      />
      <el-table-column
        prop="question"
        label="提问者"
        width="180"
      />

      <el-table-column
        prop="answer"
        label="回答者"
        width="180"
      />
    </el-table>
  </el-container>
</template>

<script>
export default {
  name: 'Review',
  data() {
    return {
      reviewlist: [],
    };
  },
  created() {
    fetch('/v1/question/reviewlist', {
      method: 'GET',
      // headers: {}
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('获取待审核问题列表失败！');
        }
        return resp.json();
      })
      .then((data) => {
        this.tableData = data.reviewlist;
        console.log(data);
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
  },
  methods: {
    filterTag(value, row) {
      return row.tag === value;
    },
  },
};
</script>
