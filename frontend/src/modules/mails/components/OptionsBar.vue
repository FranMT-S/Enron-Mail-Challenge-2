<script setup lang="ts">
import { useConfigStore } from '../store/useConfig';
import { storeToRefs } from 'pinia';

import { TypeVisualization } from '@/enums/TypeVisualization';
import CalendarIcon from '../../components/icons/CalendarIcon.vue';
import TableIcon from '@/modules/components/icons/TableIcon.vue';
import VerticalWindows from '@/modules/components/icons/VerticalWindows.vue';
import IconButton from './IconButton.vue';

const {isTableViewMode} = storeToRefs(useConfigStore())


</script>

<template>
<div class="h-16 flex items-center justify-between">
    <div class="flex items-center">
        <div class="flex items-center">
            <div class="flex items-center space-x-2">
                <IconButton
                  :is-active="isTableViewMode == TypeVisualization.Calendar"
                  @on-click="() => isTableViewMode = TypeVisualization.Calendar"
                 >
                  <CalendarIcon />
                </IconButton>
                <IconButton
                  :is-active="isTableViewMode == TypeVisualization.Table"
                  @on-click="() => isTableViewMode = TypeVisualization.Table"
                 >
                  <TableIcon />
                </IconButton>
                <IconButton
                  :is-active="isTableViewMode == TypeVisualization.Vertical"
                  @on-click="() => isTableViewMode = TypeVisualization.Vertical"
                 >
                  <VerticalWindows />
                </IconButton>
            </div>
        </div>
    </div>
  </div>
</template>
