<template>
  <el-dialog
        style="text-align: center"
        title="登陆"
        :show-close=false
        width="30%">
    <el-form label-width="80px">
      <el-form-item label="用户名">
        <el-input placeholder="username" v-model="state.username" autocomplete="off"></el-input>
        <span v-if="state.username_valid===false" style="color: red">请设置合法用户名!</span>
      </el-form-item>
    </el-form>
	<el-form label-width="80px">
      <el-form-item label="密码">
        <el-input placeholder="password" v-model="state.password" autocomplete="off"></el-input>
        <span v-if="state.password_valid===false" style="color: red">密码至少4位仅限字母数字</span>
      </el-form-item>
    </el-form>
	<!--el-form label-width="80px">
      <el-form-item label="邮箱">
        <el-input placeholder="email" v-model="state.email" autocomplete="off"></el-input>
        <span v-if="state.email_valid===false" style="color: red">请输入正确的邮箱!</span>
      </el-form-item>
    </el-form-->
    <span class="dialog-footer">
      <el-button v-on:click="quit">取 消</el-button>
      <el-button type="primary" v-on:click="register"
                  :disabled="!state.valid"
                  >注 册</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
	name: "Register",
	props:{
		dialogVisible: {
			type: Boolean,
			default: () => true
		},
		
	},
	data(){
		return {
			state: {
				username: "",
				username_valid: true,
				password: "",
				password_valid: true,
				email:"",
				//email_valid: true,
				//valid: false,
			}
		}
	},
	methods: {
		register: function(){
			this.$parent.register_(this.state.username,this.state.password)
		},
		quit: function(){
			this.$parent.quit()
		},
	},
	watch: {
		"state.username": {
			handler(newName) {
				this.state.username_valid = /^[A-Za-z\u4e00-\u9fa5][-A-Za-z0-9\u4e00-\u9fa5_]*$/.test(newName),
				this.state.valid = this.state.username_valid&&this.state.password_valid;
				if(newName=="")
					this.state.username_valid = true;
				if(this.state.username==""||this.state.password==""/*||this.state.email==""*/)
					this.state.valid = false;
			}
		},
		"state.password": {
			handler(newName) {
				this.state.password_valid = /^[-A-Za-z0-9_]{4,20}$/.test(newName),
				this.state.valid = this.state.username_valid&&this.state.password_valid;
				if(newName=="")
					this.state.password_valid = true;
				if(this.state.username==""||this.state.password==""/*||this.state.email==""*/)
					this.state.valid = false;
			}
		},/*
		"state.email": {
			handler(newName) {
				this.state.email_valid = /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$/.test(newName),
				this.state.valid = this.state.username_valid&&this.state.password_valid;
				if(newName=="")
					this.state.password_valid = true;
				if(this.state.username==""||this.state.password==""||this.state.email=="")
					this.state.valid = false;
			}
		}*/
	}
}
</script>