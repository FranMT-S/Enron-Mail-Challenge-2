<script setup lang="ts">
import SearchBarFilterFieldInput from './SearchBarFilterFieldInput.vue';
import { reactive, ref, watch, watchEffect } from 'vue';
import { useClickOutside } from '../composables/useClickOutside';
import Slider from './Slider.vue';
import CrossIcon from '@/modules/components/icons/CrossIcon.vue';
import type { IFilterFormData } from '@/models/FormFilter';





interface IEvents {
  (e: 'on-close',isActive:boolean): void
  (e: 'on-submit',data:IFilterFormData): void
}

interface Props {
  isActive:boolean
}

const {isActive} = defineProps<Props>()

const emit = defineEmits<IEvents>()
const formData = reactive<IFilterFormData>({
  from:null,
  to:null,
  subject:null,
  after:null,
  before:null,
  since:null,
  until:null,
})

const dateAfter = ref<Date|null>(null)
const dateBefore = ref<Date|null>(null)
const dateSince = ref<Date|null>(null)
const dateUntil = ref<Date|null>(null)


const isBetween = ref(false)
const isPickerOpen = ref(false)

const handleClickOutside = (event: MouseEvent) => {
  if (!isPickerOpen.value && formRef.value && !formRef.value.contains(event.target as Node)) {
    close()
  }
};

const sendData = () =>{
  formData.before = dateBefore.value
  formData.after = dateAfter.value
  formData.since = dateSince.value
  formData.until = dateUntil.value

  emit('on-submit',formData)
  close();
}

const close = () => {
  emit('on-close',false)
}

watch(isBetween, () =>{
  dateAfter.value = null
  dateBefore.value = null
  dateSince.value = null
  dateUntil.value = null
})

const {formRef} = useClickOutside(handleClickOutside)

</script>

<template>
  <form ref="formRef"  @submit.prevent="sendData" @keypress.enter.prevent="sendData" v-if="isActive" class="min-h-fit w-full max-w-[600px] right-0  flex flex-col justify-center top-full absolute w-100 z-[300] ">
      <div class="relative  px-4 py-6  max-h-[86vh] md:max-h-fit overflow-auto bg-white md:ms-8 ms-0 shadow rounded-lg shadow-modal-1">
        <div class="max-w-md mx-auto">
          <div class="divide-y divide-gray-200">
            <section class=" text-base leading-6 space-y-4 text-gray-700 sm:text-lg sm:leading-7">
              <SearchBarFilterFieldInput type="text" v-model="formData.from" label="From"/>
              <SearchBarFilterFieldInput type="text" v-model="formData.to" label="To"/>
              <SearchBarFilterFieldInput type="text" v-model="formData.subject" label="Subject"/>

              <div class="flex flex-col  gap-[4px] border-gray-200 border p-[10px] rounded-[5px] py-[12px]">
                <div class="flex flex-row row gap-[10px] mb-[8px] ">
                  <Slider v-model="isBetween"
                    :switch-style=" isBetween ? {'background-color':'#2563eb'} : undefined"
                  />
                  <span :class="{'font-semibold text-[#2049d5]': isBetween}">Between</span>
                </div>
                <div v-if="!isBetween" class="flex gap-[4px] flex-col items-stretch">
                  <SearchBarFilterFieldInput type="date" v-model="dateAfter" label="After"/>
                  <SearchBarFilterFieldInput type="date" v-model="dateBefore" label="Before"/>
                </div>
                <div v-else class="flex items-start flex-wrap gap-[4px] ">
                  <SearchBarFilterFieldInput type="date" v-model="dateSince" label="Since"/>
                  <SearchBarFilterFieldInput type="date" v-model="dateUntil" label="Until"/>
                </div>
              </div>
            </section>
            <footer class="pt-4 flex items-center gap-4 mb-[5px]">
                <button @click="close" class="flex justify-center items-center border opacity-70 hover:opacity-100 border-red-500 text-red-500 hover:text-white hover:bg-red-500 w-full  px-4 py-3 rounded-md focus:outline-none">
                  <CrossIcon/>Cancel
                </button>
                <button type="submit" class="bg-blue-500 flex hover:bg-blue-600  justify-center items-center w-full text-white px-4 py-3 rounded-md focus:outline-none">Create</button>
            </footer>
          </div>
        </div>
      </div>
    </form>
</template>
