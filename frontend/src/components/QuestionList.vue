<template>
  <el-main>
    <el-container>
      <el-table
        ref="filterTable"
        :data="tableData"
        :default-sort="{ prop: 'date', order: 'descending' }"
        style="width: 100%"
      >
        <el-table-column type="expand">
          <template #default="props">
            <p>{{ props.row.content.split(' ').join('\n') }}</p>
            <el-button
              v-if="props.row.asked && props.row.state === 'created'"
              @click="pay(props.row.id)"
            >
              去支付
            </el-button>
            <el-button
              v-if="!props.row.asked && props.row.state === 'created'"
              disabled
            >
              等待对方支付
            </el-button>
            <el-button
              v-if="props.row.state === 'paid'"
              disabled
            >
              审核中
            </el-button>
            <el-button
              v-if="props.row.asked && props.row.state === 'reviewed'"
              disabled
            >
              等待对方接单
            </el-button>
            <el-button
              v-if="!props.row.asked && props.row.state === 'reviewed'"
              type="success"
              @click="accept(props.row.id, true)"
            >
              接受提问
            </el-button>
            <el-button
              v-if="!props.row.asked && props.row.state === 'reviewed'"
              type="danger"
              @click="accept(props.row.id, false)"
            >
              拒绝提问
            </el-button>
            <el-button
              v-if="['accepted', 'done'].includes(props.row.state)"
              type="primary"
              @click="openChat(props.row)"
            >
              开始聊天
            </el-button>
            <el-button
              v-if="props.row.state === 'accepted'"
              type="primary"
              @click="done(props.row.id)"
            >
              完成问答
            </el-button>
            <el-button
              v-if="
                props.row.asked &&
                  !['accepted', 'done', 'canceled'].includes(props.row.state)
              "
              type="danger"
              @click="cancel(props.row.id)"
            >
              取消提问
            </el-button>
          </template>
        </el-table-column>
        <el-table-column
          prop="title"
          label="问题"
        />
        <el-table-column
          prop="price"
          label="金额"
          sortable
          min-width="10%"
        />
        <el-table-column
          prop="ausername"
          label="回答者"
          min-width="10%"
        />
        <el-table-column
          prop="qusername"
          label="提问者"
          min-width="10%"
        />
        <el-table-column
          prop="state"
          label="状态"
          :formatter="stateFormat"
          sortable
          min-width="10%"
        />

        <el-table-column
          prop="asked"
          label="类型"
          min-width="10%"
          :filters="[
            { text: '我提出的', value: true },
            { text: '我回答的', value: false },
          ]"
          :filter-method="filterTag"
          filter-placemeidnt="bottom-end"
        >
          <template #default="scope">
            <el-tag
              :type="scope.row.asked ? 'warning' : 'success'"
              disable-transitions
            >
              {{ scope.row.asked ? "我提出的" : "我回答的" }}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-container>
    <el-container>
      <beautiful-chat
        style="z-index: 1000"
        :participants="participants"
        :on-message-was-sent="onMessageWasSent"
        :message-list="messageList"
        :new-messages-count="newMessagesCount"
        :is-open="isChatOpen"
        :close="closeChat"
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
      tim: {},
      participants: [],
      messageList: [], // the list of the messages to show, can be paginated and adjusted dynamically
      newMessagesCount: 0,
      currentPage: 1,
      pageSize: 10,
      total: 1000,
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
      tableData: [],
      nextReqMessageID: '',
      isCompleted: false,
      chatid: 0,
    };
  },
  async created() {
    this.refresh();
    const options = {
      SDKAppID: 1400586942,
    };
    this.tim = TIM.create(options);
    this.tim.setLogLevel(1);
    this.tim.on(TIM.EVENT.MESSAGE_RECEIVED, this.onMessageReceived);
    await this.tim.logout().catch((e) => console.log(e));
    fetch('/v1/user/gensig', {
      method: 'GET',
      headers: {
        Authorization: window.localStorage.getItem('token'),
      },
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
  },
  methods: {
    stateFormat(row) {
      switch (row.state) {
        case 'created':
          return '未支付';
        case 'paid':
          return '等待审核';
        case 'reviewed':
          return '等待接单';
        case 'accepted':
          return '进行中';
        case 'done':
          return '已完成';
        case 'canceled':
          return '已终止';
        default:
          return '';
      }
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`);
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
    },
    filterTag(value, row) {
      return row.asked === value;
    },
    pay(id) {
      this.$router.push({
        name: 'Pay',
        query: { id },
      });
    },
    accept(id, choice) {
      fetch('/v1/question/accept', {
        method: 'POST',
        headers: {
          'content-type': 'application/json',
          authorization: window.localStorage.getItem('token'),
        },
        body: JSON.stringify({
          choice,
          questionid: id,
        }),
      }).then((resp) => {
        if (!resp.ok) {
          throw new Error('确认失败!');
        }
      }).then(() => {
        this.$message({
          message: '确认成功',
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
    done(id) {
      fetch('/v1/question/close', {
        method: 'POST',
        headers: {
          'content-type': 'application/json',
          authorization: window.localStorage.getItem('token'),
        },
        body: JSON.stringify({
          questionid: id,
        }),
      }).then((resp) => {
        if (!resp.ok) {
          throw new Error('确认失败!');
        }
      }).then(() => {
        this.$message({
          message: '确认成功',
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
    cancel(id) {
      fetch('/v1/question/cancel', {
        method: 'POST',
        headers: {
          'content-type': 'application/json',
          authorization: window.localStorage.getItem('token'),
        },
        body: JSON.stringify({
          questionid: id,
        }),
      }).then((resp) => {
        if (!resp.ok) {
          throw new Error('取消失败!');
        }
      }).then(() => {
        this.$message({
          message: '取消成功',
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
    openChat(row) {
      console.log(row);
      this.participants = [
        {
          id: 'other',
          name: row.asked ? row.ausername : row.qusername,
        },
      ];
      this.tim
        .getMessageList({ conversationID: `GROUP${row.id}`, count: 15 })
        .then((imResponse) => {
          this.messageList = imResponse.data.messageList.map((x) => ({
            type: 'text',
            author: x.flow === 'in' ? 'other' : 'me',
            data: { text: x.payload.text },
          }));
          this.chatid = row.id;
          console.log(imResponse.data.messageList);
          this.nextReqMessageID = imResponse.data.nextReqMessageID;
          this.isCompleted = imResponse.data.isCompleted;
        });
      this.isChatOpen = true;
      this.newMessagesCount = 0;
    },
    onMessageWasSent(message) {
      console.log(this.chatid);
      const msg = this.tim.createTextMessage({
        to: `${this.chatid}`,
        conversationType: TIM.TYPES.CONV_GROUP,
        payload: {
          text: message.data.text,
        },
      });
      this.tim
        .sendMessage(msg)
        .then((resp) => {
          console.log(resp);
          this.messageList = [...this.messageList, message];
        })
        .catch((error) => {
          this.$message({
            message: error,
            type: 'error',
          });
        });
    },
    closeChat() {
      this.isChatOpen = false;
    },
    onMessageReceived(event) {
      event.data.forEach((msg) => {
        if (!this.isChatOpen || msg.conversationID !== `GROUP${this.chatid}`) {
          const q = this.tableData.find((row) => `GROUP${row.id}` === msg.conversationID);
          this.$message({
            message: `你有新的来自问题${q.title}的消息`,
            type: 'success',
          });
        }
        if (msg.conversationID === `GROUP${this.chatid}`) {
          this.messageList = [
            ...this.messageList,
            {
              type: 'text',
              author: msg.flow === 'in' ? 'other' : 'me',
              data: { text: msg.payload.text },
            },
          ];
        }
      });
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
              author: x.flow === 'in' ? 'other' : 'me',
              data: { text: x.payload.text },
            })),
            ...this.messageList,
          ];
          this.nextReqMessageID = imResponse.data.nextReqMessageID;
          this.isCompleted = imResponse.data.isCompleted;
        });
    },
    refresh() {
      fetch('/v1/question/mine', {
        method: 'GET',
        headers: {
          authorization: window.localStorage.getItem('token'),
        },
      })
        .then((resp) => {
          if (!resp.ok) { throw new Error('获取我的问题列表失败！'); }
          return resp.json();
        })
        .then((data) => {
          this.tableData = [...(data.answeredlist.map((v) => Object.assign(v, { asked: false }))), ...(data.askedlist.map((v) => Object.assign(v, { asked: true })))];
          this.total = data.answerednum + data.askednum;
          console.log(data);
        })
        .catch((error) => {
          this.$message({
            message: error,
            type: 'error',
          });
        });
    },
  },
};
</script>
