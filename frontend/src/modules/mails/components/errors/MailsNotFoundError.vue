<script setup lang="ts">
import MagnifyingGlassImage from '../images/MagnifyingGlassImage.vue';

interface Props{
  onlyTitle?:boolean
}

const {onlyTitle} = defineProps<Props>()


</script>

<template>
  <div class="w-full flex flex-col justify-center items-center">
      <MagnifyingGlassImage class="max-w-[300px]"/>
      <div class="text-center">
        <h1 class="text-gray-800 text-lg font-semibold ">We didnt find anything</h1>
        <h2 class="text-gray-500 text-base font-semibold" v-if="!onlyTitle">Try a different words</h2>
      </div>
  </div>
</template>
