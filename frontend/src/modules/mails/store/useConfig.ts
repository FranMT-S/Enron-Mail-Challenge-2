import { TypeVisualization } from "@/enums/TypeVisualization";
import { computed } from "@vue/reactivity";
import { defineStore } from "pinia";
import { onMounted, ref, watch } from "vue";

const key = 'visualization-mode'
const initializeValue = () =>{
  let value = Number(localStorage.getItem(key))
  if(TypeVisualization[value]){
    return value
  }

  return TypeVisualization.Table
}

export const useConfigStore = defineStore('config',() =>{

  const visualizationMode = ref(initializeValue())

  watch(visualizationMode, (newValue) =>{
    localStorage.setItem(key,newValue.toString())
  })

  onMounted(() =>{
    let value = Number(localStorage.getItem(key))
    if(TypeVisualization[value]){
      return visualizationMode.value = value
    }

    visualizationMode.value = TypeVisualization.Table
  })



  return {visualizationMode}
})
