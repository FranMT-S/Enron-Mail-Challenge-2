<script setup lang="ts">
import { ref, type StyleValue } from 'vue'

interface Props{
  switchStyle?:StyleValue,
  switchClass?:any,
  spinStyle?:StyleValue,
  spingClass?:any,
}

const {switchStyle,spinStyle} = defineProps<Props>()

const emit = defineEmits<{
  (event: 'onChange'): void;
}>();

const model = defineModel()

const handleCheckboxChange = () => {
  emit('onChange')
}

</script>

<template>
  <label class="flex cursor-pointer select-none items-center">
    <div class="relative select-none">
      <input  style="" id="switch" v-model="model" type="checkbox" class="switch sr-only" @change="handleCheckboxChange" />
      <div :style="switchStyle" :class="switchClass" class="switch-container block h-[26px] w-14 rounded-full bg-[#E5E7EB]"></div>
      <div :style="spinStyle" :class="spingClass" class="spin dot absolute  top-1 h-[75%] w-[50%] rounded-full bg-[#fff] transition"></div>
    </div>
  </label>
</template>

<style>


#switch ~ .spin{
  transform: translateX(20%)  ;
}

#switch:checked ~ .spin{
  transform: translateX(80%)  ;
}
</style>
