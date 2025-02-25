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


const route = useRoute();
const router = useRouter()
const emailId = route.params.id as string;

const mail = ref<Mail | null>(null)
const relatedMails = ref<MailSummary[]>([])
const isLoading = ref(true);
const mailsMap = new Map<string,Mail>()
const error = ref<Error | undefined>(undefined)

const {getMailByID,getMails,abortPreviousRequest} = useMailService()

const InitializeData = async (id:string) =>{
  try {
    const _mail =  await getMailByID(id)
    let _mails:MailSummary[] = []
    _mails = await getRelatedMails(_mail)
    mail.value =  _mail
    relatedMails.value =  _mails
    mailsMap.set(_mail.id,_mail)
  } catch (err) {
    console.log("error");
    console.log(err);
    if(!isAbortError(err))
      isLoading.value = false
    error.value = ValidateError(err)
  }
}

const setMail = async (payload:MailSummary) =>{
  try {

    if( mailsMap.has(payload.id)){
      mail.value = mailsMap.get(payload.id)!
    }else{
      isLoading.value = true
      const _mails = await getMailByID(payload.id)


      mailsMap.set(_mails.id,_mails)
      mail.value = _mails
      isLoading.value = false
    }
  } catch (error) {
    if(!isAbortError(error))
      isLoading.value = false
  }
}



const getRelatedMails = async(mail:Mail) =>{
  try {
    let subject = mail.subject.replace(/Re:|FW:/i,"").trim();
    console.log(subject);
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

onUnmounted(() => abortPreviousRequest())

</script>

<template>
<main class="grid grid-cols-[minmax(30vw,1fr)_minmax(50vw,100%)] w-full h-full shadow-lg rounded-3xl col-span-2 row-span-2 p-4  gap-[25px]" >
      <MailsSideBar @OnMail="setMail"  :mails="relatedMails" :idMailSelected="mail?.id" title="Related emails">
        <div class="flex sticky top-0 pt-[6px] bg-[#fff]">
            <BackArrow @click="() => goIndex()"  class="cursor-pointer h-[30px] text-[#7a03ea]  hover:opacity-90 opacity-70 "/>
        </div>
      </MailsSideBar>
      <section
        :class="{'justify-center w-full':isLoading || error, }"
        class="w-full px-4 flex flex-col py-4   bg-white rounded-r-3xl h-full overflow-y-scroll "
      >
        <MailLoader showText v-if="isLoading"/>
        <MailDetails v-else-if="mail != null" :mail="mail"/>
        <div v-else>
          <MailsError :error="error" :onlyTitle="true"/>
        </div>
      </section>
  </main>
</template>


