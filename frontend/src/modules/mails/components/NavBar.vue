
<script setup lang="ts">
import { ref } from 'vue';
import SearchBar from './SearchBar.vue';
import LogoApp from './LogoApp.vue';
  interface IEvents {
    (e: 'onSearch',query:string): void
  }

  const emit = defineEmits<IEvents>()
  const onSearch = (query:string) =>{
    emit('onSearch',query)
  }


</script>

<template>
  <nav class="w-full flex py-[6px] px-[30px]  gap-[20px]  bg-[#5c308e] text-white ">
      <LogoApp class="scale-125 "/>
      <SearchBar @on-search="onSearch" class="w-[50%] mx-auto"/>
  </nav>
</template>
