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
        <span v-if="state.username_valid===false" style="color: red">请设置合法密码!</span>
      </el-form-item>
    </el-form>
    <span class="dialog-footer">
      <el-button v-on:click="quit">取 消</el-button>
      <el-button type="primary" v-on:click="pushList"
                  :disabled="state.username_valid===false"
                  >确 定</el-button>
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
				type: Object,
				default: () => {
					return {
						username: "",
						username_valid: false,
						password: "",
						password_valid: false,
					}
				}
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
	watch: { // 用于实时检测username是否合法
		"state.username": {
			handler(newName) {
				this.state.username_valid = /^[A-Za-z\u4e00-\u9fa5][-A-Za-z0-9\u4e00-\u9fa5_]*$/.test(newName)
			}
		},
		"state.password": {
			handler(newName) {
				this.state.password_valid = /^[A-Za-z\u4e00-\u9fa5][-A-Za-z0-9\u4e00-\u9fa5_]*$/.test(newName)
			}
		}
	}
}
</script>