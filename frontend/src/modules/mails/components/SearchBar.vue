
<script setup lang="ts">
  import { ref } from 'vue';
  import FilterIcon from './icons/FilterIcon.vue';
  import SearchIcon from './icons/SearchIcon.vue';
  import SearchBarFilterModal, { type IFilterFormData } from './SearchBarFilterModal.vue';

  interface IEvents {
    (e: 'on-search',query:string): void
  }

  const isFilterActive = ref(false)
  const query = ref("")
  const emit = defineEmits<IEvents>()

  const EmitSearch = () =>{
    emit('on-search',query.value)
  }

  const onFilterClick = (event:MouseEvent) =>{
    isFilterActive.value = !isFilterActive.value
  }

  const onSubmit = (data:IFilterFormData) =>{
    let newQuery = "";
    let key: keyof typeof data;
    for (key in data) {
      if(!data[key])
        continue;

        newQuery += `${key}:(${data[key]}) `;
    }

    query.value = newQuery;
    EmitSearch()
  }

  const closeFilter = (isActive:boolean) =>{
    isFilterActive.value = isActive
  }


</script>

<template>
  <div  class="flex w-full rounded bg-white relative">
    <button @click="EmitSearch" type="submit" class="m-2 rounded-full flex justify-center items-center hover:bg-purple-100  p-2 text-purple-500 ">
       <SearchIcon/>
    </button>
    <input @keyup.enter="EmitSearch"  v-model="query" class=" w-full border-none bg-transparent px-4 py-1 text-gray-400 outline-none focus:outline-none " type="search" name="search" placeholder="Search..." />
    <button @click.stop="onFilterClick" type="submit" class="m-2 rounded-full flex justify-center items-center hover:bg-purple-100  p-2 text-purple-500 ">
       <FilterIcon/>
    </button>
    <SearchBarFilterModal @on-submit="onSubmit"  @on-close="closeFilter" :is-active="isFilterActive" v-if="isFilterActive"/>
  </div>
</template>
