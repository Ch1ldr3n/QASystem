<template>
  <el-container>
    <el-header>
      <el-input
        v-model="newname"
        placeholder="管理员昵称"
        clearable
      >
        <template #append>
          <el-button @click="adding">
            添加
          </el-button>
        </template>
      </el-input>
    </el-header>
    <el-table
      ref="filterTable"
      :data="tableData"
      :default-sort="{ prop: 'date', order: 'descending' }"
      style="width: 100%"
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
        prop="date"
        label="日期"
        sortable
        min-width="20%"
        column-key="date"
      />
      <el-table-column
        prop="title"
        label="问题"
      />

      <el-table-column
        prop="state"
        label="状态"
        sortable
        min-width="10%"
      />
    </el-table>
  </el-container>
</template>

<script>
export default {
  name: 'AdminQuestionList',
  data() {
    return {
      questionlist: [],
      newname: '',
    };
  },
  created() {
    fetch('/v1/question/list', {
      method: 'GET',
      // headers: {}
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('获取管理员问题列表失败！');
        }
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
  methods: {
    filterTag(value, row) {
      return row.tag === value;
    },
    adding() {
      this.$message({
        message: '添加成功',
        type: 'success',
      });
    },
  },
};
</script>
