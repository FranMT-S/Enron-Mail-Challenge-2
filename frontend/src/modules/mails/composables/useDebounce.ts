import type { ActionFN } from "@/models/Function";
import { onUnmounted } from "vue";

export function useDebounce(wait:number){
  var timeout:number

  const debounce = (fn:ActionFN) => {
    clearInterval(timeout)
    timeout = setTimeout(()=>{
      fn()
    }, wait);
  }

  onUnmounted(() => {
    clearInterval(timeout)
  });

  return {debounce}
}
