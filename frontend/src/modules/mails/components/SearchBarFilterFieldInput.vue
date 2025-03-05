<script setup lang="ts">
import type { InputTypeHTMLAttribute } from 'vue';
import CalendarIconTwo from '@/modules/components/icons/CalendarIconTwo.vue';
const model = defineModel()

const {label,type,placeholder} = defineProps<{
  label?:string
  placeholder?:string
  type:InputTypeHTMLAttribute
}>()

const isDate = type == 'date' || type == 'datetime-local'

</script>

<template>

    <div class="flex flex-col grow shrink">
      <label v-if="label" class="leading-[1.6] text-base">{{ label }}</label>
      <input v-model="model" :type="type"  class="px-4 h-[35px] py-2 border focus:ring-gray-500 focus:border-[#c13be2] w-full sm:text-sm border-gray-300 rounded-md focus:outline-none text-gray-600" :placeholder="placeholder">

    </div>



</template>

