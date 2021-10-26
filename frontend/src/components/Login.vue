<template>
  <el-main>
    <el-row justify="center">
      <el-col :span="8">
        <el-card>
          <div class="card-header">
            <h2>登录</h2>
          </div>
          <div>
            <el-form :model="model" :rules="rules" ref="form">
              <el-form-item prop="username">
                <el-input v-model="model.username" placeholder="用户名"/>
              </el-form-item>
              <el-form-item prop="password">
                <el-input v-model="model.password" placeholder="密码" type="password"/>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="submit">登录</el-button>
              </el-form-item>
            </el-form>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </el-main>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      model: {
        username: "",
        password: "",
      },
      rules: {
        username: [
          {
            required: true,
            message: "请输入用户名",
            trigger: "blur"
          },
          {
            min: 5,
            message: "用户名长度至少为5",
            trigger: "blur"
          }
        ],
        password: [
          {
            required: true,
            message: "请输入密码",
            trigger: "blur"
          },
          {
            pattern: /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/,
            message: "密码长度至少为8，且至少有一个字母和数字",
            trigger: "blur"
          }
        ],
      },
    }
  },
  methods: {
    submit() {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          fetch("/v1/user/login", {
            method: "POST",
            body: JSON.stringify(this.model)
          })
          .then(resp => {
            if (!resp.ok) {
              throw new Error("用户名或密码错误")
            }
            return resp.json()
          })
          .then(data => {
            window.localStorage.setItem("token", data.token) 
            this.$message({
              message: "登录成功",
              type: "success",
            })
          })
          .catch(error => {
            this.$message({
              message: error,
              type: "error",
            })
          })
        } else {
          return false
        }
      })
    },
  },
};
</script>

