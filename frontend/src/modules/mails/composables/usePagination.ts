import { computed, ref } from "vue"

export const usePagination = (InitialPage:number,itemsPerPage:number,Total:number) => {

  const currentPage = ref(InitialPage)
  const totalPages = computed(() => Math.ceil(Total / itemsPerPage));

  const currentFrom = computed(() => {
    const page = currentPage.value -1
    return page == 0 ? 1 : page * itemsPerPage
  })

  const currentTo = computed(() => {
    return currentPage.value == totalPages.value ? Total : currentPage.value * itemsPerPage
  })

  const isFirstPage = computed(() => {
    return currentPage.value == 1
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
