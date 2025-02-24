import { TypeVisualization } from "@/enums/TypeVisualization";
import { defineStore } from "pinia";
import { ref } from "vue";



export const useConfigStore = defineStore('config',() =>{

  const isTableViewMode = ref(TypeVisualization.Table)

  return {isTableViewMode}
})
