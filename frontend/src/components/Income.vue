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
            label="年份"
            min-width="10%"
          />
          <el-table-column
            prop="month"
            label="月份"
            sortable
            min-width="10%"
            column-key="month"
          />
          <el-table-column
            prop="earning"
            label="收入"
            sortable
            min-width="20%"
          />
          <el-table-column
            prop="spending"
            label="支出"
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
      list: [],
    };
  },
  created() {
    fetch('/v1/question/aggreg', {
      method: 'GET',
      headers: { authorization: window.localStorage.getItem('token') },
    })
      .then((resp) => {
        if (!resp.ok) { throw new Error('获取收入统计失败！'); }
        return resp.json();
      })
      .then((data) => {
        this.list = data.list;
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
  },
  methods: {

  },
};
</script>
