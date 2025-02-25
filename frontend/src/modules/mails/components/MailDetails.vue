<script setup lang="ts">
import type { Mail } from '@/models/Mails';
import MailDetailsBody from './MailDetailsBody.vue';
import MailDetailsHeader from './MailDetailsHeader.vue';

interface Props {
  mail:Mail
}

defineProps<Props>()
</script>

<template>
  <MailDetailsHeader :mail="mail"/>
  <MailDetailsBody class="mt-[20px]" :mail="mail"/>
</template>



