<script setup lang="ts">
import type { InputTypeHTMLAttribute } from 'vue';
import CalendarIconTwo from '@/modules/components/icons/CalendarIconTwo.vue';
const model = defineModel()

const {label,type,} = defineProps<{
  label?:string
  type:InputTypeHTMLAttribute
}>()

const isDate = type == 'date' || type == 'datetime-local'

</script>

<template>

    <div class="flex flex-col grow shrink">
      <label v-if="label" class="leading-loose text-base">{{ label }}</label>
      <template v-if="isDate">
        <div class="relative focus-within:text-gray-600 text-gray-400">
          <input :type="type" v-model="model" class="px-4 mobile:pl-10 py-2 border focus:ring-gray-500 focus:border-gray-900 w-full sm:text-sm border-gray-300 rounded-md focus:outline-none text-gray-600" placeholder="25/02/2020">
          <div class="absolute left-3 top-2 max-mobile:hidden">
            <CalendarIconTwo/>
          </div>
        </div>
      </template>
      <template v-else>
        <div class="flex flex-col"  >
          <input v-model="model" :type="type"  class="px-4 py-2 border focus:ring-gray-500 focus:border-gray-900 w-full sm:text-sm border-gray-300 rounded-md focus:outline-none text-gray-600" placeholder="email@gmail.com">
        </div>
      </template>
    </div>



</template>

