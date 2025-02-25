<script setup lang="ts">
import { formatDate } from '@/helpers/formatDate';
import { formatXField } from '@/helpers/mails';
import type { Mail } from '@/models/Mails';
import { ref, watchEffect } from 'vue';
import Chip from './Chip.vue';

interface Props {
  mail:Mail
}

const props = defineProps<Props>()
const xFrom = ref("")
const xToList = ref<string[]>([])
const dateFormated = ref("")

watchEffect(() =>{
  dateFormated.value = formatDate(props.mail.date)
  if(props.mail.x_from.length > 0){
    xFrom.value = props.mail.x_from.split("<")[0].trim()
    xToList.value =  formatXField(props.mail.x_to)
  }
})

</script>

<template>
  <header class="border px-[8px] rounded-[10px] flex flex-col gap-[10px] text-left items-start py-[20px] mb-[10px]  ">
    <div class="flex flex-col justify-between  w-full">
      <h3 class="font-semibold  text-light text-[#939393]">{{dateFormated}}</h3>
      <div class="flex flex-col gap-[3px]">
        <h3 class="font-semibold text-lg text-ellipsis text-nowrap overflow-hidden">{{xFrom}} <{{mail.from}}></h3>
        <div class="flex flex-row gap-[4px] flex-wrap text-gray-700">
          To: <Chip v-for="xTo in xToList" :value="xTo"/>
        </div>
      </div>
    </div>
  </header>
</template>



