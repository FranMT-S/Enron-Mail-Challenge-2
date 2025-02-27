<script setup lang="ts">
import { useConfigStore } from '../store/useConfig';
import { storeToRefs } from 'pinia';

import { TypeVisualization } from '@/enums/TypeVisualization';
import CalendarIcon from '../../components/icons/CalendarIcon.vue';
import TableIcon from '@/modules/components/icons/TableIcon.vue';
import VerticalWindows from '@/modules/components/icons/VerticalWindows.vue';
import IconButton from './IconButton.vue';

interface Emit{
  (e:'onchange',value:TypeVisualization): void
}

const {visualizationMode} = storeToRefs(useConfigStore())
const emit = defineEmits<Emit>()

const emitOnChange = (visualization:TypeVisualization) =>{
  visualizationMode.value = visualization
  emit('onchange',visualization)
}

</script>

<template>
  <div class="py-[10px] mobile-sm:py-[16px] flex items-center justify-end mobile-sm:justify-between">
    <div class="flex items-center">
        <div class="flex items-center">
            <div class="flex items-center space-x-2">
                <IconButton
                  :is-active="visualizationMode == TypeVisualization.Table"
                  @on-click="() => emitOnChange(TypeVisualization.Table)"
                 >
                  <TableIcon />
                </IconButton>
                <IconButton
                  :is-active="visualizationMode == TypeVisualization.Vertical"
                  @on-click="() => emitOnChange(TypeVisualization.Vertical)"
                 >
                  <VerticalWindows />
                </IconButton>
                <!-- <IconButton
                  :is-active="visualizationMode == TypeVisualization.Calendar"
                  @on-click="() => emitOnChange(TypeVisualization.Calendar)"
                 >
                  <CalendarIcon />
                </IconButton> -->
            </div>
        </div>
    </div>
  </div>
</template>
