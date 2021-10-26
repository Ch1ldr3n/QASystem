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
            <el-button @click="submit">向他提问</el-button>
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
        //   tag: 'ask',
          content: '1+1=2是怎么回事呢？1+1相信大家都很熟悉，但是1+1=2是怎么回事呢，下面就让小编带大家一起了解吧。1+1=2，其实就是两个一加起来就变成了二，大家可能会很惊讶1+1怎么会=2呢？但事实就是这样，小编也感到非常惊讶。这就是关于1+1=2的事情了，大家有什么想法呢，欢迎在评论区告诉小编一起讨论哦！',
        },
        {
          date: '2016-05-02',
          name: 'Tom',
          area: 'Math',
          money: '100',
        //   tag: 'que',
          content: '2+2=4是怎么回事呢？2+2相信大家都很熟悉，但是2+2=4是怎么回事呢，下面就让小编带大家一起了解吧。2+2=4，其实就是两个二加起来就变成了四，大家可能会很惊讶2+2怎么会=4呢？但事实就是这样，小编也感到非常惊讶。这就是关于2+2=4的事情了，大家有什么想法呢，欢迎在评论区告诉小编一起讨论哦！',
        },
        {
          date: '2016-05-04',
          name: 'Tom',
          area: 'Physics',
          money: '10',
        //   tag: 'ask',
          content: '3+3=6是怎么回事呢？1+1相信大家都很熟悉，但是3+3=6是怎么回事呢，下面就让小编带大家一起了解吧。3+3=6，其实就是两个三加起来就变成了六，大家可能会很惊讶3+3怎么会=6呢？但事实就是这样，小编也感到非常惊讶。这就是关于3+3=6的事情了，大家有什么想法呢，欢迎在评论区告诉小编一起讨论哦！',
        },
        {
          date: '2016-05-01',
          name: 'Tom',
          area: 'Chinese',
          money: '10000',
        //   tag: 'que',
          content: '4+4=8是怎么回事呢？1+1相信大家都很熟悉，但是4+4=8是怎么回事呢，下面就让小编带大家一起了解吧。4+4=8，其实就是两个四加起来就变成了八，大家可能会很惊讶4+4怎么会=8呢？但事实就是这样，小编也感到非常惊讶。这就是关于4+4=8的事情了，大家有什么想法呢，欢迎在评论区告诉小编一起讨论哦！',
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
        // this.$message({
        //     message: "登出成功",
        //     type: "success"
        // })
    },
  },
}
</script>
