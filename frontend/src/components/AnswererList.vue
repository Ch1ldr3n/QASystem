<template>
    <el-table
    ref="filterTable"
    :data="userlist"
    :default-sort="{ prop: 'username', order: 'descending' }"
    style="width: 100%"
  >
    <el-table-column type="expand">
      <template #default="props">
        <p>{{ props.row.content }}</p>
            <el-button @click="ask">向他提问</el-button>
      </template>
    </el-table-column>
    <el-table-column prop="username" label="姓名" width="180" column-key="name"/>
    <el-table-column prop="price" label="价格" sortable width="180" />

    <el-table-column
      prop="profession"
      label="专业领域"
      width="100"
      :filters="[
        { text: '生命科学', value: 'Biology' },
        { text: '计算机科学', value: 'CS' },
        { text: '数学', value: 'Math'},
        { text: '物理学', value: 'Physics'},
        { text: '中国文学', value: 'Chinese'},
        { text: '英语', value: 'English'},
      ]"
      :filter-method="filterTag"
      filter-placement="bottom-end"
    >
      <template #default="scope">
        <el-tag
          >{{ scope.row.area }}</el-tag
        >
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  data() {
    return {
      userlist: []
    }
  },
  created() {
    fetch("/v1/user/filter?answerer=true", {
      method: "GET",
      headers: {"authorization": window.localStorage.getItem("token")},
    })
    .then(resp => {
      if (!resp.ok) {
        throw new Error("获取回答者列表失败")
      }
      return resp.json()
    })
    .then(data => {
      this.userlist = data.userlist
      console.log(data)
    })
    .catch(error => {
      this.$message({
        message: error,
        type: "error",
      })
    })
  },
  methods: {
    formatter(row) {
      return row.area
    },
    filterTag(value, row) {
      return row.area === value
    },
    ask() {
        this.$router.push({
            name: "Submit"
        })
    },
  },
}
</script>
