<script setup lang="ts">
  import {ref} from 'vue'

  const controller = ref<AbortController | null>(null);
  const stateRequest = ref<string>("Sin estado");

  const fetchData = async () => {

    // Cancela la petición anterior si existe
    if (controller.value) {
      controller.value.abort();
    }

    stateRequest.value = "Pendiente"

    // Crea un nuevo controlador de aborto
    controller.value = new AbortController();
    const signal = controller.value.signal;

    try {
      const res = await fetch("http://localhost:3000/long-request", {
        signal,
        method:"Post"
      });
      stateRequest.value = await res.text();
    } catch (err) {
      if (err instanceof DOMException && err.name === "AbortError") {
        // stateRequest.value = "Petición cancelada manualmente";
      } else {
        stateRequest.value = "Error en la petición";
      }
    }
  };

  const cancelRequest = () => {
    if (controller.value) {
      controller.value.abort();
    }
  };

</script>

<template>
  <div class="w-[500px] flex flex-col">
    <p class="text mt-5 font-semibold text-center">{{ stateRequest }}</p>
      <div class="text-center">
        <button v-on:click="fetchData" class="border p-1 mt-4">
          Request
      </button>
      <button v-on:click="cancelRequest" class="border p-1 mt-4 mx-3">
          Cancelar
      </button>
    </div>

  </div>
</template>

<style>


</style>
