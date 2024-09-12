<script setup>
import {onMounted,ref} from "vue"
import { List } from "../../wailsjs/go/main/App";

const topics = ref([]);
const getTopics = () => {
  List().then((result) => {
      topics.value = result;
  });
}
onMounted(()=>{
  getTopics();
})

const showCtx=(topic)=>{
console.log(topic)
}

</script>

<template>
  <header>
    <div class="p-2 flex justify-center items-center text-sky-400" style="background:linear-gradient(135deg,#ff75c3,#ffa647,#ffe83f,#9fff5b,#70e2ff,#cd93ff)">
      <img src="@/assets/images/douyin.png" alt="" class="w-8">
      <span class="font-bold text-2xl text-green-900">评论区座谈会</span>
    </div>
  </header>
  <main class="mt-2 overflow-auto">
    <v-container>
      <v-row no-gutters>
        <v-col cols="4">
          <v-sheet   class="text-white mx-auto">
            <v-card
                loading
            >
              <v-list>
                <v-list-subheader>榜单</v-list-subheader>
                <v-list-item
                    v-for="(item, i) in topics"
                    :key="i"
                    :value="item"
                    color="primary"
                    rounded="shaped"
                >
                  <template v-slot:prepend>
                    <v-icon icon="mdi-video-high-definition"></v-icon>
                  </template>
                  <v-list-item-title v-text="item.label"></v-list-item-title>
                </v-list-item>
              </v-list>
            </v-card>
          </v-sheet>
        </v-col>

        <v-col cols="8">
          <v-sheet class="pa-2 ma-2">
            <h4 class="text-h5 font-weight-bold mb-4">Your Privacy</h4>

            <p class="mb-8">
              This business determines the use of personal data collectied on our media properties and across the internet. We may collect data that you submit to us directly or data that we collect automatically including from cookies (such as device information or IP address).

              <br>
              <br>

              Please read our <a class="text-red-accent-2" href="#">Privacy Policy</a> to learn about our privacy practices or click "Your Preferences" to exercise control over your data.
            </p>

            <v-btn
                class="text-none text-black mb-4"
                color="red-accent-2"
                size="x-large"
                variant="flat"
                block
            >
              Accept
            </v-btn>

            <v-btn
                class="text-none text-black"
                color="red-accent-2"
                size="x-large"
                variant="outlined"
                block
            >
              Your Preferences
            </v-btn>
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
  </main>
  <footer>

  </footer>
</template>

<style scoped>

</style>
