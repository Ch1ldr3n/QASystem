<template>
  <el-table
    ref="filterTable"
    :data="tableData"
    :default-sort="{ prop: 'date', order: 'descending' }"
    style="width: 100%"
  >
    <el-table-column type="expand">
      <template #default="props">
        <p>{{ props.row.content }}</p>
      </template>
    </el-table-column>
    <el-table-column prop="date" label="日期" sortable min-width="20%" column-key="date" />

    <el-table-column prop="question" label="问题" :formatter="formatter" />
    <el-table-column prop="name" label="昵称" min-width="10%" />
    <el-table-column prop="money" label="金额" sortable min-width="10%" />

    <el-table-column
      prop="tag"
      label="Tag"
      min-width="10%"
      :filters="[
        { text: '我提的问题', value: 'ask' },
        { text: '别人问我的问题', value: 'que' },
      ]"
      :filter-method="filterTag"
      filter-placement="bottom-end"
    >
      <template #default="scope">
        <el-tag
          :type="scope.row.tag === 'ask' ? 'warning' : 'success'"
          disable-transitions
          >{{ scope.row.tag }}</el-tag
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
          name: '李四',
          money: '1000',
          tag: 'ask',
          question: '什么是软件工程',
          content: '每个人都不得不面对这些问题。 在面对这种问题时， 我们不得不面对一个非常尴尬的事实，那就是， 生活中，若软件工程出现了，我们就不得不考虑它出现了的事实。 软件工程因何而发生?我们一般认为，抓住了问题的关键，其他一切则会迎刃而解。 每个人都不得不面对这些问题。 ',
        },
        {
          date: '2016-05-02',
          name: '赵六',
          money: '100',
          tag: 'que',
          question: '什么是后端',
          content: '而这些并不是完全重要，更加重要的问题是， 我们不得不面对一个非常尴尬的事实，那就是， 既然如何， 那么， 裴斯泰洛齐在不经意间这样说过，今天应做的事没有做，明天再早也是耽误了。我希望诸位也能好好地体会这句话。 我们不得不面对一个非常尴尬的事实，那就是， 一般来说， 而这些并不是完全重要，更加重要的问题是， 就我个人来说，后端对我的意义，不能不说非常重大。 后端，到底应该如何实现。 我们一般认为，抓住了问题的关键，其他一切则会迎刃而解。',
        },
      ],
    }
  },
  methods: {
    formatter(row) {
      return row.question
    },
    filterTag(value, row) {
      return row.tag === value
    },
  },
}
</script>
