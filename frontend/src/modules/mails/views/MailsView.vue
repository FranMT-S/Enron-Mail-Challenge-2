<script setup lang="ts">
import type { MailSummary } from '@/models/Mails';
import type { ResponseMailSummary } from '@/models/Response';
import { onMounted, ref } from 'vue';
import MailsTable from '../components/MailsTable.vue';
import NavBar from '../components/NavBar.vue';
import OptionsBar from '../components/OptionsBar.vue';
import { useMailService } from '../composables/useMailService';

  const emailsSummary = ref<MailSummary[]>([])
  const total = ref(0)
  const {getMails,getMailByID} = useMailService()

  const handlerGetMails = async () => {
    try {
      const mails:ResponseMailSummary = await getMails()
      total.value = mails.total
      emailsSummary.value = mails.emails
    } catch (err) {
      if (err instanceof DOMException && err.name === "AbortError") {

      } else {

      }
    }
  };

  const handlerGetMailByID = async (id:string) => {
    try {
      const mail =  await getMailByID(id)

    } catch (err) {
      if (err instanceof DOMException && err.name === "AbortError") {
        // stateRequest.value = "Petición cancelada manualmente";
        console.log("abort error");
      } else {
        console.log("error");
      }
    }
  };

  onMounted(handlerGetMails)

</script>

<template>

  <NavBar/>

<div class="w-full bg-white shadow-xl rounded-lg flex overflow-x-auto custom-scrollbar">
    <div class="flex-1 px-2" >
      <!-- <button class="border " @click="handlerGetMails">fetchMails</button>
      <button class="border " @click="() => handlerGetMailByID('2fxU108FkGc')">by id</button> -->
        <OptionsBar/>
        <MailsTable :mails="emailsSummary"/>

    </div>

  </div>
</template>

<style>


</style>
