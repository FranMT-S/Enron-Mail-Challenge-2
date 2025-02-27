<script setup lang="ts">
import { isAbortError, ValidateError } from '@/helpers/errors';
import type { MailSummary } from '@/models/Mails';
import { type ResponseMailSummary } from '@/models/Response';
import { storeToRefs } from 'pinia';
import { onMounted, onUnmounted, ref, watch } from 'vue';
import MailsTable from '../components/MailsTable.vue';
import NavBar from '../components/NavBar.vue';
import OptionsBar from '../components/OptionsBar.vue';
import Pagination from '../components/Pagination.vue';
import { useDebounce } from '../composables/useDebounce';
import { useSearchStore } from '../store/useSearchStore';
import { useMailsStore } from '../store/useMailsStore';
import { useConfigStore } from '../store/useConfig';
import { TypeVisualization } from '@/enums/TypeVisualization';
import MailVerticalSplit from '../components/MailVerticalSplit.vue';
import router from '@/routes';
import { NamesRouter } from '@/routes/routesNames';
import { navigateToDetails } from '@/helpers/navigate';

const itemPerPageList = [100,500,1000]
const {queryString,itemPerPage,page} = storeToRefs(useSearchStore())
const {fillEmailsSummary,resetEmailSummarySelected } = useMailsStore()
const {emailsSummary,total} = storeToRefs(useMailsStore())
const {visualizationMode} = storeToRefs(useConfigStore())

itemPerPage.value = itemPerPageList[0]

const isLoading = ref(false)
const error = ref<Error | undefined>()

const {debounce} = useDebounce(200)

const handlerGetMails = async (query:string = "", page:number = 1, size:number=  30) => {
  try {
    isLoading.value = true
    error.value = undefined
    await fillEmailsSummary(query,page,size)
    isLoading.value = false
  } catch (err) {
    if(!isAbortError(err))
      isLoading.value = false

    error.value = ValidateError(err)
  }
};

const onSelectMailTable = (mail:MailSummary) =>{
  navigateToDetails(mail.id)
}

const onChangePage = (newPage:number) => page.value = newPage
const onChangeItemPerPage = (newValue:number) => itemPerPage.value = newValue
const onChangeQuery = (newQuery:string) => {
  queryString.value = newQuery;
  page.value = 1;
}

const onChangeVisualization = (v:any) =>{
  resetEmailSummarySelected()
}

watch([queryString,page,itemPerPage],([newQuery,newPage,newTotalPerPage]) => {
  debounce(async () =>{
    await handlerGetMails(newQuery,newPage,newTotalPerPage)
  })
})

onMounted(async () => {
  if(emailsSummary.value.length > 0)
    return

  await handlerGetMails(queryString.value,1,itemPerPage.value)
})

</script>

<template>
  <NavBar @on-search="onChangeQuery"/>
  <div   class="w-full    flex overflow-x-auto custom-scrollbar px-[20px] mb-5">
    <div class="flex-1 px-2 grid grid-rows-[auto_1fr] bg-white shadow-xl rounded-lg" >
        <div class="flex flex-row justify-between mx-[20px] max-mobile-sm:flex-col ">
          <OptionsBar @onchange="onChangeVisualization"/>
          <Pagination
            @on-change-page="onChangePage"
            @on-item-per-list-change="onChangeItemPerPage"
            :itemPerPageList="itemPerPageList"
            :page="page"
            :itemPerPage="itemPerPage"
            :totalElement="total"
            />
        </div>
        <MailVerticalSplit v-if="visualizationMode == TypeVisualization.Vertical" :mails="emailsSummary" :error="error" :isLoading="isLoading"/>
        <MailsTable @onselectmail="onSelectMailTable" v-else :mails="emailsSummary" :error="error" :isLoading="isLoading"/>
      </div>
    </div>
</template>

<style>


</style>
