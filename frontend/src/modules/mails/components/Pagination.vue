<script setup>

import { usePagination } from '../composables/usePagination';
  const totalElement = 91
  const itemPerPage = 30
  const {currentPage,currentFrom,currentTo,nextPage,previousPage,totalPages,isFirstPage,isLastPage } = usePagination(1,itemPerPage,totalElement)

</script>

<template>
  <div>


    <div>
      <span class="text-sm text-gray-500">{{currentFrom}}-{{ currentTo }} of {{totalElement}}</span>
        <div class="flex items-center space-x-2">
            <button @click="previousPage" class="bg-gray-200    p-1.5 rounded-lg"
            :class="[isFirstPage ? 'text-gray-400' : 'text-gray-700 hover:bg-gray-300']"
            >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd"></path>
                </svg>
            </button>
            <select id="pageSelect" v-model="currentPage">
            <option v-for="page in totalPages" :key="page" :value="page">
               {{ page }}
            </option>
          </select>
            <button @click="nextPage" class="bg-gray-200  p-1.5 rounded-lg " title="Older"
             :class="[isLastPage ? 'text-gray-400' : 'text-gray-700 hover:bg-gray-300']"
            >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"></path>
                </svg>
            </button>
        </div>

    </div>
  </div>
</template>
