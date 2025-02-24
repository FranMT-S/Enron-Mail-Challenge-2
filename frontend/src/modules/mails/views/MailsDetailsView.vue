<script setup lang="ts">
import type { Mail } from '@/models/Mails';
import DOMPurify from 'dompurify';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useMailService } from '../composables/useMailService';
import MailDetailsHeader from '../components/maildetails/MailDetailsHeader.vue';

const route = useRoute();
const emailId = route.params.id as string;

const mail = ref<Mail | null>(null)



const {getMailByID} = useMailService()

const handlerGetMailByID = async (id:string) =>{
  try {
    mail.value = await getMailByID(id)
    mail.value.body = DOMPurify.sanitize(mail.value.body,{
      FORBID_ATTR: ['href']
    })
  } catch (error) {
    console.log("error");
    console.log(error);
  }
}

onMounted(async () => {
  await handlerGetMailByID(emailId)
})

</script>

<template>
<main class="flex w-full h-full shadow-lg rounded-3xl col-span-2 row-span-2 p-4  gap-[25px]" >
    <section class="flex flex-col pt-3 w-4/12 bg-gray-50 h-full overflow-y-scroll max-w-[250px]">
      <ul class="mt-6">
        <li class="py-5 border-b px-3 transition hover:bg-indigo-100">
          <a href="#" class="flex justify-between items-center">
            <h3 class="text-lg font-semibold">Akhil Gautam</h3>
            <p class="text-md text-gray-400">23m ago</p>
          </a>
          <div class="text-md italic text-gray-400">You have been invited!</div>
        </li>
      </ul>
    </section>

    <template v-if="mail != null">
      <section class="w-10/12 px-4 flex flex-col bg-white rounded-r-3xl h-full overflow-y-scroll ">
        <MailDetailsHeader :mail="mail"/>
        <section>
          <h1 class="font-bold text-2xl mb-3">{{mail?.subject ?? ""}}</h1>
          <article class=" text-gray-500 leading-7 whitespace-pre-wrap break-words" v-html="mail?.body  ?? ''">
          </article>
        </section>
      </section>
    </template>
    <template v-else>
      <div>
        error
      </div>
    </template>
  </main>
</template>

<style scoped>
/* Estilos adicionales si son necesarios */
</style>


