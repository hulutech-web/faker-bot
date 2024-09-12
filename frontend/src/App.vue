<script setup>
import {computed, ref} from "vue";

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
    <header>
      <h1 class="text-3xl font-bold  text-red-800">
        评论区座谈会
      </h1>
    </header>
    <main>
      <router-view></router-view>
    </main>
      <v-layout class="overflow-visible" style="height: 56px;">
        <v-bottom-navigation
            v-model="value"
            :bg-color="color"
            mode="shift"
            grow
        >
          <v-btn>
            <v-icon>mdi-television-play</v-icon>

            <span>Video</span>
          </v-btn>

          <v-btn>
            <v-icon>mdi-music-note</v-icon>

            <span>Music</span>
          </v-btn>

          <v-btn>
            <v-icon>mdi-book</v-icon>

            <span>Book</span>
          </v-btn>

          <v-btn>
            <v-icon>mdi-image</v-icon>

            <span>Image</span>
          </v-btn>
        </v-bottom-navigation>
      </v-layout>
  </v-app>
</template>

<style>

</style>
