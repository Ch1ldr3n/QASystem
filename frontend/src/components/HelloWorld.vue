<template>
  <div class="hello">
    <c style="font-weight:bold;">{{ msg }}</c>
    <!-- <h1>{{ msg }}</h1> -->

    <br>

    <h3>海量题目</h3>
    <h3>专业认知</h3>
    <h3>安全可信</h3>

    <br>

    <h5>
        计算机、电子、金融、医药<br>
        面面俱到
    </h5>

    <h5>
      领域前沿学者<br>
      客观解答
    </h5>

    <h5>
      第三方支付平台<br>
      可信结算
    </h5>

    <h4>开发人员</h4>
     <p>
      队伍：紫荆我的最爱 队长：曹隽诚  前端：邱翊君 周满盈 车熙嘉   后端：曹隽诚 吴志立
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
 -->
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
  display: inline-block;
  text-align:center;
  /* border-style: solid;
  border-width:15px; */
  margin: 10px 200px 50px 160px;
  font-size: 18px;
}
h3 {
  display: inline-block;
  margin: 100px 200px 50px 200px;
  text-align:left;
  font-size: 25px;
}
h4 {
  margin: 100px 0 0;
}
ul {
  display: inline;
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
