import { fieldRegex, specialCharactersRegex } from "@/helpers/regex";
import type { SortByItem, SortType } from "@/models/Sort";

import { defineStore } from "pinia";
import { computed, ref } from "vue";
export const useSearchStore = defineStore('search',() =>{


  const queryString = ref("")
  const page = ref(1)
  const itemPerPage = ref(100)
  const isHightlightActive = ref(true)
  const sortValue = ref<SortByItem>({
    field:'date',
    sort:'desc',
    label:'Date',
    sortLabel:'Newest'
  })

  const setSortField = (newSort:SortByItem) => {
    sortValue.value = newSort
  }

  const hightlightWords = computed(() =>{
    let words = []
    let match:RegExpExecArray | null = null
    const fields = new RegExp(fieldRegex,'gi')
    let query = queryString.value

    // find fields and value
    while((match = fields.exec(query)) != null){
      let value = match[1]
      if(!value)
        continue

      value = cleanValue(value)
      if(value == "")
        continue

      let wordsList = GetWordsHightlight(value)

      words.push(...wordsList)
      query = query.replace(match[0],"")
    }

    query = cleanValue(query)
    if(query.length > 0)
      words.push(...GetWordsHightlight(query))


    return  words.sort((a,b) => b.length - a.length)
  })


  return {queryString,page,itemPerPage,isHightlightActive,hightlightWords,sortValue,setSortField}
})

// Auxilar functions
function cleanValue(value:string){
  value = value.replace(/[()]/g,"")
  value = value.trim()
  value = value == "." ? "" : value
  return value
}

function NotExcludeWordAndNotEmpty(value:string):boolean{
  return value.length > 0 && value[0] != "-"
}

function splitWordBySpace(value:string){
  return value.split(" ").map(v => v.trim()).filter(NotExcludeWordAndNotEmpty)
}

function splitUserFromMail(value:string){
  const userMail = value.split("@");
  if(userMail.length == 1)
    return []


  return userMail[0].split('.').map(v => v.trim()).filter(Boolean)
}

function GetWordsHightlight(value:string){
  let words:string[] = []
  let wordsList = splitWordBySpace(value)
  // split user words from mail
  for(var w of wordsList){
    if(w[0] == "-")
      continue

    let nameSplited = splitUserFromMail(w)
    words.push(...nameSplited)
  }
  words.push(...wordsList)
  return words
}
