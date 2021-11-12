<template>
  <el-main>
    <el-container>
      <ExploreList2 />
    </el-container>
  </el-main>
</template>

<script>
// import ExploreList from './ExploreList.vue';
import ExploreList2 from './ExploreList_2.vue';

export default {
  name: 'Explore',
  components: {
    // ExploreList,
    ExploreList2,
  },
};
</script>
