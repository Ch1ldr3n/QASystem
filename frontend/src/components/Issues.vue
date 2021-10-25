<template>
  <div>
    <el-container>
      <el-header>Header</el-header>
      <el-container>
        <el-aside width="200px">Aside</el-aside>
        <el-main>
          <Issueslist/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import Issueslist from './Issueslist.vue'

export default {
  name: "Homepage",
  components: {
    Issueslist,
  },
  data() {
    return
  },
  methods:{
  }
};
</script>
