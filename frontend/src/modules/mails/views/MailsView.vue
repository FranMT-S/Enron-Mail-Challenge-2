<script setup lang="ts">
  import type { MailSummary } from '@/models/Mails';
import type { ResponseMailSummary } from '@/models/Response';
import { onBeforeMount, onMounted, reactive, ref, watch } from 'vue';
import MailsTable from '../components/MailsTable.vue';
import NavBar from '../components/NavBar.vue';
import OptionsBar from '../components/OptionsBar.vue';
import Pagination from '../components/Pagination.vue';
import { useMailService } from '../composables/useMailService';
import { useDebounce } from '../composables/useDebounce';

const emailsSummary = ref<MailSummary[]>([])
const total = ref(1)
// const query = ref("este es un ejemplo")
const query = ref("")
const page = ref(1)
const {getMails} = useMailService()
const itemPerPageList = [100,500,1000]
const totalPerPage = ref(itemPerPageList[0])
const {debounce} = useDebounce(400)

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


watch([query,page,totalPerPage],([newQuery,newPage,newTotalPerPage]) => {
  if(query.value != newQuery){
    newPage  = 1
    page.value  = newPage
  }
  debounce(() =>{
    handlerGetMails(newQuery,newPage,newTotalPerPage)
  })
})

const onChangePage = (newPage:number) => page.value = newPage
const onChangeItemPerPage = (newValue:number) => totalPerPage.value = newValue
const onChangeQuery = (newQuery:string) => query.value = newQuery


onMounted(async () => await handlerGetMails(query.value,1,totalPerPage.value))

</script>

<template>
  <NavBar @on-search="onChangeQuery"/>

  <div class="w-full    flex overflow-x-auto custom-scrollbar px-[20px] mb-5">
    <div class="flex-1 px-2 grid grid-rows-[64px_1fr] bg-white shadow-xl rounded-lg" >
        <div class="flex justify-between mx-[20px]">
          <OptionsBar/>
          <Pagination
            @on-change-page="onChangePage"
            @on-item-per-list-change="onChangeItemPerPage"
            :itemPerPageList="itemPerPageList"
            :page="page"
            :itemPerPage="totalPerPage"
            :totalElement="total"
          />
        </div>
        <MailsTable :mails="emailsSummary"/>
    </div>
  </div>
</template>

<style>


</style>
