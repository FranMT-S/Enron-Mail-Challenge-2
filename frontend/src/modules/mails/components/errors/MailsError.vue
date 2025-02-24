<script setup lang="ts">
import ErrorComponent from './ErrorComponent.vue';
import MailsNotFoundError from './MailsNotFoundError.vue';

interface Props{
  error?:Error,
  onlyTitle:boolean
}

const {error,onlyTitle} = defineProps<Props>()

</script>
<template>
  <ErrorComponent :error="error" v-if="error"/>
  <MailsNotFoundError :onlyTitle="onlyTitle" v-else/>
</template>
