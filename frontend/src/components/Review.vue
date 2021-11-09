<template>
  <el-container>
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
            @click="approve(props.row.id, true)"
          >
            审核通过
          </el-button>
          <el-button
            @click="approve(props.row.id, false)"
          >
            拒绝通过
          </el-button>
        </template>
      </el-table-column>
      <el-table-column
        prop="title"
        label="问题"
      />
      <el-table-column
        prop="questioner"
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
    this.refresh();
  },
  methods: {
    refresh() {
      fetch('/v1/question/review', {
        method: 'GET',
        headers: {
          'content-type': 'application/json',
          authorization: window.localStorage.getItem('admintoken'),
        },
      })
        .then((resp) => {
          if (!resp.ok) {
            throw new Error('刷新待审核问题列表失败！');
          }
          return resp.json();
        })
        .then((data) => {
          this.reviewlist = data;
          console.log(data);
        })
        .catch((error) => {
          this.$message({
            message: error,
            type: 'error',
          });
        });
    },
    filterTag(value, row) {
      return row.tag === value;
    },
    approve(id, choice) {
      fetch('/v1/question/review', {
        method: 'POST',
        headers: {
          'content-type': 'application/json',
          authorization: window.localStorage.getItem('admintoken'),
        },
        body: JSON.stringify({
          choice,
          questionid: id,
        }),
      }).then((resp) => {
        if (!resp.ok) {
          throw new Error('审核通过失败！');
        }
      }).then(() => {
        this.$message({
          message: '确认成功',
          type: 'success',
        });
        this.refresh();
      }).catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
    },
  },
};
</script>
