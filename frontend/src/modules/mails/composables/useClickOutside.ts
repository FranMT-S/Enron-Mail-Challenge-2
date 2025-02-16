import { onMounted, onUnmounted, ref } from "vue";

export const useClickOutside = (callback: (event: MouseEvent) => void) =>{
  const formRef = ref<HTMLElement | null>(null);
  const controller = new AbortController()

  const handleClickOutside = (event: MouseEvent) => {
    callback(event)
  };

  onMounted(() => {
    document.addEventListener("click", handleClickOutside,{signal:controller.signal});
  });

  onUnmounted(() => {
    controller.abort()
  });

  return {formRef}
}
