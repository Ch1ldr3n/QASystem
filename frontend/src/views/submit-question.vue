<template>
  <div>
    <el-main>
      <el-container>
        <el-aside width="unset">
          <el-card class="box-card">
            <div class="card-header">
              <el-image
                :src="
                  'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png'
                "
              />
            </div>
            <div>
              <el-descriptions column="1">
                <el-descriptions-item label="用户名"
                  >测试回答者</el-descriptions-item
                >
                <!-- <el-descriptions-item label="专业方向">信息科学</el-descriptions-item> -->
                <el-descriptions-item label="专业方向"
                  ><el-tag size="small">信息科学</el-tag></el-descriptions-item
                >
                <el-descriptions-item label="价格">100</el-descriptions-item>
              </el-descriptions>
            </div>
          </el-card>
        </el-aside>
        <el-container style="margin-left: 50px; margin-top: 50px;">
          <el-form>
            <el-form-item label="问题名称">
              <el-input v-model="qname"></el-input>
            </el-form-item>
            <el-form-item label="问题描述">
              <el-input
                type="textarea"
                :autosize="{ minRows: 6, maxRows: 10 }"
                v-model="qdesc"
              ></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmit">提交问题</el-button>
              <el-button>取消提问</el-button>
            </el-form-item>
          </el-form>
        </el-container>
      </el-container>
    </el-main>
  </div>
</template>

<script>
export default {
  name: 'pay',
  data() {
    return {
      qname: '',
      qdesc: '',
    }
  },
  methods: {
    onSubmit() {
      this.$router.push({
        name: 'Pay',
      })
      fetch('/v1/question/summit', {
        method: 'POST',
        headers: {
          'content-type': 'application/json',
        },
        body: JSON.stringify({
          answererid: 0,
          content: this.qdesc,
          price: 0,
          questionerid: 0,
          title: this.qname,
        }),
      }).then((resp) => {
        if (!resp.ok) {
          throw new Error('修改失败!')
        }
        this.$message({
          message: '修改成功',
          type: 'success',
        })
        //跳转到支付页面
        this.$router.push({
          name: 'Pay',
        })
      })
    },
  },
}
</script>

<style scoped>
.el-aside {
  height: 100%;
}
.el-input {
  width: 1000px;
}
</style>
