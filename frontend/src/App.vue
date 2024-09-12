<script setup>
import {computed, ref} from "vue";

const toHome = () => {
  console.log("toHome")
}
const toCategory = () => {
  console.log("toCategory")
}
const toRank = () => {
  console.log("toRank")
}
const toSetting = () => {
  console.log("toSetting")
}
const value = ref(0)
const color = computed(() => {
  switch (value.value) {
    case 0:
      return 'blue-grey'
    case 1:
      return 'teal'
    case 2:
      return 'brown'
    case 3:
      return 'indigo'
    default:
      return 'blue-grey'
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
