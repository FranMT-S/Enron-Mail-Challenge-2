<script setup lang="ts">
import { formatDate } from '@/helpers/formatDate';
import type { Mail, MailSummary } from '@/models/Mails';
import router from '@/routes';
import { NamesRouter } from '@/routes/routesNames';
import { computed, toRef } from 'vue';

interface MailTableProps{
  mails:MailSummary[]
}

const columns = ["From","To","Subject","Date"]

const props = defineProps<MailTableProps>();

const navigateToDetails = (id:string) =>{
  router.push({
    name:NamesRouter.MailDetail,
    params:{id}
  })
}

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
<div class="overflow-x-auto h-full mb-5 pr-[5px] rounded-t-[10px]">
        <table class="min-w-max w-full table-auto ">
            <thead class="sticky top-0 bg-[#6a2485] text-[#f3b2b2]">
                <tr class=" uppercase text-sm leading-normal">
                    <th class="py-3 px-6 text-center w-[200px]" v-for="column in columns" :key="column">
                        {{ column }}
                    </th>
                </tr>
            </thead>
            <tbody class="text-gray-600 text-sm font-light text-center">

              <tr @click="() => navigateToDetails(mail.id)" class="border-b border-gray-200 hover:bg-gray-100" v-for="mail in newMails" :key="mail.id" >
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
