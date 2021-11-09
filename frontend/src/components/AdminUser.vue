<template>
  <el-main>
    <el-row justify="center">
      <el-col :span="8">
        <el-card shadow="hover">
          <template #header>
            <div class="clearfix">
              <span>更改密码</span>
            </div>
          </template>
          <el-form
            ref="form"
            label-width="120px"
            :model="model"
            status-icon
            :rules="rules"
          >
            <el-form-item
              label="新密码: "
              prop="password1"
            >
              <el-input
                v-model="model.password1"
                type="password"
              />
            </el-form-item>
            <el-form-item
              label="确认密码: "
              prop="password2"
            >
              <el-input
                v-model="model.password2"
                type="password"
              />
            </el-form-item>

            <el-form-item>
              <el-button
                type="primary"
                @click="onSubmit"
              >
                保存
              </el-button>
              <el-button @click="quit">
                返回
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>
  </el-main>
</template>

<script>
export default {
  name: 'AdminUser',
  data() {
    return {
      model: {
        password1: '',
        password2: '',
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
            validator: (rule, value) => value === this.model.password1,
            message: '两次输入的密码不一致',
            trigger: 'blur',
          },
        ],
      },
    };
  },

  methods: {
    quit() {
      this.$router.go(-1);
    },
    onSubmit() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          fetch('/v1/admin/edit', {
            method: 'POST',
            headers: {
              Authorization: window.localStorage.getItem('admintoken'),
              'content-type': 'application/json',
            },
            body: JSON.stringify({
              password: this.model.password2,
            }),
          }).then((resp) => {
            if (!resp.ok) {
              throw new Error('修改失败!');
            }
            this.$message({
              message: '修改成功',
              type: 'success',
            });
            this.$forceUpdate();
          });
        }
      });
    },
  },
};
</script>
