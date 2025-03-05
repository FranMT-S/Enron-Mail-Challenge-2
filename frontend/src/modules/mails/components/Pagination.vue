<script setup lang="ts">
import { ref, toRef, toRefs, watch } from 'vue';
import { usePagination } from '../composables/usePagination';
import LeftArrowIcon from '../../components/icons/LeftArrowIcon.vue';
import SelectCustom from './SelectCustom.vue';

interface Props {
  page: number;
  itemPerPage: number;
  totalElement: number;
  itemPerPageList: number[]
}

interface IEvents {
  (e: 'onChangePage',number:number): void
  (e: 'onItemPerListChange',number:number): void
}

const ButtonTextColor = 'text-[#9e00c6ab] hover:text-[#9d00c6] hover:bg-[#fcf8fb]'

const props = defineProps<Props>();
const emit = defineEmits<IEvents>()
const {page,itemPerPage,totalElement} = toRefs(props)

const {
  currentPage,
  currentFrom,
  currentTo,
  nextPage,
  previousPage,
  totalPages,
  isFirstPage,
  isLastPage
} = usePagination(page,itemPerPage,totalElement)

const isNumberRegex = /\d+/

const validIsNumber = (e:KeyboardEvent) =>{
  if(!isNumberRegex.test(e.key)){
    e.preventDefault()
  }
}

const validatePage = (e:Event) =>{
  const input = e.target as HTMLInputElement
  let newValue = Number(input.value);
  if(newValue < 0){
    input.value = '1'
    currentPage.value = 1
    return
  }

  if (newValue >= totalPages.value) {
    currentPage.value = totalPages.value
    input.value = totalPages.value.toString()
    return
  }

  currentPage.value = newValue
}

watch(currentPage,() => {
  emit('onChangePage',currentPage.value)
})

watch(page,(newValue) => {
  currentPage.value = newValue
})

const onChangeItemsPerList = (newValue:number | string) =>{
  currentPage.value = 1
  emit('onItemPerListChange',Number(newValue))
}

</script>

<template>

    <div class="flex items-end justify-center flex-col mb-[4px] mobile-sm:mb-[0px] max-mobile-sm:gap-y-[4px] gap-y-[2px] ">
        <section class="text-sm text-gray-500 flex flex-row  gap-[5px] row-gap-[30px]">
          <div class="flex  justify-between gap-[2px] ">
            <div  class="flex flex-row gap-[10px] items-end">
              <input  @keypress="validIsNumber" @input="validatePage"  type="number"  min="1" :max="totalPages" :value="currentPage"
              class=" hide-arrow outline-none border-b-[1px] border-[#bababa] text-center focus:border-[#6f28a8]"
              />
              <span>of</span>
              <span>{{ totalPages }}</span>
            </div>
          </div>
          <SelectCustom class="min-w-[65px]" :items="itemPerPageList" @OnSelect="onChangeItemsPerList"/>
        </section>
        <section class="flex items-center gap-[5px] ">
          <span class="text-sm text-gray-500">
            {{totalElement > 0 ? currentFrom : 0}}-{{  totalElement > 0 ? currentTo : 0 }} of {{totalElement}}
          </span>
          <div class="flex items-center gap-[2px] min-w-[65px]">
            <button @click="previousPage" class="px-[4px] py-[4px] rounded-lg" :disabled="isFirstPage"
              :class="[isFirstPage ? 'text-gray-400' : ButtonTextColor]"
            >
              <LeftArrowIcon class="rotate-180"/>
            </button>
            <button @click="nextPage" class="px-[4px] py-[4px] rounded-lg " title="Older" :disabled="isLastPage"
             :class="[isLastPage ? 'text-gray-400' : ButtonTextColor]"
            >
               <LeftArrowIcon/>
            </button>
          </div>
        </section>

    </div>
</template>
