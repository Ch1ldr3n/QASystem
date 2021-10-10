<template>
  <div class="hello">
    <c>{{ msg }}</c>
    <!-- <h1>{{ msg }}</h1> -->
    <!-- <p>
      For a guide and recipes on how to configure / customize this project,<br>
      check out the
      <a href="https://cli.vuejs.org" target="_blank" rel="noopener">vue-cli documentation</a>.
    </p> -->

    <h3>功能介绍</h3>
     <h5>
      提问者可以向指定回答者发起提问并支付，问题经过审核系统审核后，<br>
      将会被推送给回答者，回答者在限定的时间内进行作答。<br>
      问答服务完成后将会自动对订单进行结算，<br>
      回答者和平台会进行收入和佣金分配。
     </h5>

    <h4>开发人员</h4>
     <p>
      队伍：紫荆我的最爱<br>
      队长：曹隽诚<br>
      前端：车熙嘉 邱翊君 周满盈<br>
      后端：曹隽诚 吴志立
     </p>


    <!-- <h3>Installed CLI Plugins</h3>
    <ul>
      <li><a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-babel" target="_blank" rel="noopener">babel</a></li>
      <li><a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-eslint" target="_blank" rel="noopener">eslint</a></li>
    </ul>
    <h3>Essential Links</h3>
    <ul>
      <li><a href="https://vuejs.org" target="_blank" rel="noopener">Core Docs</a></li>
      <li><a href="https://forum.vuejs.org" target="_blank" rel="noopener">Forum</a></li>
      <li><a href="https://chat.vuejs.org" target="_blank" rel="noopener">Community Chat</a></li>
      <li><a href="https://twitter.com/vuejs" target="_blank" rel="noopener">Twitter</a></li>
      <li><a href="https://news.vuejs.org" target="_blank" rel="noopener">News</a></li>
    </ul>
    <h3>Ecosystem</h3>
    <ul>
      <li><a href="https://router.vuejs.org" target="_blank" rel="noopener">vue-router</a></li>
      <li><a href="https://vuex.vuejs.org" target="_blank" rel="noopener">vuex</a></li>
      <li><a href="https://github.com/vuejs/vue-devtools#vue-devtools" target="_blank" rel="noopener">vue-devtools</a></li>
      <li><a href="https://vue-loader.vuejs.org" target="_blank" rel="noopener">vue-loader</a></li>
      <li><a href="https://github.com/vuejs/awesome-vue" target="_blank" rel="noopener">awesome-vue</a></li>
    </ul> -->
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
c {
  margin: 20px 0 0;
  font-size: 50px;
}
h5 {
  text-align:center;
  /* border-style: solid;
  border-width:15px; */
  margin: 30px 1250px 100px 250px;
  background-color: pink;
  font-size: 18px;
}
h3 {
  margin: 50px 0 50px 400px;
  text-align:left;
  font-size: 25px;
}
h4 {
  margin: 20px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}

</style>
