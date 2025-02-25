<script setup lang="ts">
import { sanatizeBodyMail } from '@/helpers/mails';
import type { Mail } from '@/models/Mails';
import { ref, watchEffect } from 'vue';

interface Props {
  mail:Mail
}

const props = defineProps<Props>()
const body = ref("")
const subject = ref("")

watchEffect(() =>{
  subject.value = props.mail.subject
  body.value = sanatizeBodyMail(props.mail)
})

</script>

<template>
  <section>
    <h1 class="font-bold text-2xl mb-3 break-words">{{subject}}</h1>
    <article class="border border-gray-200 p-[10px] rounded-[10px] text-[#696969] leading-7  whitespace-pre-wrap break-words overflow-auto" v-html="body">
    </article>
  </section>
</template>



