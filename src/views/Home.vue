<template>
  <div class="home">
    <h2>Kaomojy</h2>
    <div class="des-area border-dot">
      顔文字だけで、プログラムが書けます！！<br />面白いor難しいプログラムかけたら、
      <button class="twitter" v-on:click="tweetShare()">Tweet #kaomojy</button
      >してください！ <br />
      <a
        href="https://github.com/ryomak/brainfuck-web/issues/new"
        target="blank"
        >issueに</a
      >挙げていただけたら
      exampleに追加していこうと思います。（修正もお待ちしております）
      <br />仕組みは<a href="https://ryomak.info/2019/brainfuck-go/">ここ</a
      >を参照してください
    </div>
    <h3>対応表</h3>
    <table border="1" align="center" class="border-s">
      <tr>
        <td>😩:"&gt;"</td>
        <td>😄:"&lt;"</td>
        <td>😏:"+"</td>
        <td>😆:"-"</td>
        <td>😍:"["</td>
        <td>😴:"]"</td>
        <td>😡:"."</td>
      </tr>
    </table>

    <div class="execute-area">
      <div class="input-area">
        <h3>INPUT</h3>
        <div class="code">
          <textarea class="border-s" v-model="inputStr" col="5"></textarea>
        </div>
        <select class="border-s ma" style="font-size:1.3em;" v-model="selected">
          <option v-for="ex in exampleList" :key="ex" :value="ex">{{
            ex
          }}</option>
        </select>
        <button class="btn-rich ma" v-on:click="load()">load</button>
        <button
          class="btn-rich"
          style="display:block;margin:0 auto;"
          v-on:click="run()"
        >
          start
        </button>
      </div>
      <div class="result-area">
        <h3>OUTPUT</h3>
        <div class="border-s ma code output-area">
          {{ result }}
        </div>
      </div>
    </div>
  </div>
</template>
<script>
// @ is an alias to /src
import start from "@/assets/brainfuck.js";
import example from "@/assets/example.js";
import { getParam, addParam } from "@/assets/url.js";

export default {
  name: "home",
  components: {},
  data() {
    return {
      inputStr: getParam("input"),
      result: "",
      selected: ""
    };
  },
  computed: {
    exampleList: function() {
      return Object.keys(example);
    }
  },
  methods: {
    run: function() {
      this.result = start(this.inputStr);
    },
    load: function() {
      this.inputStr = example[this.selected];
    },
    tweetShare: function() {
      const message = `「${this.result.substring(
        0,
        50
      )}...」ができた%23kaomojy%23brainfxxk`;
      //const url = `https://twitter.com/intent/tweet?text=${message}&url=${addParam("input",this.inputStr)}`;
      const url = `https://twitter.com/intent/tweet?text=${message}&url=${window.location.origin}`;
      const option = "status=1,width=818,height=400,top=100,left=100";
      window.open(url, "twitter", option);
    }
  }
};
</script>
<style>
.des-area {
  width: 80vw;
  margin: 0 auto;
}
.output-area {
  padding: 2px;
  overflow: scroll;
}
.border-s {
  border: 4px solid #2c3e50;
}
.border-dot {
  border: 4px dotted #2c3e50;
}

.btn-rich {
  position: relative;
  display: inline-block;
  padding: 0.25em 0.5em;
  text-decoration: none;
  color: #fff;
  background: #2c3e50; /*色*/
  border: solid 1px #2c3e50; /*線色*/
  border-radius: 4px;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2);
  text-shadow: 0 1px 0 rgba(0, 0, 0, 0.2);
  font-size: 1.3em;
}

.btn-rich:active {
  /*押したとき*/
  border: solid 1px #2c3e50;
  box-shadow: none;
  text-shadow: none;
}

.ma {
  padding: 2px;
  margin: 1em auto;
}

.twitter {
  background-color: #55acee;
  border: 2px solid #55acee;
  border-radius: 20px;
  color: #fff;
  padding: 2px 15px;
  -webkit-transition: all 0.3s;
  transition: all 0.3s;
}
.twitter:hover {
  background-color: #fff;
  color: #55acee;
}

@media screen and (min-width: 601px) {
  .execute-area {
    width: 80%;
    display: flex;
    padding: 3em;
    margin: 0 auto;
    justify-content: center;
  }
  .input-area {
    margin: 0 auto;
    width: 40%;
  }
  .result-area {
    width: 40%;
    margin: 0 auto;
  }

  .code {
    margin: 0 auto;
    height: 8em;
  }

  textarea {
    width: 100%;
    height: 100%;
  }
}
@media screen and (max-width: 600px) {
  .execute-area {
    display: block;
    width: 100%;
  }
  .input-area {
    margin: 0 auto;
    width: 80%;
  }
  .result-area {
    margin: 0 auto;
    width: 80%;
  }
  .code {
    margin: 0 auto;
    width: 100%;
    height: 8em;
  }
  textarea {
    width: 100%;
    height: 100%;
  }
}
</style>
