<script setup lang="ts">
interface Props{
  textColor?:string,
  backgroundColor?:string
}
defineProps<Props>()

</script>
<template>
  <p :style="{backgroundColor:backgroundColor,color:textColor}" class="text-light  text-[#f8eaff] inline-block bg-[#7759e3] py-[1px] px-[5px] rounded-lg w-max">
    <slot></slot>
  </p>
</template>

