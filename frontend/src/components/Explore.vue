<template>
  <el-main>
    <el-container>
      <ExploreList/>
    </el-container>
  </el-main>
</template>

<script>
import ExploreList from './ExploreList.vue'
export default {
  name: "Explore",
  components: {
    ExploreList,
  }
};
</script>
