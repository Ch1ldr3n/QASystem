<template>
  <el-main>
    <el-container>
      <el-aside width="unset">
        <MyInfoCard />
      </el-aside>
      <el-container style="margin-left: 50px; margin-top: 20px;">
        <el-table
          ref="filterTable"
          :data="list"
          :default-sort="{ prop: 'date', order: 'descending' }"
          style="width: 100%"
        >
          <el-table-column
            prop="month"
            label="月份"
            sortable
            min-width="10%"
            column-key="month"
          />
          <el-table-column
            prop="income"
            label="收入"
            sortable
            min-width="20%"
          />
        </el-table>
      </el-container>
    </el-container>
  </el-main>
</template>

<script>
import MyInfoCard from './MyInfoCard.vue';

export default {
  name: 'Income',
  components: {
    MyInfoCard,
  },
  data() {
    return {
      tableData: [],
      list: [],
    };
  },
  created() {
    fetch('/v1/question/list', {
      method: 'GET',
    })
      .then((resp) => {
        if (!resp.ok) { throw new Error('获取我的问题列表失败！'); }
        return resp.json();
      })
      .then((data) => {
        this.tableData = data.questionlist;
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
    this.tableData.forEach(function (item) {
      if (item.tag === 'que') {
        const date = this.list.get(item.date);
        const money = item.price;
        if (date === null) {
          this.list.push({ month: date, income: money });
        } else {
          this.list.forEach((value) => {
            if (value.date === item.date) {
              // eslint-disable-next-line no-param-reassign
              value.income += money;
            }
          });
        }
      }
    });
  },
  methods: {

  },
};
</script>
