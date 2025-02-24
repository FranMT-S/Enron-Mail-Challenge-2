<script setup lang="ts">
import { ValidateError } from '@/helpers/errors';
import type { MailSummary } from '@/models/Mails';
import { type ResponseMailSummary } from '@/models/Response';
import { storeToRefs } from 'pinia';
import { onMounted, ref, watch } from 'vue';
import MailsTable from '../components/MailsTable.vue';
import NavBar from '../components/NavBar.vue';
import OptionsBar from '../components/OptionsBar.vue';
import Pagination from '../components/Pagination.vue';
import { useDebounce } from '../composables/useDebounce';
import { useMailService } from '../composables/useMailService';
import { useSearchStore } from '../store/useSearchStore';

const itemPerPageList = [100,500,1000]
const emailsSummary = ref<MailSummary[]>([])
const store = useSearchStore()
const {queryString,itemPerPage,page} = storeToRefs(store)
itemPerPage.value = itemPerPageList[0]

const isLoading = ref(false)
const error = ref<Error | undefined>()
const total = ref(1)
const {getMails} = useMailService()
const {debounce} = useDebounce(200)

const handlerGetMails = async (query:string = "", page:number = 1, size:number=  30) => {
  try {
    isLoading.value = true
    error.value = undefined
    const mails:ResponseMailSummary = await getMails(query,page,size)
    total.value = mails.total
    emailsSummary.value = mails.emails
  } catch (err) {
    error.value = ValidateError(err)
  }finally{
    isLoading.value = false
  }
};


watch([queryString,page,itemPerPage],([newQuery,newPage,newTotalPerPage]) => {
  if(queryString.value != newQuery){
    newPage  = 1
    page.value  = newPage
  }
  debounce(() =>{
    handlerGetMails(newQuery,newPage,newTotalPerPage)
  })
})

const onChangePage = (newPage:number) => page.value = newPage
const onChangeItemPerPage = (newValue:number) => itemPerPage.value = newValue
const onChangeQuery = (newQuery:string) => queryString.value = newQuery


onMounted(async () => await handlerGetMails(queryString.value,1,itemPerPage.value))

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
            :itemPerPage="itemPerPage"
            :totalElement="total"
          />
        </div>
        <MailsTable :mails="emailsSummary" :error="error" :isLoading="isLoading"/>
    </div>
  </div>
</template>

<style>


</style>
