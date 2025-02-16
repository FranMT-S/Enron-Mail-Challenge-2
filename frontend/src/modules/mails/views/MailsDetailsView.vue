<script setup lang="ts">
import type { Mail } from '@/models/Mails';
import { defineProps, h, onMounted, ref, type VNode } from 'vue';
import { useRoute } from 'vue-router';
import { useMailService } from '../composables/useMailService';
import DOMPurify from 'dompurify';
import { formatDate } from '@/helpers/formatDate';

const route = useRoute();
const emailId = route.params.id as string;

const mail = ref<Mail | null>(null)
const dateFormated = ref("")
const xFrom = ref("")

const {getMailByID} = useMailService()

const handlerGetMailByID = async (id:string) =>{
  try {
    mail.value = await getMailByID(id)
    mail.value.body = DOMPurify.sanitize(mail.value.body,{
      FORBID_ATTR: ['href']
    })
    dateFormated.value = formatDate(mail.value.date)
    if(mail.value.x_from.length > 0){
      xFrom.value = mail.value.x_from.split("<")[0].trim()
    }
    console.log(mail.value);
  } catch (error) {
    console.log("no se puedo obtener el correo");
    console.log(error);
  }
}

onMounted(async () => await handlerGetMailByID(emailId))

</script>

<template>
  <!-- <div class="mx-auto bg-white  p-6  ">

    <div class="text-sm text-gray-600">
      <p><strong>De:</strong> {{ mail?.from ?? "" }}</p>
      <p><strong>Para:</strong> {{ mail?.to ?? ""}}</p>
      <p><strong>Fecha:</strong> {{ dateFormated}}</p>
    </div>
    <div class="mt-4 p-4 bg-gray-100 rounded-lg text-gray-700 whitespace-pre-wrap" v-html="mail?.body">
    </div>
  </div> -->

  <!-- component -->
<main class="flex w-full h-full shadow-lg rounded-3xl col-span-2 row-span-2 p-4" >
    <section class="flex flex-col pt-3 w-4/12 bg-gray-50 h-full overflow-y-scroll">
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

    <section class="w-6/12 px-4 flex flex-col bg-white rounded-r-3xl h-full overflow-y-scroll">
      <div class="flex justify-between items-center h-48 border-b-2 py-10 mb-4 ">
        <div class="flex space-x-4 items-center">
          <div class="flex flex-col">
            <h3 class="font-semibold text-lg">{{xFrom}}</h3>
            <p class="text-light text-gray-400">{{mail?.from ?? ""}}</p>
          </div>
        </div>
        <div>
          <ul class="flex text-gray-400 space-x-4">
            <li class="w-6 h-6">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M11 15l-3-3m0 0l3-3m-3 3h8M3 12a9 9 0 1118 0 9 9 0 01-18 0z" />
              </svg>
            </li>
            <li class="w-6 h-6">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M13 9l3 3m0 0l-3 3m3-3H8m13 0a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </li>

            <li class="w-6 h-6">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z" />
              </svg>
            </li>
            <li class="w-6 h-6">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </li>
            <li class="w-6 h-6">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
              </svg>
            </li>
          </ul>
        </div>
      </div>
      <section>
        <h1 class="font-bold text-2xl mb-3">{{mail?.subject ?? ""}}</h1>
        <!-- <article class="mt-8 text-gray-500 leading-7 tracking-wider"> -->
        <article class=" text-gray-500 leading-7 break-words" v-html="mail?.body  ?? ''">

        </article>
      </section>
    </section>
  </main>
</template>

<style scoped>
/* Estilos adicionales si son necesarios */
</style>


