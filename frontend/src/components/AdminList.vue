<template>
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
</template>

<script>
export default {
  name: 'AdminList',
  components: {
  },
  data() {
    return {
      adminlist: [],
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
  },
};
</script>
