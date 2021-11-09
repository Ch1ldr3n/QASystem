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
      :data="reviewlist"
      :default-sort="{ prop: 'date', order: 'descending' }"
      style="width: 100%"
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
      newname: '',
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
        this.reviewlist = data.reviewlist;
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
