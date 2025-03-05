<script setup lang="ts">
import { ref } from "vue";
import { useClickOutside } from "../composables/useClickOutside";
import { useSearchStore } from "../store/useSearchStore";
import { type SortByItem, type SortByItemList, type SortType } from "@/models/Sort";
import SortBySelectItem from "./SortBySelectItem.vue";
import SortDescLetter from "@/modules/components/icons/SortDescLetter.vue";
import SortAscLetter from "@/modules/components/icons/SortAscLetter.vue";
import SortAsc from "@/modules/components/icons/SortAsc.vue";
import SortDesc from "@/modules/components/icons/SortDesc.vue";

interface SelectedElement{
  label:string,
  sortLabel:string,
}

const items = ref<SortByItemList[]>([
  {
    label:'From',
    field:'from',
    ascLabel:'Asc',
    descLabel:'Desc',
  },
  {
    label:'To',
    field:'to',
    ascLabel:'Asc',
    descLabel:'Desc',
  },
  {
    label:'Subject',
    field:'subject',
    ascLabel:'Asc',
    descLabel:'Desc',
  },
])

const {sortValue,setSortField} = useSearchStore()
const selected = ref<SelectedElement>({
  label:sortValue.label,
  sortLabel:sortValue.sortLabel
})
const isOpen = ref(false);

const onSelect = (value:SortByItem) => {
  selected.value = {label:value.label,sortLabel:value.sortLabel}
  setSortField(value)
  isOpen.value = false
}

const handleClickOutside = (event: MouseEvent) => {
  if (isOpen.value && formRef.value && !formRef.value.contains(event.target as Node)) {
    isOpen.value = false
  }
};

const {formRef} = useClickOutside(handleClickOutside)

</script>

<template>
  <div class="relative w-fit  " ref="formRef">
    <button
      @click="isOpen = !isOpen"
      :class="{ 'border-[#7530be]': isOpen }"
      class="w-full w-min-[120px] outline-none bg-white border border-gray-300 rounded-[5px] px-2 py-[1px] text-left flex justify-between items-center"
    >
      <span class="flex flex-row gap-[4px] divide-x-2">
        <h1 class=" min-w-[55px] w-auto text-center">{{ selected.label }}</h1>
        <h2 class="px-[4px] min-w-[60px] text-center"> {{ selected.sortLabel }}</h2>
      </span>
      <span :class="{ 'rotate-180 !text-[#792acf]': isOpen }" class="text-gray-400 hover:text-[#792acf] transition duration-100 ms-[5px]">▼</span>
    </button>
    <ul
      v-if="isOpen"
      class="flex flex-col items-end gap-[4px] py-[8px] px-[8px] absolute  bg-white border border-gray-300 rounded-lg mt-1 shadow-lg z-10 w-auto "
    >
      <SortBySelectItem
        ascLabel="Oldest"
        descLabel="Newest"
        field="date"
        label="Date"
        @onSelect="onSelect"
      >
        <template #imgAsc>
          <SortAsc class="h-[20px] invisible group-hover/text:visible"/>
        </template>
        <template #imgDesc>
          <SortDesc class="h-[24px] invisible group-hover/text:visible"/>
        </template>
      </SortBySelectItem>
      <SortBySelectItem
        v-for="i of items"
        :ascLabel="i.ascLabel"
        :descLabel="i.descLabel"
        :field="i.field"
        :label="i.label"
        @onSelect="onSelect"
      >
        <template #imgAsc>
          <SortDescLetter class="h-[20px] invisible group-hover/text:visible"/>
        </template>
        <template #imgDesc>
          <SortAscLetter class="h-[20px] invisible group-hover/text:visible"/>
        </template>
      </SortBySelectItem>
    </ul>
  </div>
</template>
