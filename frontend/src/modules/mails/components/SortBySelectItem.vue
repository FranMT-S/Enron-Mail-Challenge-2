<script setup lang="ts">
import type { SortByItem, SortByItemList, SortType } from "@/models/Sort";

interface Emits{
  (e:'onSelect',value:SortByItem):void
}

const {ascLabel,descLabel,field,label} = defineProps<SortByItemList>()
const emit = defineEmits<Emits>()

const onSelect = (field:string,label:string,sort:SortType,sortLabel:string) => {
    emit('onSelect',{field,label,sort,sortLabel})
}

</script>

<template>
  <li
    class="cursor-pointer text-center w-full"
  >
    <div class="group flex flex-row justify-end   items-center gap-[2px] divide-x-2 border rounded hover:border-[#b574ff] ">
        <h1 class="px-2 group-hover:text-[#4f198d] w-full">{{label}}</h1>
        <div class="flex flex-col divide-y-2 item-center justify-center">
          <h2 @click="() => onSelect(field,label,'asc',ascLabel)" class="group/text flex items-center px-2 b hover:bg-[#4f198d] hover:text-[#fff]">
            <slot name="imgAsc"></slot>
            {{ascLabel}}
          </h2>
          <h2 @click="() => onSelect(field,label,'desc',descLabel)" class="group/text flex items-center px-2 hover:bg-[#4f198d] hover:text-[#fff] min-w-[93px]">
            <slot name="imgDesc"></slot>
            {{ descLabel }}
          </h2>
        </div>
    </div>
  </li>
</template>
