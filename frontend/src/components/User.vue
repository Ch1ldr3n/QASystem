<template>
  <el-main>
    <el-row justify="center">
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>
            <div class="clearfix">
              <span>账户编辑</span>
            </div>
          </template>
          <el-form
            label-width="120px"
            ref="form"
            :model="model"
            status-icon
            :rules="rules"
          >
            <el-form-item label="用户名: ">{{ model.name }}</el-form-item>
            <el-form-item label="身份:" v-if="model.answerer"
              >回答者</el-form-item
            >
            <el-form-item label="身份:" v-else>提问者</el-form-item>
            <el-form-item label="新密码: " prop="password1">
              <el-input type="password" v-model="model.password1"></el-input>
            </el-form-item>
            <el-form-item label="确认密码: " prop="password2">
              <el-input type="password" v-model="model.password2"> </el-input>
            </el-form-item>

            <el-form-item label="是否成为回答者">
              <el-switch v-model="model.answerer"></el-switch>
            </el-form-item>
            <el-form-item label="email">
              <el-input v-model="model.email"> </el-input>
            </el-form-item>

            <el-form-item label="手机号码">
              <el-input v-model="model.phone"> </el-input>
            </el-form-item>

            <el-form-item label="账户余额">
              {{ model.price }}
            </el-form-item>

            <el-form-item label="职业">
              <el-input v-model="model.profession"> </el-input>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" @click="onSubmit">保存</el-button>
              <el-button>返回</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>
  </el-main>
</template>

<script>
export default {
  name: 'User',
  data() {
    return {
      model: {
        password1: '',
        password2: '',
        name: '', //这里和后端那个数据绑定
        answerer: false,
        email: '',
        phone: '',
        price: 0,
        profession: '',
      },
      rules: {
        password1: [
          {
            message: '请输入密码',
            trigger: 'blur',
          },
          {
            pattern: /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/,
            message: '密码长度至少为8，且至少有一个字母和数字',
            trigger: 'blur',
          },
        ],
        password2: [
          {
            validator: (rule, value) => value == this.model.password1,
            message: '两次输入的密码不一致',
            trigger: 'blur',
          },
        ],
      },
    }
  },
  methods: {
    onSubmit() {
      console.log('dian')
      //向后端请求修改数据
      this.$refs['form'].validate((valid) => {
        if (valid) {
          fetch('/v1/user/edit', {
            method: 'POST',
            headers: {
              Authorization: window.localStorage.getItem('token'),
              'content-type': 'application/json',
            },
            body: JSON.stringify({
              answerer: this.model.answerer,
              email: this.model.email,
              phone: this.model.phone,
              price: this.model.price,
              profession: this.model.profession,
              token: window.localStorage.getItem('token'),
            }),
          }).then((resp) => {
            if (!resp.ok) {
              throw new Error('修改失败!')
            }
            this.$message({
              message: '修改成功',
              type: 'success',
            })
            location.reload()
          })
        }
      })
    },
  },

  created() {
    fetch('/v1/user/info', {
      method: 'GET',
      headers: {
        Authorization: window.localStorage.getItem('token'),
      },
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('用户信息加载失败')
        }
        return resp.json()
      })
      .then((data) => {
        console.log(data)
        this.model.name = data.username
        this.model.answerer = data.answerer
        this.model.email = data.email
        this.model.phone = data.phone
        this.model.price = data.price
        this.model.profession = data.profession
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        })
      })
  },
}
</script>

<style scoped>
.info {
  text-align: center;
  padding: 35px 0;
}
.info-image {
  position: relative;
  margin: auto;
  width: 100px;
  height: 100px;
  background: #f8f8f8;
  border: 1px solid #eee;
  border-radius: 50px;
  overflow: hidden;
}
.info-image img {
  width: 100%;
  height: 100%;
}
.info-edit {
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  opacity: 0;
  transition: opacity 0.3s ease;
}
.info-edit i {
  color: #eee;
  font-size: 25px;
}
.info-image:hover .info-edit {
  opacity: 1;
}
.info-name {
  margin: 15px 0 10px;
  font-size: 24px;
  font-weight: 500;
  color: #262626;
}
.crop-demo-btn {
  position: relative;
}
.crop-input {
  position: absolute;
  width: 100px;
  height: 40px;
  left: 0;
  top: 0;
  opacity: 0;
  cursor: pointer;
}
</style>
