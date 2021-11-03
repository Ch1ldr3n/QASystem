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
          <el-button
            @click="pay(props.row.id)"
            v-if="props.row.asked && props.row.state === 'created'"
            >去支付</el-button
          >
          <el-button
            disabled
            v-if="!props.row.asked && props.row.state === 'created'"
            >等待对方支付</el-button
          >
          <el-button
            disabled
            v-if="props.row.asked && props.row.state === 'paid'"
            >等待对方接单</el-button
          >
          <el-button
            type="success"
            @click="accept(props.row.id, true)"
            v-if="!props.row.asked && props.row.state === 'paid'"
            >接受提问</el-button
          >
          <el-button
            type="danger"
            @click="accept(props.row.id, false)"
            v-if="!props.row.asked && props.row.state === 'paid'"
            >拒绝提问</el-button
          >
          <el-button
            type="primary"
            @click="openChat(props.row)"
            v-if="['accepted', 'done'].includes(props.row.state)"
            >开始聊天</el-button
          >
          <el-button
            type="primary"
            @click="done(props.row.id)"
            v-if="props.row.state === 'accepted'"
            >完成问答</el-button
          >
          <el-button
            type="danger"
            @click="cancel(props.row.id)"
            v-if="
              props.row.asked &&
              !['accepted', 'done', 'canceled'].includes(props.row.state)
            "
            >取消提问</el-button
          >
        </template>
      </el-table-column>
      <el-table-column prop="title" label="问题" />
      <el-table-column prop="price" label="金额" sortable min-width="10%" />
      <el-table-column prop="ausername" label="回答者" min-width="10%" />
      <el-table-column prop="qusername" label="提问者" min-width="10%" />
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
            >{{ scope.row.asked ? "我提出的" : "我回答的" }}</el-tag
          >
        </template>
      </el-table-column>
    </el-table>
    <beautiful-chat
      style="z-index: 1000"
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
      participants: [],
      messageList: [], // the list of the messages to show, can be paginated and adjusted dynamically
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
    stateFormat(row) {
      switch (row.state) {
        case 'created':
          return '未支付';
        case 'paid':
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
      })
        .then((resp) => {
          if (!resp.ok) {
            throw new Error('确认失败!');
          }
        })
        .then(() => {
          this.$message({
            message: '确认成功',
            type: 'success',
          });
          // TODO: refresh page
        })
        .catch((error) => {
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
      })
        .then((resp) => {
          if (!resp.ok) {
            throw new Error('确认失败!');
          }
        })
        .then(() => {
          this.$message({
            message: '确认成功',
            type: 'success',
          });
          // TODO: refresh page
        })
        .catch((error) => {
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
      })
        .then((resp) => {
          if (!resp.ok) {
            throw new Error('取消失败!');
          }
        })
        .then(() => {
          this.$message({
            message: '取消成功',
            type: 'success',
          });
          // TODO: refresh page
        })
        .catch((error) => {
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
      tim
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
      const msg = tim.createTextMessage({
        to: `${this.chatid}`,
        conversationType: TIM.TYPES.CONV_GROUP,
        payload: {
          text: message.data.text,
        },
      });
      tim
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
      tim
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
        if (!resp.ok) {
          throw new Error('获取我的问题列表失败！');
        }
        return resp.json();
      })
      .then((data) => {
        this.tableData = [
          ...data.answeredlist.map((v) => Object.assign(v, { asked: false })),
          ...data.askedlist.map((v) => Object.assign(v, { asked: true })),
        ];
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
