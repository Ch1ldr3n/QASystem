<template>
  <el-main>
    <el-container>
      <el-table
        ref="filterTable"
        :data="tableData.slice((currentPage-1)*pageSize,currentPage*pageSize)"
        :default-sort="{ prop: 'date', order: 'descending' }"
        style="width: 100%"
        stripe
      >
        <el-table-column type="expand">
          <template #default="props">
            <p>{{ props.row.content }}</p>
            <el-button @click="openChat(props.row)">
              查看详情
            </el-button>
          </template>
        </el-table-column>
        <el-table-column
          prop="title"
          label="问题"
        />
        <el-table-column
          prop="qusername"
          label="提问者"
          min-width="10%"
        />
        <el-table-column
          prop="ausername"
          label="回答者"
          min-width="10%"
        />
      </el-table>
    </el-container>
    <el-container>
      <el-pagination
        v-model:currentPage="currentPage"
        :page-size="10"
        layout="prev, pager, next, jumper"
        :total="total"
        style="margin:0 auto"
        :hide-on-single-page="true"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
      <beautiful-chat
        style="z-index: 1000"
        :participants="participants"
        :message-list="messageList"
        :is-open="isChatOpen"
        :close="closeChat"
        :on-message-was-sent="(msg) => console.log(msg)"
        :show-emoji="false"
        :open="() => {}"
        :show-file="false"
        :show-edition="false"
        :show-deletion="false"
        :show-launcher="false"
        :show-close-button="true"
        :colors="colors"
        :always-scroll-to-bottom="false"
        :disable-user-list-toggle="true"
        :message-styling="true"
        @scrollToTop="handleScrollToTop"
      />
    </el-container>
  </el-main>
</template>

<script>
import TIM from 'tim-js-sdk';

export default {
  data() {
    return {
      tableData: [],
      currentPage: 1,
      pageSize: 10,
      total: 1000,
      tim: {},
      participants: [],
      messageList: [], // the list of the messages to show, can be paginated and adjusted dynamically
      isChatOpen: false, // to determine whether the chat window should be open or closed
      showTypingIndicator: '', // when set to a value matching the participant.id it shows the typing indicator for the specific user
      colors: {
        header: {
          bg: '#4e8cff',
          text: '#ffffff',
        },
        launcher: {
          bg: '#4e8cff',
        },
        messageList: {
          bg: '#ffffff',
        },
        sentMessage: {
          bg: '#4e8cff',
          text: '#ffffff',
        },
        receivedMessage: {
          bg: '#eaeaea',
          text: '#222222',
        },
        userInput: {
          bg: '#f4f7f9',
          text: '#565867',
        },
      },
      nextReqMessageID: '',
      isCompleted: false,
      row: {},
      chatid: 0,
    };
  },
  async created() {
    const options = {
      SDKAppID: 1400586942,
    };
    this.tim = TIM.create(options);
    this.tim.setLogLevel(1);
    // this.tim.on(TIM.EVENT.MESSAGE_RECEIVED, this.onMessageReceived);
    await this.tim.logout().catch((e) => console.log(e));
    fetch('/v1/user/genpublicsig', {
      method: 'GET',
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error('获取imsdk签名失败');
        }
        return resp.json();
      })
      .then((data) => this.tim.login({ userID: data.userid, userSig: data.signature }))
      .then((resp) => {
        console.log(resp);
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
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
        this.total = data.num;
        console.log(data);
      })
      .catch((error) => {
        this.$message({ message: error, type: 'error' });
      });
  },
  methods: {
    openChat(row) {
      console.log(row);
      this.participants = [
        {
          id: 'other',
          name: row.ausername,
        },
      ];
      this.tim
        .getMessageList({ conversationID: `GROUP${row.id}`, count: 15 })
        .then((imResponse) => {
          this.messageList = imResponse.data.messageList.map((x) => ({
            type: 'text',
            author: x.from === `${row.answererid}` ? 'other' : 'me',
            data: { text: x.payload.text },
          }));
          this.chatid = row.id;
          this.row = row;
          console.log(imResponse.data.messageList);
          this.nextReqMessageID = imResponse.data.nextReqMessageID;
          this.isCompleted = imResponse.data.isCompleted;
        });
      this.isChatOpen = true;
      this.newMessagesCount = 0;
    },
    handleScrollToTop() {
      this.tim
        .getMessageList({
          conversationID: `GROUP${this.chatid}`,
          count: 15,
          nextReqMessageID: this.nextReqMessageID,
        })
        .then((imResponse) => {
          this.messageList = [
            ...imResponse.data.messageList.map((x) => ({
              type: 'text',
              author: x.from === `${this.row.answererid}` ? 'other' : 'me',
              data: { text: x.payload.text },
            })),
            ...this.messageList,
          ];
          this.nextReqMessageID = imResponse.data.nextReqMessageID;
          this.isCompleted = imResponse.data.isCompleted;
        });
    },
    closeChat() {
      this.isChatOpen = false;
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`);
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
    },
  },
};
</script>
