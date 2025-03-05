<script setup lang="ts">
import { formatDate } from '@/helpers/formatDate';
import { formatXField } from '@/helpers/mails';
import type { Mail } from '@/models/Mails';
import { onBeforeMount, onMounted, onUnmounted, onUpdated, ref, watch, watchEffect } from 'vue';
import Chip from './Chip.vue';
import Highlight from './Highlight.vue';
import ArrowUp from '@/modules/components/icons/ArrowUp.vue';

interface Props {
  mail:Mail
}

interface ToPerson{
  to:string,
}

const props = defineProps<Props>()
const xFrom = ref("")
const xToList = ref<string[]>([])
const ToList = ref<string[]>([])
const dateFormated = ref("")
const chipListRef = ref<HTMLDivElement | null>(null)
const isShowCompleteList = ref(false)
const hasScrollbarHide = ref(false)
const minHeight = ref(28)

const showHide = () =>{
  const chipLisContainer = chipListRef.value
  if(chipLisContainer){

    isShowCompleteList.value = !isShowCompleteList.value
  }
}

watchEffect(() =>{
  dateFormated.value = formatDate(props.mail.date)
  if(props.mail.x_from.length > 0){
    xFrom.value = props.mail.x_from.split("<")[0].trim()
    xToList.value =  formatXField(props.mail.x_to)
  }

  ToList.value = props.mail.to.split(/,\s*/).filter(Boolean)
})

const controller = new AbortController()
const checkScrollbar = () =>{
  const chipLisContainer = chipListRef.value
  if(chipLisContainer){
    hasScrollbarHide.value = chipLisContainer.scrollHeight > minHeight.value ? true : false
  }
}


onMounted(() => {
  checkScrollbar();
  window.addEventListener("resize", checkScrollbar,{signal:controller.signal});
});

onUnmounted(() => {
  controller.abort()
});



</script>

<template>
  <header :id="mail.id" class="border px-[8px] rounded-[10px] flex flex-col gap-[10px] text-left items-start py-[20px] mb-[10px]  ">
    <div class="flex flex-col justify-between  w-full">
      <h3 class="font-semibold  text-light text-[#939393]">{{dateFormated}}</h3>
      <div class="flex flex-col gap-[3px]">
        <h3 class="font-semibold text-lg text-ellipsis text-nowrap overflow-hidden">{{xFrom}} <{{mail.from}}></h3>
        <div
          :class="{'active':isShowCompleteList,}"
          class="chip-list  chip-list ">
          <div ref="chipListRef" :style="{minHeight:minHeight + 'px'}" class=" flex flex-row gap-[4px] flex-wrap text-gray-700   overflow-hidden">
            <!-- <button v-if="hasScrollbarHide" @click="showHide" class="px-2 border rounded">test</button> -->
            <ArrowUp v-if="hasScrollbarHide" @click="showHide" class="text-gray-400 hover:text-gray-500 h-[24px]"
            :class="{'rotate-180 !text-[#681dca]':isShowCompleteList,}"/>
            To: <Chip textColor="#d3f2ff" backgroundColor="#681dca"  v-for="(xTo,i) in xToList">
              <Highlight>
                {{ xTo }} <p v-if="ToList.length == xToList.length" class="contents text-[#ffdbb5]"><{{ToList[i]}}></p>
              </Highlight>
            </Chip>

          </div>
        </div>
        <!-- <div>
          TO:{{ mail.to }},
          <hr>
          XTO:{{ mail.x_to}}
          TO:{{ ToList.length }}, XTO:{{ xToList.length}}
        </div> -->
      </div>
    </div>
  </header>
</template>


<style>


  .chip-list {
    display: grid;
    grid-template-rows: 0fr;

    transition: grid-template-rows 0.3s ease-in-out;
  }

  .chip-list.active {
    grid-template-rows: 1fr;
  }
</style>
