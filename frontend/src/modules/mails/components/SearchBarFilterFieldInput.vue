<script setup lang="ts">
import type { InputTypeHTMLAttribute } from 'vue';
const model = defineModel()

defineProps<{
  label?:string
  placeholder?:string
  type:InputTypeHTMLAttribute
  errormsg?:string,
  isShowError?: boolean
  min?:number | string
  max?:number | string
  click?:() => void
  blur?:() => void
}>()

</script>

<template>
  <div class="flex flex-col grow shrink max-w-full">
    <label v-if="label" class="leading-[1.6] text-base">{{ label }}</label>
    <input :max="max" :min="min" @click="click" @blur="blur"  v-model="model" :type="type"  class="px-4 h-[35px] py-2 border focus:ring-gray-500 focus:border-[#c13be2] w-full sm:text-sm border-gray-300 rounded-md focus:outline-none text-gray-600" :placeholder="placeholder">
    <label v-if="isShowError" class="text-red-600 text-[15px] first-letter:capitalize "> {{ errormsg }}</label>
  </div>
</template>

