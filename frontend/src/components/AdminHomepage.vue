<template>
  <el-main>
    <el-container style="margin-left: 50px; margin-top: 20px; margin-right: 50px;">
      <Review />
    </el-container>
  </el-main>
</template>

<script>
import Review from './Review.vue';

export default {
  name: 'AdminHomepage',
  components: {
    Review,
  },
  data() {
    return {};
  },
  methods: {},
};
</script>
