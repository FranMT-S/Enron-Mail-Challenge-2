<script setup lang="ts">
import SearchBarFilterFieldInput from './SearchBarFilterFieldInput.vue';
import { ref } from 'vue';
import { useClickOutside } from '../composables/useClickOutside';

export interface IFilterFormData {
  from?:string,
  to?:string,
  subject?:string,
  startDate?:Date,
  endDate?:Date,
}


interface IEvents {
  (e: 'on-close',isActive:boolean): void
  (e: 'on-submit',data:IFilterFormData): void
}

interface Props {
  isActive:boolean
}

const {isActive} = defineProps<Props>()
const emit = defineEmits<IEvents>()
const formData = ref<IFilterFormData>({
  from:undefined,
  to:undefined,
  subject:undefined,
  startDate:undefined,
  endDate:undefined,
})

const handleClickOutside = (event: MouseEvent) => {
  if (formRef.value && !formRef.value.contains(event.target as Node)) {
    close()
  }
};

const sendData = () =>{
  emit('on-submit',formData.value)
  close();
}

const close = () => {
  emit('on-close',false)
}

const {formRef} = useClickOutside(handleClickOutside)

</script>

<template>
  <form ref="formRef"  @submit.prevent="sendData" v-if="isActive" class="min-h-fit w-full max-w-[600px] right-0  flex flex-col justify-center top-full absolute w-100 z-[300] ">
      <div class="relative  px-4 py-10 bg-white mx-8 md:mx-0 shadow rounded-lg sm:p-10 shadow-modal-1">
        <div class="max-w-md mx-auto">
          <div class="divide-y divide-gray-200">
            <div class="py-8 text-base leading-6 space-y-4 text-gray-700 sm:text-lg sm:leading-7">
              <SearchBarFilterFieldInput type="text" v-model="formData.from" label="From"/>
              <SearchBarFilterFieldInput type="text" v-model="formData.to" label="To"/>
              <SearchBarFilterFieldInput type="text" v-model="formData.subject" label="Subject"/>
              <div class="flex items-center space-x-4">
                <SearchBarFilterFieldInput type="date" v-model="formData.startDate" label="Start"/>
                <SearchBarFilterFieldInput type="date" v-model="formData.endDate" label="End"/>
              </div>
            </div>
            <div class="pt-4 flex items-center space-x-4">
                <button @click="close" class="flex justify-center items-center w-full text-gray-900 px-4 py-3 rounded-md focus:outline-none">
                  <svg class="w-6 h-6 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg> Cancel
                </button>
                <button type="submit" class="bg-blue-500 flex justify-center items-center w-full text-white px-4 py-3 rounded-md focus:outline-none">Create</button>
            </div>
          </div>
        </div>
      </div>
    </form>
</template>
