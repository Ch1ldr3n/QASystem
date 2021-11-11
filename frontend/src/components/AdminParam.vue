<template>
  <el-main>
    <el-row justify="center">
      <el-col :span="10">
        <el-card shadow="hover">
          <template #header>
            <div class="clearfix">
              <span>编辑系统参数</span>
            </div>
          </template>
          <el-form
            ref="form"
            label-width="140px"
            status-icon
            :model="model"
            :rules="rules"
          >
            <el-form-item
              label="问题价格区间"
              prop="price"
            >
              <el-row>
                <el-col :span="9">
                  <el-input
                    v-model="model.min_price"
                    type="number"
                  />
                </el-col>
                <el-col
                  :span="4"
                  align="middle"
                >
                  —
                </el-col>
                <el-col :span="11">
                  <el-input
                    v-model="model.max_price"
                    type="number"
                  >
                    <template #append>
                      元
                    </template>
                  </el-input>
                </el-col>
              </el-row>
            </el-form-item>
            <el-form-item
              label="接单等待时长"
              prop="take_order_waiting_time"
            >
              <el-input
                v-model="model.take_order_waiting_time"
                type="number"
              >
                <template #append>
                  秒
                </template>
              </el-input>
            </el-form-item>
            <el-form-item
              label="作答等待时长"
              prop="first_answer_waiting_time"
            >
              <el-input
                v-model="model.first_answer_waiting_time"
                type="number"
              >
                <template #append>
                  秒
                </template>
              </el-input>
            </el-form-item>
            <el-form-item
              label="最大问答次数"
              prop="max_qa_times"
            >
              <el-input
                v-model="model.max_qa_times"
                type="number"
              >
                <template #append>
                  次
                </template>
              </el-input>
            </el-form-item>
            <el-form-item
              label="最长服务时长"
              prop="max_service_time"
            >
              <el-input
                v-model="model.max_service_time"
                type="number"
              >
                <template #append>
                  秒
                </template>
              </el-input>
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
  name: 'AdminParam',
  data() {
    const checkPrice = (rules, value, callback) => {
      if (!this.model.min_price || !this.model.max_price) {
        callback(new Error('价格区间不能为空！'));
      }
      setTimeout(() => {
        if (this.model.min_price > this.model.max_price) {
          callback(new Error('最低价格不能超过最高价格！'));
        } else {
          callback();
        }
      }, 500);
    };
    return {
      model: {
        min_price: '',
        max_price: '',
        take_order_waiting_time: '',
        first_answer_waiting_time: '',
        max_qa_times: '',
        max_service_time: '',
      },
      rules: {
        price: [
          {
            required: true,
            validator: checkPrice,
            trigger: ['blur'],
          },
        ],
        take_order_waiting_time: [
          {
            required: true,
            trigger: ['blur'],
            message: '接单等待时长不能为空！',
          },
        ],
        first_answer_waiting_time: [
          {
            required: true,
            trigger: ['blur'],
            message: '首次作答等待时长不能为空！',
          },
        ],
        max_qa_times: [
          {
            required: true,
            trigger: ['blur'],
            message: '最大问答次数不能为空！',

          },
        ],
        max_service_time: [
          {
            required: true,
            trigger: ['blur'],
            message: '最长服务时长不能为空！',
          },
        ],
      },
    };
  },

  created() {
    fetch('/v1/admin/param', {
      method: 'GET',
      headers: {
        Authorization: window.localStorage.getItem('admintoken'),
      },
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('用户信息加载失败');
        }
        return resp.json();
      })
      .then((data) => {
        this.model.min_price = data.min_price;
        this.model.max_price = data.max_price;
        this.model.take_order_waiting_time = data.accept_deadline;
        this.model.first_answer_waiting_time = data.answer_deadline;
        this.model.max_qa_times = data.answer_limit;
        this.model.max_service_time = data.done_deadline;
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
      // admintoken
      this.$refs.form.validate((valid) => {
        if (valid) {
          fetch('/v1/admin/param', {
            method: 'POST',
            headers: {
              Authorization: window.localStorage.getItem('admintoken'),
              'content-type': 'application/json',
            },
            body: JSON.stringify({
              accept_deadline: parseInt(this.model.take_order_waiting_time, 10),
              answer_deadline: parseInt(this.model.first_answer_waiting_time, 10),
              answer_limit: parseInt(this.model.max_qa_times, 10),
              done_deadline: parseInt(this.model.max_service_time, 10),
              max_price: parseFloat(this.model.max_price),
              min_price: parseFloat(this.model.min_price),
              token: window.localStorage.getItem('admintoken'),
            }),
          })
            .then((resp) => {
              if (!resp.ok) {
                throw new Error('无修改系统参数权限!');
              }
              this.$message({
                message: '修改成功',
                type: 'success',
              });
              this.$forceUpdate();
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
