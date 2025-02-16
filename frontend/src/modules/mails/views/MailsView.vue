<script setup lang="ts">
  import type { MailSummary } from '@/models/Mails';
  import type { ResponseMailSummary } from '@/models/Response';
  import { onBeforeMount, onMounted, ref } from 'vue';
  import MailsTable from '../components/MailsTable.vue';
  import NavBar from '../components/NavBar.vue';
  import OptionsBar from '../components/OptionsBar.vue';
  import Pagination from '../components/Pagination.vue';
  import { useMailService } from '../composables/useMailService';





  const emailsSummary = ref<MailSummary[]>([])
  const total = ref(1)
  const query = ref("este es un ejemplo")
  const page = ref(1)
  const totalPerPage = ref(30)
  const {getMails} = useMailService()

  const handlerGetMails = async (query:string = "", page:number = 1, size:number=  30) => {
    try {

      const mails:ResponseMailSummary = await getMails(query,page,size)
      total.value = mails.total
      emailsSummary.value = mails.emails

    } catch (err) {
      if (err instanceof DOMException && err.name === "AbortError") {

      } else {
      }
    }
  };

  const onChangePage = (page:number) =>{
    handlerGetMails(query.value,page,totalPerPage.value)
  }

  const onChangeQuery = (query:string) =>{
    handlerGetMails(query,1,totalPerPage.value)
  }

  onMounted(async () => await handlerGetMails(query.value,1,totalPerPage.value))

</script>

<template>
  <NavBar @on-search="onChangeQuery"/>

  <div class="w-full bg-white shadow-xl rounded-lg flex overflow-x-auto custom-scrollbar">
    <div class="flex-1 px-2 grid grid-rows-[64px_1fr]" >
        <div class="flex justify-between">
          <OptionsBar/>
          <Pagination @on-change-page="onChangePage " :page="page" :itemPerPage="30" :totalElement="total"/>
        </div>
        <MailsTable :mails="emailsSummary"/>

    </div>

  </div>
</template>

<style>


</style>
