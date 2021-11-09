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
            @click="approve(props.row.id, true)"
          >
            审核通过
          </el-button>
          <el-button
          @click="approve(props.row.id, false)">
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
      reviewlist: [
        {
          title: 'weishenme',
          questioner: 'cxj',
          answer: 'zmy',
        },
      ],
      newname: '',
    };
  },
  created() {
    this.refresh();
  },
  methods: {
    refresh() {
      fetch('/v1/question/reviewlist', {
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
    filterTag(value, row) {
      return row.tag === value;
    },
    adding() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          fetch('/v1/admin/add', {
            method: 'POST',
            headers: {
              Authorization: window.localStorage.getItem('admintoken'),
              'content-type': 'application/json',
            },
            body: JSON.stringify({
              token: window.localStorage.getItem('admintoken'),
              username: this.newname,
            }),
          }).then((resp) => {
            if (!resp.ok) {
              throw new Error('添加失败!');
            }
          }).then((data) => {
            this.$message({
              showClose: true,
              message: `添加成功，初始密码为 ${data.password}`,
              type: 'success',
              duration: 0,
            });
            this.$forceUpdate();
          });
        }
      });
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
