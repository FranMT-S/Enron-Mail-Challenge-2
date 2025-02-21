<script setup lang="ts">
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
  <nav class="rounded-b-[10px] w-full flex py-[6px] px-[30px]  gap-[20px]  bg-[#f7ebff] text-white shadow-[1px_3px_20px_2px_#aa6af56e]">
      <LogoApp class="scale-125 "/>
      <SearchBar @on-search="onSearch" class="w-[50%] mx-auto"/>
  </nav>
</template>

