<template>
  <el-container>
  <el-table
    ref="filterTable"
    :data="tableData"
    :default-sort="{ prop: 'date', order: 'descending' }"
    style="width: 100%"
  >
    <el-table-column type="expand">
      <template #default="props">
        <p>{{ props.row.content }}</p>
        <el-button type="primary" @click="openChat(props.row)">开始聊天</el-button>
      </template>
    </el-table-column>
    <el-table-column prop="title" label="问题"/>
    <el-table-column prop="price" label="金额" sortable min-width="10%" />
    <el-table-column prop="ausername" label="回答者" min-width="10%" />
    <el-table-column prop="qusername" label="提问者" min-width="10%" />
    <el-table-column prop="state" label="状态" sortable min-width="10%" />

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
          >{{ scope.row.asked ? '我提出的' : '我回答的' }}</el-tag
        >
      </template>
    </el-table-column>
  </el-table>
  <beautiful-chat
    style="z-index: 1000;"
    :participants="participants"
    :onMessageWasSent="onMessageWasSent"
    :messageList="messageList"
    :newMessagesCount="newMessagesCount"
    :isOpen="isChatOpen"
    :close="closeChat"
    :showEmoji="false"
    :open="() => {}"
    :showFile="false"
    :showEdition="false"
    :showDeletion="false"
    :showLauncher="false"
    :showCloseButton="true"
    :colors="colors"
    :alwaysScrollToBottom="false"
    :disableUserListToggle="true"
    @scrollToTop="handleScrollToTop"
    :messageStyling="true"
    />
  </el-container>
</template>

<script>
import TIM from 'tim-js-sdk';

const options = {
  SDKAppID: 1400586942,
};
const tim = TIM.create(options);
tim.setLogLevel(0);

export default {
  data() {
    return {
      participants: [
      ],
      messageList: [
      ], // the list of the messages to show, can be paginated and adjusted dynamically
      newMessagesCount: 0,
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
  methods: {
    filterTag(value, row) {
      return row.tag === value;
    },
    openChat(row) {
      console.log(row);
      this.participants = [
        {
          id: 'other',
          name: row.asked ? row.ausername : row.qusername,
        },
      ];
      tim.getMessageList({ conversationID: `GROUP${row.id}`, count: 15 }).then((imResponse) => {
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
      const msg = tim.createTextMessage({
        to: `${this.chatid}`,
        conversationType: TIM.TYPES.CONV_GROUP,
        payload: {
          text: message.data.text,
        },
      });
      tim.sendMessage(msg).then((resp) => {
        console.log(resp);
        this.messageList = [...this.messageList, message];
      }).catch((error) => {
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
        if (msg.conversationID === `GROUP${this.chatid}`) {
          this.messageList = [...this.messageList, { type: 'text', author: msg.flow === 'in' ? 'other' : 'me', data: { text: msg.payload.text } }];
        }
      });
    },
    handleScrollToTop() {
      tim.getMessageList({ conversationID: `GROUP${this.chatid}`, count: 15, nextReqMessageID: this.nextReqMessageID }).then((imResponse) => {
        this.messageList = [...imResponse.data.messageList.map((x) => ({
          type: 'text',
          author: x.flow === 'in' ? 'other' : 'me',
          data: { text: x.payload.text },
        })), ...this.messageList];
        this.nextReqMessageID = imResponse.data.nextReqMessageID;
        this.isCompleted = imResponse.data.isCompleted;
      });
    },
  },
  created() {
    // TODO: filter question related to user
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
        console.log(data);
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
    tim.on(TIM.EVENT.MESSAGE_RECEIVED, this.onMessageReceived);
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
      .then((data) => tim.login({ userID: data.userid, userSig: data.signature }))
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
};
</script>
