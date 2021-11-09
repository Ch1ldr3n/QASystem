<template>
  <el-container style="margin-left: 50px; margin-top: 20px;">
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
    <el-main>
      <el-table
        ref="filterTable"
        :data="adminlist"
        :default-sort="{ prop: 'username', order: 'descending' }"
        style="width: 100%"
      >
        <el-table-column type="expand">
          <template #default="props">
            <p>{{ props.row.content }}</p>
          </template>
        </el-table-column>
        <el-table-column
          prop="username"
          label="姓名"
        />
        <el-table-column
          prop="role"
          label="职务"
          width="180"
        />
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
export default {
  name: 'AdminList',
  components: {
  },
  data() {
    return {
      adminlist: [],
      newname: '',
    };
  },
  created() {
    this.refresh();
  },
  methods: {
    refresh() {
      fetch('/v1/admin/list', {
        method: 'GET',
        headers: {
          'content-type': 'application/json',
          authorization: window.localStorage.getItem('admintoken'),
        },
      })
        .then((resp) => {
          if (!resp.ok) {
            throw new Error('刷新管理员列表失败！');
          }
          return resp.json();
        })
        .then((data) => {
          this.adminlist = data.adminlist;
          console.log(data);
        })
        .catch((error) => {
          this.$message({
            message: error,
            type: 'error',
          });
        });
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
  },
};
</script>
