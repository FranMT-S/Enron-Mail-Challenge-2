<script setup lang="ts">
import { ref } from 'vue';
import FilterIcon from '../../components/icons/FilterIcon.vue';
import SearchIcon from '../../components/icons/SearchIcon.vue';
import SearchBarFilterModal from './SearchBarFilterModal.vue';
import type { IFilterFormData } from '@/models/FormFilter';
import { formatDateYYYYMMDD } from '@/helpers/formatDate';

interface IEvents {
  (e: 'on-search',query:string): void
}

interface Props{
  initialQuery?:string
}

const isFilterActive = ref(false)

const {initialQuery} = defineProps<Props>()
const query = ref(initialQuery ?? "")
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

    switch (key) {
      case 'after':
      case 'before':
      case 'since':
      case 'until':
        const date = data[key]
        if(!date)
          break
        const dateFormated = formatDateYYYYMMDD(date)
        newQuery += `${key}:(${dateFormated}) `
        break;

      default:
        newQuery += `${key}:(${data[key]}) `;
        break;
    }

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
    <input @keyup.enter="EmitSearch"  v-model="query" class=" w-full border-none bg-transparent px-4 py-1 text-gray-400 focus:text-gray-700 outline-none focus:outline-none " type="search" name="search" placeholder="Search..." />
    <button @click.stop="onFilterClick" type="submit" class="m-2 rounded-full flex justify-center items-center hover:bg-purple-100  p-2 text-purple-500 ">
       <FilterIcon/>
    </button>
    <SearchBarFilterModal @on-submit="onSubmit"  @on-close="closeFilter" :is-active="isFilterActive" v-if="isFilterActive"/>
  </div>
</template>
