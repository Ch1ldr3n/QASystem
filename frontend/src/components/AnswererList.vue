<template>
    <el-table
    ref="filterTable"
    :data="userlist"
    :default-sort="{ prop: 'username', order: 'descending' }"
    style="width: 100%"
  >

    <el-table-column prop="username" label="姓名" min-width="10%" column-key="name"/>
    <el-table-column prop="price" label="价格" sortable min-width="10%" />
    <el-table-column
      prop="profession"
      label="专业领域"
      min-width="10%"
      :filters="Array.from(new Set(userlist.map(x => x.profession))).map(x => { return { text: x, value: x } })"
      :filter-method="filterTag"
      filter-placement="bottom-end"
    >
      <template #default="scope">
        <el-tag
          >{{ scope.row.profession }}</el-tag
        >
      </template>
    </el-table-column>
    <el-table-column min-width="10%" >
      <template #default="scope">
        <el-button @click="ask(scope.row.id)">向他提问</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  data() {
    return {
      userlist: [],
    };
  },
  created() {
    fetch('/v1/user/filter?answerer=true', {
      method: 'GET',
      headers: { authorization: window.localStorage.getItem('token') },
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('获取回答者列表失败');
        }
        return resp.json();
      })
      .then((data) => {
        this.userlist = data.userlist;
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
  },
  methods: {
    formatter(row) {
      return row.profession;
    },
    filterTag(value, row) {
      return row.profession === value;
    },
    ask(id) {
      this.$router.push({
        name: 'Submit',
        query: { id },
      });
    },
  },
};
</script>
