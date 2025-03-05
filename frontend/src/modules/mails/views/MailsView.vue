<script setup lang="ts">
import { TypeVisualization } from '@/enums/TypeVisualization';
import { isAbortError, ValidateError } from '@/helpers/errors';
import { navigateToDetails } from '@/helpers/navigate';
import type { MailSummary } from '@/models/Mails';
import { storeToRefs } from 'pinia';
import { onMounted, ref, watch } from 'vue';
import MailsTable from '../components/MailsTable.vue';
import MailVerticalSplit from '../components/MailVerticalSplit.vue';
import NavBar from '../components/NavBar.vue';
import ViewOptionsBar from '../components/OptionsBar.vue';
import Pagination from '../components/Pagination.vue';
import { useDebounce } from '../composables/useDebounce';
import { useConfigStore } from '../store/useConfig';
import { useMailsStore } from '../store/useMailsStore';
import { useSearchStore } from '../store/useSearchStore';
import SortBySelect from '../components/SortBySelect.vue';
import type { SortField } from '@/models/Sort';

const itemPerPageList = ref([100,500,1000])
const {queryString,itemPerPage,page,sortValue} = storeToRefs(useSearchStore())
const {fillEmailsSummary,resetEmailSummarySelected } = useMailsStore()
const {emailsSummary,total} = storeToRefs(useMailsStore())
const {visualizationMode} = storeToRefs(useConfigStore())

itemPerPage.value = itemPerPageList.value[0]

const isLoading = ref(false)
const error = ref<Error | undefined>()

const {debounce} = useDebounce(200)

const handlerGetMails = async (query:string = "", page:number = 1, size:number=  30,sortList:SortField[] = []) => {
  try {
    isLoading.value = true
    error.value = undefined
    await fillEmailsSummary(query,page,size,sortList)
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
const onChangeQuery = async (newQuery:string) => {
  queryString.value = newQuery;
  page.value = 1;

  // force request when press button
  if(queryString.value == newQuery)
    await handlerGetMails(newQuery,page.value,itemPerPage.value,[sortValue.value])
}

const onChangeVisualization = (v:any) =>{
  resetEmailSummarySelected()
}

watch([page,itemPerPage],([newPage,newTotalPerPage,]) => {
  debounce(async () =>{
    await handlerGetMails(queryString.value,newPage,newTotalPerPage,[sortValue.value])
  })
})

watch(sortValue,async () => {
  if(page.value ==1)
    await handlerGetMails(queryString.value,page.value,itemPerPage.value,[sortValue.value])

  page.value = 1
})

onMounted(async () => {
  if(emailsSummary.value.length > 0)
    return

  await handlerGetMails(queryString.value,1,itemPerPage.value,[sortValue.value])
})

</script>

<template>
  <NavBar @on-search="onChangeQuery"/>
  <div   class="w-full    flex overflow-x-auto custom-scrollbar px-[20px] mb-5">
    <div class="flex-1 px-2 grid grid-rows-[auto_1fr] bg-white shadow-xl rounded-lg" >
      <div class="flex flex-row justify-between mx-[20px] max-mobile:flex-col ">
          <div class="flex items-center justify-end py-[10px] mobile:py-[16px] gap-[4px] mobile:justify-center">
            <ViewOptionsBar @onchange="onChangeVisualization"/>
            <SortBySelect />
          </div>
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
