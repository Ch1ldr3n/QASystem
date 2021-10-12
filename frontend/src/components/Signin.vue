<template>
  <el-dialog
        style="text-align: center"
        title="登陆"
        :show-close=false
        width="30%">
    <el-form label-width="80px">
      <el-form-item label="用户名">
        <el-input placeholder="username" v-model="state.username" autocomplete="off"></el-input>
        <span v-if="state.username_valid===false" style="color: red">请输入合法用户名!</span>
      </el-form-item>
    </el-form>
	<el-form label-width="80px">
      <el-form-item label="密码">
        <el-input placeholder="password" v-model="state.password" autocomplete="off"></el-input>
        <span v-if="state.password_valid===false" style="color: red">请输入合法密码!</span>
      </el-form-item>
    </el-form>
    <span class="dialog-footer">
      <el-button v-on:click="quit">取 消</el-button>
      <el-button type="primary" v-on:click="signin"
                  :disabled="!state.valid"
                  >登 陆</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
	name: "Signin",
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
				valid: false,
			}
		}
	},
	methods: {
		signin: function(){
			this.$parent.signin_()//this.state.username,this.state.password),
			this.state.username="",
			this.state.password=""
			},
		quit: function(){
			this.$parent.quit(),
			this.state.username="",
			this.state.password=""
		},
	},
	watch: {
		"state.username": {
			handler(newName) {
				this.state.username_valid = /^[A-Za-z\u4e00-\u9fa5][-A-Za-z0-9\u4e00-\u9fa5_]*$/.test(newName),
				this.state.valid = this.state.username_valid&&this.state.password_valid;
				if(newName=="")
					this.state.username_valid = true;
				if(this.state.username==""||this.state.password=="")
					this.state.valid = false;
			}
		},
		"state.password": {
			handler(newName) {
				this.state.password_valid = /^[A-Za-z][-A-Za-z0-9_]*$/.test(newName),
				this.state.valid = this.state.username_valid&&this.state.password_valid;
				if(newName=="")
					this.state.password_valid = true;
				if(this.state.username==""||this.state.password=="")
					this.state.valid = false;
			}
		}
	}
}
</script>