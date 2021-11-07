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
            ref="form"
            label-width="120px"
            :model="model"
            status-icon
            :rules="rules"
          >
            <el-form-item label="用户名: ">
              {{ model.name }}
            </el-form-item>
            <el-form-item
              v-if="model.answerer"
              label="身份:"
            >
              回答者
            </el-form-item>
            <el-form-item
              v-else
              label="身份:"
            >
              提问者
            </el-form-item>

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

            <el-form-item
              label="email"
              prop="email"
            >
              <el-input v-model="model.email" />
            </el-form-item>

            <el-form-item
              label="手机号码"
              prop="phone"
            >
              <el-input v-model="model.phone" />
            </el-form-item>

            <el-form-item label="钱包余额">
              {{ model.balance }}
            </el-form-item>

            <el-divider content-position="left">
              成为回答者
            </el-divider>
            <el-form-item label="是否成为回答者">
              <el-switch v-model="model.answerer" />
            </el-form-item>

            <el-form-item
              v-if="model.answerer"
              label="问题定价"
              prop="price"
            >
              <el-input
                v-model="model.price"
                type="number"
              />
            </el-form-item>

            <el-form-item
              v-if="model.answerer"
              label="专业领域方向"
              prop="profession"
            >
              <el-input v-model="model.profession" />
            </el-form-item>

            <el-form-item
              v-if="model.answerer"
              label="个人简介"
              prop="description"
            >
              <el-input
                v-model="model.description"
                type="textarea"
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
  name: 'User',
  data() {
    const checkPrice = (rule, value, callback) => {
      if (!value && this.model.answerer) {
        return callback(new Error('请完善回答者个人信息'));
      }
      if (value > 500) return callback(new Error('问题定价不能超过500'));
      return true;
    };
    return {
      model: {
        password1: '',
        password2: '',
        name: '',
        answerer: false,
        email: '',
        phone: '',
        price: 0,
        profession: '',
        balance: 0,
        description: '',
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
        email: [
          {
            pattern: /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/,
            message: '请输入正确的邮箱',
            trigger: 'blur',
          },
        ],
        phone: [
          {
            pattern: /^1[3|4|5|7|8][0-9]{9}$/,
            message: '请输入正确的电话号码',
            trigger: 'blur',
          },
        ],
        price: [
          {
            required: true,
            message: '请完善回答者的问题定价',
            trigger: 'blur',
          },
          {
            validator: checkPrice,
            trigger: 'blur',
          },
        ],
        profession: [
          {
            required: true,
            message: '请完善回答者的职业信息',
            trigger: 'blur',
          },
        ],
        description: [
          {
            required: true,
            message: '请完善回答者的个人简介',
            trigger: 'blur',
          },
        ],
      },
    };
  },

  created() {
    this.model.description = window.localStorage.getItem('description');
    fetch('/v1/user/info', {
      method: 'GET',
      headers: {
        Authorization: window.localStorage.getItem('token'),
      },
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('用户信息加载失败');
        }
        return resp.json();
      })
      .then((data) => {
        this.model.name = data.username;
        this.model.answerer = data.answerer;
        this.model.email = data.email;
        this.model.phone = data.phone;
        this.model.price = data.price;
        this.model.profession = data.profession;
        this.model.balance = data.balance;
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
  },
  methods: {
    quit() {
      this.$router.go(-1);
    },
    onSubmit() {
      // 向后端请求修改数据
      this.$refs.form.validate((valid) => {
        if (valid) {
          window.localStorage.setItem('description', this.model.description);
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
              price: parseFloat(this.model.price),
              profession: this.model.profession,
              token: window.localStorage.getItem('token'),
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
