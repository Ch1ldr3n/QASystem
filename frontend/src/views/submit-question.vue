<template>
  <div>
    <el-container>
      <el-aside width="400px" style="background-color: rgb(238, 241, 246)">
        <el-image
          style="width: 100; height: 100"
          :src="require('../assets/logo.png')"
          :fit="fill"
        >
        </el-image>
        <el-menu :default-openeds="['1', '3']">
          <el-sub-menu index="1">
            <template #title
              ><i class="el-icon-message"></i>Navigator One</template
            >
            <el-menu-item-group>
              <template #title>Group 1</template>
              <el-menu-item index="1-1">Option 1</el-menu-item>
              <el-menu-item index="1-2">Option 2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="Group 2">
              <el-menu-item index="1-3">Option 3</el-menu-item>
            </el-menu-item-group>
            <el-sub-menu index="1-4">
              <template #title>Option4</template>
              <el-menu-item index="1-4-1">Option 4-1</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
          <el-sub-menu index="2">
            <template #title
              ><i class="el-icon-menu"></i>Navigator Two</template
            >
            <el-menu-item-group>
              <template #title>Group 1</template>
              <el-menu-item index="2-1">Option 1</el-menu-item>
              <el-menu-item index="2-2">Option 2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="Group 2">
              <el-menu-item index="2-3">Option 3</el-menu-item>
            </el-menu-item-group>
            <el-sub-menu index="2-4">
              <template #title>Option 4</template>
              <el-menu-item index="2-4-1">Option 4-1</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
          <el-sub-menu index="3">
            <template #title
              ><i class="el-icon-setting"></i>Navigator Three</template
            >
            <el-menu-item-group>
              <template #title>Group 1</template>
              <el-menu-item index="3-1">Option 1</el-menu-item>
              <el-menu-item index="3-2">Option 2</el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group title="Group 2">
              <el-menu-item index="3-3">Option 3</el-menu-item>
            </el-menu-item-group>
            <el-sub-menu index="3-4">
              <template #title>Option 4</template>
              <el-menu-item index="3-4-1">Option 4-1</el-menu-item>
            </el-sub-menu>
          </el-sub-menu>
        </el-menu>
      </el-aside>
      <el-main>
        <el-form label-width="120px">
          <el-form-item label="问题名称">
            <el-input v-model="qname"></el-input>
          </el-form-item>
          <el-form-item label="问题描述">
            <el-input type="textarea" v-model="qdesc"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">提交问题</el-button>
            <el-button>取消提问</el-button>
          </el-form-item>
        </el-form>
      </el-main>
    </el-container>
  </div>
</template>

<script></script>

<style scoped></style>

<script>
export default {
  name: "pay",
  data() {
    return {
      qname: "",
      qdesc: "",
    };
  },
};
</script>

<style scoped></style>
