<script setup lang="ts">
import { defineEmits } from 'vue';
import ChangeDesign from './icons/ChangeDesignIcon.vue';

const emit = defineEmits(['get-emails','get-emails-by-id']);

</script>

<template>
<div class="h-16 flex items-center justify-between">
    <div class="flex items-center">
        <div class="flex items-center">
            <div class="flex items-center space-x-2">
                <button title="Mark As Read" class="text-gray-700 p-[3px] border border-gray-300 rounded-lg shadow hover:bg-gray-200 ">
                  <ChangeDesign/>
                </button>
            </div>
        </div>
    </div>
  </div>
</template>
