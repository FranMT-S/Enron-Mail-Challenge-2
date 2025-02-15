<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  from: String,
  to: String,
  date: String,
  body: String
});
</script>

<template>
  <div class="max-w-2xl mx-auto bg-white shadow-lg rounded-xl p-6 border border-gray-200">
    <div class="mb-4">
      <h2 class="text-lg font-semibold text-gray-800">Correo Electrónico</h2>
    </div>
    <div class="text-sm text-gray-600">
      <p><strong>De:</strong> {{ props.from }}</p>
      <p><strong>Para:</strong> {{ props.to }}</p>
      <p><strong>Fecha:</strong> {{ props.date }}</p>
    </div>
    <div class="mt-4 p-4 bg-gray-100 rounded-lg text-gray-700 whitespace-pre-wrap">
      {{ props.body }}
    </div>
  </div>
</template>

<style scoped>
/* Estilos adicionales si son necesarios */
</style>
