<template>
  <el-main>
    <el-container>
      <el-aside width="unset">
        <el-card class="box-card">
          <div class="card-header">
            <el-image :src="'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png'"/>
          </div>
          <div>
            <el-descriptions direction="vertical" column="1">
              <el-descriptions-item label="用户名">测试用户</el-descriptions-item>
              <el-descriptions-item label="手机号">18100000000</el-descriptions-item>
              <el-descriptions-item label="称号"><el-tag size="small">回答者</el-tag></el-descriptions-item>
              <el-descriptions-item label="余额">999999.99</el-descriptions-item>
            </el-descriptions>
          </div>
        </el-card>
      </el-aside>
      <el-container>
        <AnswererList/>
      </el-container>
    </el-container>
  </el-main>
</template>

<script>
import AnswererList from './AnswererList.vue'

export default {
  name: "answerer",
  components: {
    AnswererList,
  },
  data() {
    return
  },
  methods:{
  }
};
</script>
