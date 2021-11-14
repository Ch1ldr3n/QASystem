<template>
    <el-main>
    <el-row :gutter="20">
        <el-col :span="6">
        <el-card :body-style="{ padding: '10px' }">
            <img
            src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
            class="image"
            />
            <div style="padding: 14px">
            <span>Yummy hamburger</span>
            <div class="bottom">
                <time class="time">{{ currentDate }}</time>
                <el-button type="primary">Operating</el-button>
            </div>
            </div>
        </el-card>
        </el-col>
        <el-col :span="6">
        <el-card :body-style="{ padding: '10px' }">
            <img
            src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
            class="image"
            />
            <div style="padding: 14px">
            <span>Yummy hamburger</span>
            <div class="bottom">
                <time class="time">{{ currentDate }}</time>
                <el-button type="primary">Operating</el-button>
            </div>
            </div>
        </el-card>
        </el-col>
        <el-col :span="6">
        <el-card :body-style="{ padding: '10px' }">
            <img
            src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
            class="image"
            />
            <div style="padding: 14px">
            <span>Yummy hamburger</span>
            <div class="bottom">
                <time class="time">{{ currentDate }}</time>
                <el-button type="primary">Operating</el-button>
            </div>
            </div>
        </el-card>
        </el-col>
            <el-col :span="6">
        <el-card :body-style="{ padding: '10px' }">
            <img
            src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
            class="image"
            />
            <div style="padding: 14px">
            <span>Yummy hamburger</span>
            <div class="bottom">
                <time class="time">{{ currentDate }}</time>
                <el-button type="primary">Operating</el-button>
            </div>
            </div>
        </el-card>
        </el-col>
    </el-row>
    </el-main>
</template>

<script>
export default ({
  data() {
    return {
      tableData: [],
    };
  },
  created() {
    // TODO: ask backend to return more related information
    fetch('/v1/question/list', {
      method: 'GET',
      headers: {},
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('获取知识广场信息失败！');
        }
        return resp.json();
      })
      .then((data) => {
        this.tableData = data.questionlist;
        console.log(data);
      })
      .catch((error) => {
        this.$message({ message: error, type: 'error' });
      });
  },
});
</script>

<style>
.bottom {
  margin-top: 13px;
  line-height: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.image {
  width: 100%;
  display: block;
}
</style>
