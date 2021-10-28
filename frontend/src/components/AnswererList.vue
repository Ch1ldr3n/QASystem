<template>
    <el-table
    ref="filterTable"
    :data="tableData"
    :default-sort="{ prop: 'name', order: 'descending' }"
    style="width: 100%"
  >
    <el-table-column type="expand">
      <template #default="props">
        <p>{{ props.row.content }}</p>
            <el-button @click="ask">向他提问</el-button>
      </template>
    </el-table-column>
    <el-table-column prop="name" label="姓名" width="180" column-key="name"/>
    <el-table-column prop="date" label="注册日期" />
    <el-table-column prop="money" label="价格" sortable width="180" />

    <el-table-column
      prop="area"
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
      tableData: [
        {
          date: '2016-05-03',
          name: 'Tom',
          area: 'Biology',
          money: '1000',
          content: '985高校生命科学博士，丰富的留学履历和科研相关工作经验，擅长分子生物学领域',
        },
        {
          date: '2016-05-02',
          name: 'Alice',
          area: 'Math',
          money: '100',
          content: '中科院数学研究员，主要研究兴趣为微分几何学，在国际指明会议发表多篇学术论文',
        },
        {
          date: '2016-05-04',
          name: 'Bob',
          area: 'Physics',
          money: '10',
          content: '拥有20年从业经验的重点中学物理教师，擅长与学生发展良好的师生关系，课堂讲解深入浅出',
        },
        {
          date: '2016-05-01',
          name: 'Jack',
          area: 'Chinese',
          money: '1600',
          content: '中国文学学士，长期从事中西方文化交流宣传活动，著有文集《清华的那段又冷又饿又困的时光》等',
        },
      ],
    }
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
