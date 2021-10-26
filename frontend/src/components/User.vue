<template>
  <div>
    <div>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-card shadow="hover">
            <template #header>
              <div class="clearfix">
                <span>基础信息</span>
              </div>
            </template>
            <div class="info">
              <div class="info-image">
                <img src="../assets/logo.png" />
                <span class="info-edit">
                  <i class="el-icon-lx-camerafill"></i>
                </span>
              </div>
              <div class="info-name">{{ name }}</div>
              <div class="info-desc">{{ description }}</div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="12">
          <el-card shadow="hover">
            <template #header>
              <div class="clearfix">
                <span>账户编辑</span>
              </div>
            </template>
            <el-form
              label-width="90px"
              ref="ruleForm"
              :model="ruleForm"
              status-icon
              :rules="rules"
            >
              <el-form-item label="用户名: ">{{ name }}</el-form-item>
              <el-form-item label="旧密码: ">
                <el-input type="password" v-model="password.old"></el-input>
              </el-form-item>
              <el-form-item label="新密码: " prop="pass">
                <el-input type="password" v-model="password.new1"></el-input>
                <span v-if="this.password.valid1 === false" style="color: red;"
                  >请输入合法密码</span
                >
              </el-form-item>
              <el-form-item label="确认密码: ">
                <el-input type="password" v-model="password.new2"> </el-input>
              </el-form-item>

              <el-form-item>
                <el-button
                  type="primary"
                  @click="onSubmit"
                  :disabled="!password.valid1"
                  >保存</el-button
                >
              </el-form-item>
            </el-form>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: 'user',
  data() {
    return {
      password: {
        old: '', // 这里和后端绑定
        new1: '',
        valid1: true,
        new2: '',
      },
      name: '后端传进来', //这里和后端那个数据绑定
      description: '这个人很懒，什么都没留下',
    }
  },
  methods: {
    onSubmit() {
      console.log(this.password)
      //检查新密码前后输入是否一致
      const isSame = this.password.new1 === this.password.new2
      if (isSame) {
        //向后端发送用户信息修改请求
        const editSucc = true
        if (editSucc) {
          alert('修改成功！')
        } else {
          alert('修改失败！')
        }
        return
      } else {
        alert('新密码前后输入不一致')
      }
    },
  },
  watch: {
    'password.new1': {
      handler(newName) {
        if (newName === '') {
          this.password.valid1 = false
          return
        }
        this.password.valid1 = /^[-A-Za-z0-9_]{4,20}$/.test(newName)
      },
    },
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
