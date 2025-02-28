<template>
  <v-container>
    <v-date-picker v-model="date"></v-date-picker>
  </v-container>
</template>

<script setup lang="ts">
  import { ref } from 'vue'

  const date = ref(null)
</script>
