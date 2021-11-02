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
    <el-table-column prop="date" label="日期" sortable min-width="20%" column-key="date" />
    <el-table-column prop="title" label="问题"/>
    <el-table-column prop="name" label="昵称" min-width="10%" />
    <el-table-column prop="price" label="金额" sortable min-width="10%" />
    <el-table-column prop="state" label="状态" sortable min-width="10%" />

    <el-table-column
      prop="tag"
      label="Tag"
      min-width="10%"
      :filters="[
        { text: '我提的问题', value: 'ask' },
        { text: '别人问我的问题', value: 'que' },
      ]"
      :filter-method="filterTag"
      filter-placemeidnt="bottom-end"
    >
      <template #default="scope">
        <el-tag
          :type="scope.row.tag === 'ask' ? 'warning' : 'success'"
          disable-transitions
          >{{ scope.row.tag }}</el-tag
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
        { type: 'text', author: 'me', data: { text: 'Say yes!' } },
        { type: 'text', author: 'user1', data: { text: 'No.' } },
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
          id: row.answererid,
          name: 'username', // TODO
        },
      ];
      this.isChatOpen = true;
      this.newMessagesCount = 0;
    },
    sendMessage(text) {
      if (text.length > 0) {
        this.newMessagesCount = this.isChatOpen ? this.newMessagesCount : this.newMessagesCount + 1;
        this.onMessageWasSent({ author: 'support', type: 'text', data: { text } });
      }
    },
    onMessageWasSent(message) {
      // called when the user sends a message
      this.messageList = [...this.messageList, message];
    },
    closeChat() {
      this.isChatOpen = false;
    },
    handleScrollToTop() {
      // called when the user scrolls message list to top
      // leverage pagination for loading another page of messages
    },
  },
  created() {
    // TODO: filter question related to user
    fetch('/v1/question/list', {
      method: 'GET',
      // headers: {}
    })
      .then((resp) => {
        if (!resp.ok) { throw new Error('获取我的问题列表失败！'); }
        return resp.json();
      })
      .then((data) => {
        this.tableData = data.questionlist;
        console.log(data);
      })
      .catch((error) => {
        this.$message({
          message: error,
          type: 'error',
        });
      });
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
