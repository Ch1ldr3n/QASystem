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
        <el-button @click="submit">查看详情</el-button>
      </template>
    </el-table-column>
    <el-table-column prop="date" label="日期" sortable min-width="20%" column-key="date" />
    <el-table-column prop="question" label="问题" :formatter="formatter" />
    <el-table-column prop="askname" label="提问者" min-width="10%" />
    <el-table-column prop="quename" label="回答者" min-width="10%" />
  </el-table>
</template>

<script>
export default {
  data() {
    return {
      tableData: [
        {
          date: '2016-05-03',
          askname: '张三',
          quename: '李四',
          question: '什么是软件工程',
          content: '每个人都不得不面对这些问题。 在面对这种问题时， 我们不得不面对一个非常尴尬的事实，那就是， 生活中，若软件工程出现了，我们就不得不考虑它出现了的事实。 软件工程因何而发生?我们一般认为，抓住了问题的关键，其他一切则会迎刃而解。 每个人都不得不面对这些问题。 ',
        },
        {
          date: '2016-05-02',
          askname: '李四',
          quename: '王五',
          question: '嘉然今天吃什么',
          content: '一般来说， 这种事实对本人来说意义重大，相信对这个世界也是有一定意义的。 现在，解决嘉然今天吃什么的问题，是非常非常重要的。 所以， 现在，解决嘉然今天吃什么的问题，是非常非常重要的。 所以， 那么， 那么， 嘉然今天吃什么，发生了会如何，不发生又会如何。 现在，解决嘉然今天吃什么的问题，是非常非常重要的。 所以， 生活中，若嘉然今天吃什么出现了，我们就不得不考虑它出现了的事实。 我们一般认为，抓住了问题的关键，其他一切则会迎刃而解。 要想清楚，嘉然今天吃什么，到底是一种怎么样的存在。 经过上述讨论我们一般认为，抓住了问题的关键，其他一切则会迎刃而解。 现在，解决嘉然今天吃什么的问题，是非常非常重要的',
        },
        {
          date: '2016-05-04',
          askname: '王五',
          quename: '赵六',
          question: '什么是测试',
          content: '阿卜·日·法拉兹曾经说过，学问是异常珍贵的东西，从任何源泉吸收都不可耻。我希望诸位也能好好地体会这句话。 那么， 本人也是经过了深思熟虑，在每个日日夜夜思考这个问题。 一般来讲，我们都必须务必慎重的考虑考虑。 测试的发生，到底需要如何做到，不测试的发生，又会如何产生。 测试，发生了会如何，不发生又会如何。 测试的发生，到底需要如何做到，不测试的发生，又会如何产生。 测试因何而发生?我们都知道，只要有意义，那么就必须慎重考虑。 这种事实对本人来说意义重大，相信对这个世界也是有一定意义的。 总结的来说， 我们都知道，只要有意义，那么就必须慎重考虑。 总结的来说， 测试因何而发生?一般来说， 就我个人来说，测试对我的意义，不能不说非常重大。',
        },
        {
          date: '2016-05-01',
          askname: '赵六',
          quename: '张三',
          question: '什么是后端',
          content: '而这些并不是完全重要，更加重要的问题是， 我们不得不面对一个非常尴尬的事实，那就是， 既然如何， 那么， 裴斯泰洛齐在不经意间这样说过，今天应做的事没有做，明天再早也是耽误了。我希望诸位也能好好地体会这句话。 我们不得不面对一个非常尴尬的事实，那就是， 一般来说， 而这些并不是完全重要，更加重要的问题是， 就我个人来说，后端对我的意义，不能不说非常重大。 后端，到底应该如何实现。 我们一般认为，抓住了问题的关键，其他一切则会迎刃而解。',
        },
      ],
    };
  },
  methods: {
    formatter(row) {
      return row.question;
    },
  },
};
</script>
