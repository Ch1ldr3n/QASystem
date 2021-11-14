<template>
  <el-main>
    <el-container>
      <ExploreList />
    </el-container>
    <!-- <el-contatiner style="text-align: center;">
      <Pagination />
    </el-contatiner> -->
  </el-main>
</template>

<script>
import ExploreList from './ExploreList.vue';
// import Pagination from './Pagination.vue';
// import ExploreList2 from './ExploreList_2.vue';

export default {
  name: 'Explore',
  components: {
    ExploreList,
    // ExploreList2,
    // Pagination,
  },
};
</script>
