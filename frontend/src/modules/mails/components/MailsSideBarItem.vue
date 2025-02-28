<script setup lang="ts">
import { formatDate } from '@/helpers/formatDate';
import { formatXField } from '@/helpers/mails';
import type { MailSummary } from '@/models/Mails';
import { ref } from 'vue';
import Highlight from './Highlight.vue';

interface Props {
  mail:MailSummary,
  isSelected?:boolean
}

const {mail,isSelected} = defineProps<Props>()

const dateFormated = ref("")
const xFrom = ref("")
const xToList = ref<string[]>([])
const xTo = ref<string>("")

dateFormated.value = formatDate(mail.date)

xToList.value =  formatXField(mail.x_to).map((v) => v.replace(",",""))
xFrom.value = mail.x_from.split("<")[0].trim()
xTo.value = xToList.value.join(",")
</script>

<template>
    <li :class="{'bg-[#580696]': isSelected,'hover:bg-[#ebebff] transition-colors':!isSelected}"
      class="py-5 border-b px-3 transition relative cursor-pointer"
    >
      <slot></slot>
      <div :class="{'font-medium !text-[#ba86ef]':isSelected}" class="text-base break-words italic text-gray-600">{{dateFormated}}</div>
      <div  class="flex flex-col justify-between  w-full ">
        <h3 :class="{'text-[#e5ceff]':isSelected}" class="text-lg font-semibold     text-nowrap text-ellipsis w-full overflow-hidden">
          <Highlight>
            {{ xFrom }}
          </Highlight>
        </h3>
        <p :class="{'!text-[#d1cdd2]':isSelected}" class="text-md text-gray-700    text-nowrap text-ellipsis w-full overflow-hidden">To:
          <Highlight>
            {{ xTo }}
          </Highlight>
        </p>
        <h3 :class="{'!text-[#ffe1e1]':isSelected}" class="text-lg font-semibold text-gray-800      text-nowrap text-ellipsis w-full overflow-hidden">
          <Highlight>
            {{ mail.subject }}
          </Highlight>
        </h3>
      </div>
    </li>
</template>



