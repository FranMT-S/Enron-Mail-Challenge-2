<script setup lang="ts">
import { formatDate } from '@/helpers/formatDate';
import type { Mail } from '@/models/Mails';
import { ref } from 'vue';
import BackArrow from '../../../components/icons/BackArrow.vue';
import { NamesRouter } from '@/routes/routesNames';

interface Props {
  mail:Mail
}

const {mail} = defineProps<Props>()

const xFrom = ref("")
const dateFormated = ref("")
dateFormated.value = formatDate(mail.date)
if(mail.x_from.length > 0){
  xFrom.value = mail.x_from.split("<")[0].trim()
}
</script>

<template>
  <header class="flex flex-col gap-[10px] text-left items-start my-[20px]   h-48 border-b-2  ">
      <div class="flex  justify-between  w-full">
        <ul class="flex text-gray-400 space-x-4">
            <RouterLink :to="{name:NamesRouter.Index}"   >
                <BackArrow class="cursor-pointer text-[#7a03ea]  hover:opacity-90 opacity-70 "/>
            </RouterLink>
          </ul>
          <h3 class="font-semibold  text-light text-[#62219f]">{{dateFormated}}</h3>
      </div>
        <div class="flex space-x-4 items-center">
          <div class="flex flex-col">
            <h3 class="font-semibold text-lg">{{xFrom}}</h3>
            <p class="text-light text-gray-400">{{mail?.from ?? ""}}</p>
          </div>
        </div>
        <div>

        </div>
      </header>
</template>



