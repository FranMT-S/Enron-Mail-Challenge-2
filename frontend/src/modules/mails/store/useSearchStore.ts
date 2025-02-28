import { fieldRegex, specialCharactersRegex } from "@/helpers/regex";
import { defineStore } from "pinia";
import { computed, ref } from "vue";

function cleanValue(value:string){
  value = value.replace(/[()]/g,"")
  value = value.trim()
  value = value == "." ? "" : value
  return value
}

function wordListWithTrim(value:string){
  return value.split(" ").map(v => v.trim()).filter(Boolean)
}

export const useSearchStore = defineStore('search',() =>{

  const queryString = ref("")
  const page = ref(1)
  const itemPerPage = ref(100)
  const isHightlightActive = ref(true)

  const hightlightWords = computed(() =>{
    let words = []
    let match:RegExpExecArray | null = null
    const regex = new RegExp(fieldRegex,'gi')
    let query = queryString.value

    // find fields and value
    while((match = regex.exec(query)) != null){
      let value = match[1]
      if(!value)
        continue

      value = cleanValue(value)
      if(value == "")
        continue

      let cleanWords = wordListWithTrim(value)
      words.push(...cleanWords)
      query = query.replace(match[0],"")
    }

    query = cleanValue(query)
    if(query == "")
      return words

    words.push(...wordListWithTrim(query))
    return words
  })


  return {queryString,page,itemPerPage,isHightlightActive,hightlightWords}
})
