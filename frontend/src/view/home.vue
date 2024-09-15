
<script setup>
import {onMounted, ref} from "vue"
import router from "../router/router.js";
import {Amusement, Drama, HotSearch, Topic,HotSentence, HotVideo, List, Trending} from "../../wailsjs/go/main/App";

const topics = ref([]);
const getTopics = () => {
  List().then((result) => {
    topics.value = result;
  });
}
onMounted(() => {
  getTopics();
})

const showCtx = (topic) => {
  switch (topic.ID) {
    case 1:
      HotVideo().then(_ => {
        console.log(_)
      })
      break;
    case 2:
      Topic().then(_ => {
        console.log(_)
      })
      break;

    case 3:
      Amusement().then(_ => {
        console.log(_)
      })
          break;
    case 4:
      Drama().then(_ => {
        console.log(_)
      })
      break;
    case 5:
      HotSearch().then(_ => {
        console.log(_)
      })
      break;
    case 6:
      Trending().then(_ => {
        console.log(_)
      })
      break;
    case 7:
      HotSentence("搞笑").then(_ => {
        console.log(_)
      })
      break;
  }
}
const model = ref(1);

const toAuthorize = ()=>{
  router.push({path:"/authorize"})
}
</script>

<template>
  <header>
    <div class="p-2 flex justify-center items-center text-sky-400"
         style="background:linear-gradient(135deg,#ff75c3,#ffa647,#ffe83f,#9fff5b,#70e2ff,#cd93ff)">
      <img src="@/assets/images/douyin.png" alt="" class="w-8">
      <span class="font-bold text-2xl text-green-900">评论区座谈会</span>
    </div>
  </header>
  <main class="overflow-auto">
    <v-card>
      <v-tabs
          bg-color="deep-purple-darken-4"
          center-active
          v-model="model"
      >
        <v-tab
               v-for="(item, i) in topics"
               :key="i"
               :value="item">
          {{item.label}}
        </v-tab>

      </v-tabs>
      <v-tabs-window v-model="model">
        <v-tabs-window-item
            v-for="(item, i) in topics"
            :key="i"
            :value="item"
        >
          <v-card>
            <v-sheet class="pa-2 ma-2">
              <h4 class="text-h5 font-weight-bold mb-4">用户授权</h4>

              <p class="mb-8">
               用户授权由2部分组成：1是官方open数据；2是需要用户授权后方可使用的数据

                <br>
                <br>

                用户授权<a class="text-red-accent-2" href="#">授权页面</a> 请扫描二维码进行授权操作.
              </p>

              <v-btn
                  @click="showCtx(item)"
                  class="text-none text-black mb-4"
                  color="red-accent-2"
                  size="x-large"
                  variant="flat"
                  block
              >
                获取数据
              </v-btn>
            </v-sheet>
          </v-card>
        </v-tabs-window-item>
      </v-tabs-window>
    </v-card>
  </main>
  <footer>
    <v-btn
        @click="toAuthorize"
        class="text-none text-black mb-4"
        color="red-accent-2"
        size="x-large"
        variant="flat"
        block
    >
      授权
    </v-btn>
  </footer>
</template>

<style scoped>

</style>
