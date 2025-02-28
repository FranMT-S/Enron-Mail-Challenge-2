<script setup lang="ts">
import { computed, ref, toRef, toRefs } from '@vue/reactivity';
import { onBeforeMount, onMounted, useSlots, useTemplateRef, watch, watchEffect } from 'vue';
import { useSearchStore } from '../store/useSearchStore';
import { storeToRefs } from 'pinia';

interface Props{
  queries?:string[];
  isActive?:boolean
  minLenght?:number
}

const NotNumberLettersOrMailCharactersRegex = /[^A-z0-9 \.@_]|/g
const textRef = ref<HTMLElement | null>(null);
const textInitial = ref("");
const props = withDefaults(defineProps<Props>(),{
  isActive:true,
  minLenght:3,
  queries:undefined
})

const {queries} = toRefs(props)
const {isActive} = toRefs(props)
const { hightlightWords } = useSearchStore()
const ql = ref(!queries.value ?  hightlightWords : queries.value)

const queriesList = computed(() =>{

  var minLenghtAvaible = 3
  if(props.minLenght && props.minLenght > 3)
    minLenghtAvaible  = props.minLenght

  const queryList = new Set<string>()

  let ql:string[] | Set<string>
  ql = !queries.value ?  hightlightWords : queries.value

  for(var query of ql){
    if(query.length >= minLenghtAvaible){
      queryList.add(query)
    }
  }

  return queryList
})

const GetHightLightTag = (queries:Set<string>,text:string) =>{
  if(text.length == 0 || !isActive || queries.size  == 0)
    return ""

  const highlightRegex = createHightlightRegex(queries)
  let match:RegExpExecArray | null = null
  const matches:RegExpExecArray[] = []

  while((match = highlightRegex.exec(text)) != null){
    matches.push(match)
  }

  let textWrappedInMark = WrapTextInnerMarkTag(text,matches)

  return textWrappedInMark
}

const createHightlightRegex = (queries:Set<string>) =>{
  let regexString = ""
  for(var query of queries){
    query = query.replace(NotNumberLettersOrMailCharactersRegex,"")
    regexString += "|"  + query
  }

  // clean the first |
  regexString = regexString.substring(1)
  return new RegExp(regexString,"gi")
}

const WrapTextInnerMarkTag = (text:string,matches:RegExpExecArray[]) =>{
  let currentOffSet = 0
  const beginMarkTag = `<mark class="hightlight-text">`
  const endMarkTag = `</mark>`
  let offset= beginMarkTag.length + endMarkTag.length
  let beforeMatch = ""
  let afterMatch = ""
  let match:RegExpExecArray | null = null

  for(match of matches){
    let word = match[0]
    beforeMatch = text.substring(0,match.index + currentOffSet)
    afterMatch = text.substring(match.index + word.length + currentOffSet)
    text = `${beforeMatch}${beginMarkTag}${word}${endMarkTag}${afterMatch}`
    currentOffSet += offset;
  }

  return text
}

onMounted(() =>{
  if(isActive && textRef.value){
    textInitial.value = textRef.value.innerHTML

    GetHightLightTag(queriesList.value,textInitial.value)
  }
})

const textHighlightTag = computed(() =>{
  return GetHightLightTag(queriesList.value,textInitial.value)
})

watch([textHighlightTag,isActive],() => {
  if(!textRef.value)
    return


  textRef.value.innerHTML = isActive.value ? textHighlightTag.value : textInitial.value;
})

</script>

<template>
  <span ref="textRef" class="content"><slot></slot></span>
</template>
