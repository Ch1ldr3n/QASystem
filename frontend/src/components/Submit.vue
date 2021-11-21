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
              <el-descriptions :column="1">
                <el-descriptions-item label="用户名">
                  {{ aname }}
                </el-descriptions-item>
                <el-descriptions-item label="专业方向">
                  <el-tag size="small">
                    {{ aprof }}
                  </el-tag>
                </el-descriptions-item>
                <el-descriptions-item label="价格">
                  {{ price }}
                </el-descriptions-item>
              </el-descriptions>
            </div>
          </el-card>
        </el-aside>
        <el-container style="margin-left: 50px; margin-top: 50px;">
          <el-form>
            <el-form-item label="问题名称">
              <el-input v-model="qname" />
            </el-form-item>
            <el-form-item label="问题描述">
              <el-input
                v-model="qdesc"
                type="textarea"
                :autosize="{ minRows: 6, maxRows: 10 }"
              />
            </el-form-item>
            <el-form-item>
              <el-button
                type="primary"
                @click="onSubmit"
              >
                向他提问
              </el-button>
              <el-button
                @click="$router.go(-1)"
              >
                取消提问
              </el-button>
            </el-form-item>
          </el-form>
        </el-container>
      </el-container>
    </el-main>
  </div>
</template>

<script>
export default {
  name: 'Submit',
  data() {
    return {
      qname: '',
      qdesc: '',
      aname: '',
      aprof: '',
      price: '',
    };
  },
  created() {
    const param = new URLSearchParams({ id: parseInt(this.$route.query.id, 10) });
    fetch(`/v1/user/filter?${param}`, {
      method: 'GET',
      headers: { authorization: window.localStorage.getItem('token') },
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('获取用户信息失败');
        }
        return resp.json();
      })
      .then((data) => {
        this.aname = data.userlist[0].username;
        this.aprof = data.userlist[0].profession;
        this.price = data.userlist[0].price;
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
  },
  methods: {
    onSubmit() {
      console.log(JSON.stringify({
        answererid: parseInt(this.$route.query.id, 10),
        content: this.qdesc.split(' ').join('\n'),
        title: this.qname,
      }));
      fetch('/v1/question/submit', {
        method: 'POST',
        headers: {
          'content-type': 'application/json',
          authorization: window.localStorage.getItem('token'),
        },
        body: JSON.stringify({
          answererid: parseInt(this.$route.query.id, 10),
          content: this.qdesc.split(' ').join('\n'),
          title: this.qname,
        }),
      })
        .then((resp) => {
          if (!resp.ok) {
            throw new Error('提交失败!');
          }
          return resp.json();
        })
        .then((resp) => {
          this.$message({
            message: '提交成功',
            type: 'success',
          });
          this.$router.push({
            name: 'Pay',
            query: { id: resp.questionid },
          });
        })
        .catch((error) => {
          this.$message({
            message: error,
            type: 'error',
          });
        });
    },
  },
};
</script>

<style scoped>
.el-aside {
  height: 100%;
}
.el-input {
  width: 1000px;
}
</style>
