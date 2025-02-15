<script setup lang="ts">
import type { Mail } from '@/models/Mails';
import { defineProps, h, onMounted, ref, type VNode } from 'vue';
import { useRoute } from 'vue-router';
import { useMailService } from '../composables/useMailService';
import DOMPurify from 'dompurify';

const route = useRoute();
const emailId = route.params.id as string;

const mail = ref<Mail | null>(null)


const {getMailByID} = useMailService()

const handlerGetMailByID = async (id:string) =>{
  try {
    mail.value = await getMailByID(id)
    console.log(mail.value.body);
    mail.value.body = DOMPurify.sanitize(mail.value.body,{
      FORBID_ATTR: ['href']
    })
  } catch (error) {
    console.log("no se puedo obtener el correo");
    console.log(error);
  }
}

onMounted(async () => await handlerGetMailByID(emailId))

</script>

<template>
  <div class="max-w-2xl mx-auto bg-white shadow-lg rounded-xl p-6 border border-gray-200">
    <div class="mb-4">
      <h2 class="text-lg font-semibold text-gray-800">Correo Electrónico</h2>
    </div>
    <div class="text-sm text-gray-600">
      <p><strong>De:</strong> {{ mail?.from ?? "" }}</p>
      <p><strong>Para:</strong> {{ mail?.to ?? ""}}</p>
      <p><strong>Fecha:</strong> {{ mail?.date ?? ""}}</p>
    </div>
    <div class="mt-4 p-4 bg-gray-100 rounded-lg text-gray-700 whitespace-pre-wrap" v-html="mail?.body">

    </div>
  </div>
</template>

<style scoped>
/* Estilos adicionales si son necesarios */
</style>
