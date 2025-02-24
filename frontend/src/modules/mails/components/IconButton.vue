<script setup lang="ts">

  interface Events{
      (e: 'On-Click',event?:MouseEvent):void
  }

  interface Props{
    isActive?:boolean,
  }

  const emit = defineEmits<Events>()
  const onClick = (e:MouseEvent) => emit('On-Click',e)

  const {isActive} = defineProps<Props>()

</script>

<template>
  <button
      :class="[isActive ? 'text-purple-600' : 'text-gray-300 button']"
      @click="onClick" title="TableMode" class=" p-[3px] border border-gray-300 rounded-lg shadow  ">
      <slot></slot>
  </button>
</template>

