<script setup lang="ts">
import { isAbortError, ValidateError } from '@/helpers/errors';
import type { Mail, MailSummary } from '@/models/Mails';
import { NamesRouter } from '@/routes/routesNames';
import { onMounted, onUnmounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useMailService } from '../composables/useMailService';

import MailsSideBar from '../components/MailsSideBar.vue';
import MailLoader from '../components/MailLoader.vue';
import BackArrow from '@/modules/components/icons/BackArrow.vue';
import MailsError from '../components/errors/MailsError.vue';
import MailDetails from '../components/MailDetails.vue';
import { useMailsStore } from '../store/useMailsStore';


const route = useRoute();
const router = useRouter()
const emailId = route.params.id as string;

const mail = ref<Mail | null>(null)
const relatedMails = ref<MailSummary[]>([])
const isLoading = ref(true);
const mailsMap = new Map<string,Mail>()
const error = ref<Error | undefined>(undefined)

const {getMails} = useMailService()
const {getEmail} = useMailsStore()

const InitializeData = async (id:string) =>{
  let _mails:MailSummary[] = []
  try {
    const _mail =  await getEmail(id)
    _mails = await getRelatedMails(_mail)
    mail.value =  _mail
    relatedMails.value =  _mails
    mailsMap.set(id, _mail)
  } catch (err) {
    if(!isAbortError(err))
      isLoading.value = false

    error.value = ValidateError(err)
  }
}

const setMail = async (payload:MailSummary) =>{
  try {
    isLoading.value = true
    const _mail = await getEmail(payload.id)

    mailsMap.set(_mail.id,_mail)
    mail.value = _mail
    isLoading.value = false
  } catch (error) {
    if(!isAbortError(error))
      isLoading.value = false
  }
}

const getRelatedMails = async(mail:Mail) =>{
  try {
    let subject = mail.subject.replace(/Re:|FW:/i,"").trim();
    let origin = mail.x_origin;
    let to = mail.to;
    let from = mail.from;
    let query = `x_origin:(${origin}) *subject:(${subject}) from:(${from}) *to:(${to})`;
    return (await getMails(query,0,1000)).emails;
  } catch (error) {
    return []
  }
}

const goIndex = () => {
  router.push({name: NamesRouter.Index})
}

onMounted(async () => {
  isLoading.value = true;
  await InitializeData(emailId);
  isLoading.value = false;
})


</script>

<template>
<main class="grid grid-cols-[min(30vw,250px)_minmax(50vw,100%)] w-full h-full shadow-lg rounded-3xl col-span-2 row-span-2 p-4  gap-[25px]" >
      <MailsSideBar @onselectmail="setMail"  :mails="relatedMails" :idMailSelected="mail?.id" title="Similar Subject">
        <div class="flex sticky top-0 pt-[6px] bg-[#fff] z-10">
            <BackArrow @click="() => goIndex()"  class="cursor-pointer h-[30px] text-[#7a03ea]  hover:opacity-90 opacity-70 "/>
        </div>
      </MailsSideBar>
      <section
        :class="{'justify-center w-full':isLoading || error, }"
        class="w-full px-4 flex flex-col py-4   bg-white rounded-r-3xl h-full overflow-auto "
      >
        <MailLoader showText v-if="isLoading"/>
        <MailDetails v-else-if="mail" :mail="mail"/>
        <div v-else>
          <MailsError :error="error" :onlyTitle="true"/>
        </div>
      </section>
  </main>
</template>


