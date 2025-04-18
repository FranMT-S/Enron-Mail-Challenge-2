<script setup lang="ts">
import { isAbortError, ValidateError } from '@/helpers/errors';
import { navigateToDetails } from '@/helpers/navigate';
import type { MailsProps } from '@/models/mailProps';
import type { Mail, MailSummary } from '@/models/Mails';
import OpenWindowsIcon from '@/modules/components/icons/OpenWindowsIcon.vue';
import { storeToRefs } from 'pinia';
import { onMounted, ref, toRefs, watch } from 'vue';
import { useMailsStore } from '../store/useMailsStore';
import MailsError from './errors/MailsError.vue';
import MailDetails from './MailDetails.vue';
import MailLoader from './MailLoader.vue';
import MailsSideBarItem from './MailsSideBarItem.vue';

const props =  defineProps<MailsProps>()
const { emailSummarySelected } = storeToRefs(useMailsStore())
const { getEmail,resetEmailSummarySelected } = useMailsStore()
const {error:err,isLoading:loading,mails} =  toRefs(props)
const isLoading = ref(loading.value)
const error = ref(err.value)
const mail = ref<Mail | undefined>(undefined)

const onSelectMail = async (mailSummary:MailSummary) =>{
  emailSummarySelected.value = mailSummary
  error.value = undefined
  try {
    isLoading.value = true
    mail.value = await getEmail(mailSummary.id)
    isLoading.value = false
  } catch (err) {
    if(!isAbortError(err))
      isLoading.value = false
    error.value = ValidateError(err)
  }
}

watch(loading, () =>{
  isLoading.value = loading.value
})

watch(mails, () => {
  resetEmailSummarySelected()
})

onMounted(() =>{
  if(emailSummarySelected.value){
    onSelectMail(emailSummarySelected.value)
    const selected = document.querySelector('.selected-item')
    if(selected)
      selected.scrollIntoView({ behavior: "smooth", block: "start" });
  }
})

</script>
<template>
  <div class="mx-5 rounded gap-[10px] bg-[#f3f3f3] flex flex-row h-full overflow-auto border">
    <section class="flex flex-col w-full bg-gray-50 h-full overflow-y-scroll max-w-[250px]">
      <ul>
        <MailsSideBarItem
          class="group"
          v-for="mail in mails" :key="mail.id"
          @click="() => onSelectMail(mail)"
          :mail="mail" :isSelected="mail.id == emailSummarySelected?.id"
        >
          <span
            @click.stop="() => navigateToDetails(mail.id)"
            title="Fullpage"
            class="hidden group-hover:block absolute right-[10px] text-[#630078]  cursor-pointer "
            :class="{'!text-white !block':mail.id == emailSummarySelected?.id,}"
          >
            <OpenWindowsIcon class="h-[22px]"/>
          </span>
        </MailsSideBarItem>
      </ul>
    </section>

    <section
        :class="{'justify-center w-full':isLoading || error, }"
        class="w-full px-4 flex flex-col py-4   bg-white rounded-r-3xl h-full overflow-auto "
      >
      <MailLoader showText v-if="isLoading || loading"/>
      <MailDetails v-else-if="mail != null" :mail="mail"/>
      <MailsError v-else-if="error" :error="error" :onlyTitle="true"/>
      </section>
  </div>
</template>


