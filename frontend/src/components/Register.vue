<template>
  <el-main>
    <el-row justify="center">
      <el-col :span="8">
        <el-card>
          <div class="card-header">
            <h2>注册</h2>
          </div>
          <div>
            <el-form
              ref="form"
              :model="model"
              :rules="rules"
            >
              <el-form-item prop="username">
                <el-input
                  v-model="model.username"
                  placeholder="用户名"
                />
              </el-form-item>
              <el-form-item prop="password">
                <el-input
                  v-model="model.password"
                  placeholder="密码"
                  type="password"
                />
              </el-form-item>
              <el-form-item prop="password_confirm">
                <el-input
                  v-model="model.password_confirm"
                  placeholder="确认密码"
                  type="password"
                />
              </el-form-item>
              <el-form-item>
                <el-button
                  type="primary"
                  @click="submit"
                >
                  注册
                </el-button>
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
  name: 'Register',
  data() {
    return {
      model: {
        username: '',
        password: '',
        password_confirm: '',
      },
      rules: {
        username: [
          {
            required: true,
            message: '请输入用户名',
            trigger: 'blur',
          },
          {
            min: 5,
            message: '用户名长度至少为5',
            trigger: 'blur',
          },
        ],
        password: [
          {
            required: true,
            message: '请输入密码',
            trigger: 'blur',
          },
          {
            pattern: /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/,
            message: '密码长度至少为8，且至少有一个字母和数字',
            trigger: 'blur',
          },
        ],
        password_confirm: [
          {
            required: true,
            // eslint-disable-next-line no-unused-vars
            validator: (rule, value) => value === this.model.password,
            message: '两次输入的密码不一致',
            trigger: 'blur',
          },
        ],
      },
    };
  },
  methods: {
    submit() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          fetch('/v1/user/register', {
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: JSON.stringify({
              username: this.model.username,
              password: this.model.password,
            }),
          })
            .then((resp) => {
              if (!resp.ok) {
                throw new Error('用户名已被注册');
              }
              return resp.json();
            })
            .then((data) => {
              window.localStorage.setItem('token', data.token);
              this.$message({
                message: '注册成功',
                type: 'success',
              });
              this.$router.push({
                name: 'Question',
              });
            })
            .catch((error) => {
              this.$message({
                message: error,
                type: 'error',
              });
            });
          return true;
        }
        return false;
      });
    },
  },
};
</script>
