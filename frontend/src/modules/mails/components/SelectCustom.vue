<script setup lang="ts">
import { ref } from "vue";
import { useClickOutside } from "../composables/useClickOutside";

interface Props{
  items:string[] | number[]
}

interface Emits{
  (e:'onSelect',select:string|number):void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const options = ref(props.items);
const selectedOption = ref<string | number | null | undefined>(options.value[0] ?? "");
const isOpen = ref(false);

const selectOption = (option:number | string) => {
  selectedOption.value = option;
  isOpen.value = false;
  emit('onSelect',option)
};

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
      @click.prevent.stop="isOpen = !isOpen"
      :class="{ 'border-[#7530be]': isOpen }"
      class="w-full h-full outline-none bg-white hover:bg-gray-50 border border-gray-300 rounded-[5px] px-2 py-[1px] text-left flex justify-between items-center"
    >
      <p class="text-center w-full">{{ selectedOption }}</p>
      <span :class="{ 'rotate-180 !text-[#792acf]': isOpen }" class="text-gray-400 hover:text-[#792acf] transition duration-100 ms-[5px]">▼</span>
    </button>
    <ul v-if="isOpen" class="absolute w-full bg-white border border-gray-300 rounded-lg mt-1 shadow-lg z-10">
      <li
        v-for="option in options"
        :key="option"
        @click="selectOption(option)"
        class="px-4 py-2 hover:bg-gray-200 cursor-pointer"
      >
        {{ option }}
      </li>
    </ul>
  </div>
</template>
