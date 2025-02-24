import { defineStore } from "pinia";
import { ref } from "vue";

export const useSearchStore = defineStore('search',() =>{

  const queryString = ref("")
  const page = ref(1)
  const itemPerPage = ref(100)



  return {queryString,page,itemPerPage}
})
