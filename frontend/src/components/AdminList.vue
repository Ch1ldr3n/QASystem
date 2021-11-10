<template>
  <el-container style="margin-top: 20px;">
    <el-header>
      <el-form
        ref="form"
        :model="form"
        :rules="rules"
      >
        <el-form-item prop="newname">
          <el-input
            v-model="form.newname"
            placeholder="管理员昵称"
          >
            <template #append>
              <el-button @click="adding">
                添加
              </el-button>
            </template>
          </el-input>
        </el-form-item>
      </el-form>
    </el-header>
    <el-main>
      <el-table
        ref="filterTable"
        :data="adminlist"
        :default-sort="{ prop: 'username', order: 'descending' }"
        style="width: 100%"
      >
        <el-table-column
          prop="username"
          label="姓名"
        />
        <el-table-column
          prop="role"
          label="职务"
          :formatter="stateFormat"
        />
        <el-table-column
          prop="role"
          label="操作"
        >
          <template #default="props">
            <el-button
              v-if="props.row.role === 'none'"
              type="primary"
              @click="change(props.row.username, 'reviewer')"
            >
              赋予审核权限
            </el-button>
            <el-button
              v-if="props.row.role === 'reviewer'"
              @click="change(props.row.username, 'none')"
            >
              取消审核权限
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<script>
export default {
  name: 'AdminList',
  components: {
  },
  data() {
    return {
      adminlist: [],
      form: {
        newname: '',
      },
      rules: {
        newname: [
          {
            required: true,
            message: '请输入管理员用户名',
            trigger: 'blur',
          },
          {
            min: 4,
            message: '管理员用户名长度至少为4',
            trigger: 'blur',
          },
        ],
      },
    };
  },
  created() {
    this.refresh();
  },
  methods: {
    stateFormat(row) {
      switch (row.state) {
        case 'admin':
          return '站长';
        case 'reviewer':
          return '审核员';
        case 'none':
          return '管理员';
        default:
          return '';
      }
    },
    refresh() {
      fetch('/v1/admin/list', {
        method: 'GET',
        headers: {
          'content-type': 'application/json',
          authorization: window.localStorage.getItem('admintoken'),
        },
      })
        .then((resp) => {
          if (!resp.ok) {
            throw new Error('刷新管理员列表失败！');
          }
          return resp.json();
        })
        .then((data) => {
          this.adminlist = data.userlist;
        })
        .catch((error) => {
          this.$message({
            message: error,
            type: 'error',
          });
        });
    },
    change(username, role) {
      fetch('/v1/admin/change', {
        method: 'POST',
        headers: {
          Authorization: window.localStorage.getItem('admintoken'),
          'content-type': 'application/json',
        },
        body: JSON.stringify({
          username,
          role,
        }),
      }).then((resp) => {
        if (!resp.ok) {
          if (resp.status === 403) {
            throw new Error('无修改管理员身份权限!');
          } else {
            throw new Error('修改管理员身份时发生错误!');
          }
        }
      }).then(() => {
        this.$message({
          showClose: true,
          message: '修改成功',
          type: 'success',
        });
        this.refresh();
      }).catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
    },
    adding() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          fetch('/v1/admin/add', {
            method: 'POST',
            headers: {
              Authorization: window.localStorage.getItem('admintoken'),
              'content-type': 'application/json',
            },
            body: JSON.stringify({
              token: window.localStorage.getItem('admintoken'),
              username: this.form.newname,
            }),
          }).then((resp) => {
            if (!resp.ok) {
              if (resp.status === 403) {
                throw new Error('无添加管理员权限!');
              } else {
                throw new Error('无法重复添加同名管理员!');
              }
            }
            return resp.json();
          }).then((data) => {
            this.$message({
              showClose: true,
              message: `添加成功，用户名 ${this.form.newname}，初始密码为 ${data.password}`,
              type: 'success',
              duration: 0,
            });
            this.refresh();
          }).catch((error) => {
            this.$message({
              message: error,
              type: 'error',
            });
          });
        }
      });
    },
  },
};
</script>
