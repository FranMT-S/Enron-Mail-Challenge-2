import { computed, ref, watch, type Ref } from "vue"

export const usePagination = (
  InitialPage:Ref<number> = ref(1),
  itemsPerPage:Ref<number> = ref(30),
  Total:Ref<number> = ref(1)
) => {

  const total = ref(Total.value)

  watch(() => Total.value,(newValue) =>{
    if(newValue < 1 )
      newValue = 1

    total.value = newValue
  })

  const totalPages = computed(() => {
    if(total.value <= 0)
      return 1

    return Math.ceil(total.value / itemsPerPage.value)
  });
  if(InitialPage.value > totalPages.value)
    InitialPage.value = totalPages.value

  if(InitialPage.value < 1)
    InitialPage.value = 1

  const currentPage = ref(InitialPage.value)

  const currentFrom = computed(() => {
    const page = currentPage.value <= 0 ? 0 : currentPage.value - 1

    return 1 + ( page * itemsPerPage.value)
  })

  const currentTo = computed(() => {
    const page = currentPage.value <= 0 ? 1 : currentPage.value
    if(currentPage.value == totalPages.value)
      return Total.value

    return  page  * itemsPerPage.value
  })

  const isFirstPage = computed(() => {
    return currentPage.value <= 1
  })

  const isLastPage = computed(() => {
    return currentPage.value == totalPages.value
  })

  const nextPage = () => {
    if(isLastPage.value ){
      return
    }

    currentPage.value++
  }

  const previousPage = () => {
    if(isFirstPage.value ){
      return
    }

    currentPage.value--
  }

  const toPage = (page:number) => {
    if(page < 1 || page > totalPages.value){
      return
    }

    currentPage.value = page
  }

  return {previousPage,nextPage,toPage,totalPages,currentPage,currentTo,currentFrom,isFirstPage,isLastPage}
}
