<script setup lang="ts">
import { sanitizeBodyMail } from '@/helpers/mails';
import type { Mail } from '@/models/Mails';
import { ref, watchEffect } from 'vue';
import Highlight from './Highlight.vue';

interface Props {
  mail:Mail
}

const props = defineProps<Props>()
const body = ref("")
const subject = ref("")

watchEffect(() =>{
  subject.value = props.mail.subject
  body.value = sanitizeBodyMail(props.mail)
})

</script>

<template>
  <section>
    <h1 class="font-bold text-2xl mb-3 break-words"><Highlight>{{subject}}</Highlight></h1>
    <article class="border w-full border-gray-200 p-[10px] rounded-[10px] text-[#696969] leading-7  whitespace-pre-wrap break-words overflow-auto">
      <Highlight v-html="body"></Highlight>
    </article>
  </section>
</template>



