


<script setup lang="ts">
import { formatDate } from '@/helpers/formatDate';
import type { Mail, MailSummary } from '@/models/Mails';
import { computed, toRef } from 'vue';

export interface MailTableProps{
  mails:MailSummary[]
}

const columns = ["From","To","Subject","Date"]
// const {mails} = withDefaults(defineProps<MailTableProps>(),{
//   mails:() => []
// })
const props = defineProps<{
  mails: MailSummary[];
}>();



const newMails = computed(() => {
  return props.mails.map(mail => {
    return {
      ...mail,
      date: formatDate(mail.date)
    }
  })
});

</script>

<style scoped>

</style>


<template>
<div class="overflow-x-auto">
        <table class="min-w-max w-full table-auto">
            <thead>
                <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
                    <th class="py-3 px-6 text-center w-[200px]" v-for="column in columns" :key="column">
                        {{ column }}
                    </th>
                </tr>
            </thead>
            <tbody class="text-gray-600 text-sm font-light text-center">
                <tr class="border-b border-gray-200 hover:bg-gray-100" v-for="mail in newMails" :key="mail.id" >
                    <td class="max-w-[100px]  py-3 px-6  text-center text-nowrap text-ellipsis overflow-hidden">
                        <div class="text-ellipsis overflow-hidden text-nowrap">
                            <span class="font-medium">{{ mail.from }}</span>
                        </div>
                    </td>
                    <td class="max-w-[100px]  py-3 px-6 text-center  text-nowrap text-ellipsis overflow-hidden">
                        <div class="text-ellipsis overflow-hidden text-nowrap ">
                            <span>{{ mail.to }}</span>
                        </div>
                    </td>
                    <td class="max-w-[100px]  py-3 px-6 text-center  text-nowrap text-ellipsis overflow-hidden">
                        <div class="text-ellipsis overflow-hidden text-nowrap ">
                          {{ mail.subject }}
                        </div>
                    </td>
                    <td class="max-w-[100px]  py-3 px-6 text-center text-nowrap text-ellipsis overflow-hidden">
                        <span class="py-1 px-3 rounded-full text-xs">
                          {{ mail.date }}
                        </span>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<style>


</style>
