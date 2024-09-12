<script setup>
import {computed, ref} from "vue";
import router from "./router/router";
console.log(router)
const toHome = () => {
  router.push({path: '/home'})
  console.log("toHome")
}
const toCategory = () => {
  router.push({path: '/category'})

  console.log("toCategory")
}
const toRank = () => {
  router.push({path: '/rank'})

  console.log("toRank")
}
const toSetting = () => {
  router.push({path: '/setting'})
  console.log("toSetting")
}
const value = ref(0)
const color = computed(() => {
  switch (value.value) {
    case 0:
      return 'blue-darken-2'
    case 1:
      return 'pink-darken-3'
    case 2:
      return 'green-darken-2'
    case 3:
      return 'deep-orange-darken-3'
    default:
      return 'orange-darken-1'
  }
})
</script>

<template>
  <v-app>
    <main>
      <RouterView>
      </RouterView>
    </main>
    <v-layout class="overflow-visible" style="height: 56px;">
      <v-bottom-navigation
          v-model="value"
          :bg-color="color"
          mode="shift"
          grow
      >
        <v-btn @click="toHome">
          <v-icon>mdi-television-play</v-icon>

          <span>首页</span>
        </v-btn>

        <v-btn @click="toCategory">
          <v-icon>mdi-music-note</v-icon>

          <span>热播榜</span>
        </v-btn>

        <v-btn @click="toRank">
          <v-icon>mdi-book</v-icon>

          <span>搞笑榜</span>
        </v-btn>

        <v-btn @click="toSetting">
          <v-icon>mdi-image</v-icon>

          <span>设置</span>
        </v-btn>
      </v-bottom-navigation>
    </v-layout>
  </v-app>
</template>

<style>

</style>
