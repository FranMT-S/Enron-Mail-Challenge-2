<script setup lang="ts">
import type { MailSummary } from '@/models/Mails';
import { mergeProps, ref, watch } from 'vue';
import MailsSideBarItem from './MailsSideBarItem.vue';

interface Props{
  mails:MailSummary[]
  idMailSelected?:string,
  title?:string
}

interface Emit{
  (e:'OnMail', mail:MailSummary):void
}

const  props = withDefaults(defineProps<Props>(),{
  mails:() => []
})

const emit = defineEmits<Emit>()
const onClick = (mail:MailSummary) => emit('OnMail',mail)

const mails = ref(props.mails)


watch(() =>props.mails,(newValues) => {
  mails.value = newValues
})

</script>

<template>
  <section class="flex flex-col w-full bg-gray-50 h-full overflow-y-scroll max-w-[250px]">
    <slot></slot>
    <ul >
      <h1 class="text-gray-600 px-2">{{title}}</h1>
      <MailsSideBarItem @click="onClick(mail)" v-for="mail in mails" :key="mail.id" :mail="mail" :isSelected="mail.id == idMailSelected"/>
    </ul>
  </section>
</template>


<style scoped>

</style>
