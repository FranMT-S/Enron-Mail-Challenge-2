<script setup lang="ts">
import AlertTriangleIcon from '@/modules/components/icons/AlertTriangleIcon.vue';
import WifiErrorIcon from '@/modules/components/icons/WifiErrorIcon.vue';
import WifiNoConnectIcon from '@/modules/components/icons/WifiNoConnectIcon.vue';
import MagnifyingGlassImage from '../images/MagnifyingGlassImage.vue';

import { ResponseError } from '@/models/Response';
import { onMounted, ref, shallowRef } from 'vue';

interface Props{
  error:Error
}

const iconsSvg = {
  wifiError: WifiErrorIcon,
  alert: AlertTriangleIcon,
  wifiNoConnectIcon: WifiNoConnectIcon,
};

const props = defineProps<Props>();
const Image = shallowRef(iconsSvg.wifiError)
const message = ref("")
const code = ref(500)

const validateError = (error:Error) =>{
  if(error instanceof ResponseError){
    message.value = error.message
    code.value = error.status
    switch (error.status) {
      case 503:
        Image.value = iconsSvg.wifiError
        break;
      case 408:
        Image.value = iconsSvg.wifiError
        break;
      default:
        Image.value = iconsSvg.alert
        break;
    }
  }else{
    Image.value = iconsSvg.alert
    code.value = 500
    message.value = "There was a problem try again later"
  }
}

onMounted(() => validateError(props.error))
</script>

<template>
  <div class="w-full flex flex-col justify-center items-center ">
      <div
        :class="{'max-w-[300px]':code == 404}"
        class="flex flex-col max-w-[150px]">

        <MagnifyingGlassImage v-if="code == 404" class="ico w-full text-[#4f108e]" :class="[':text-[#c338c3]']"/>
        <Image v-else class="ico w-full text-[#4f108e]" :class="[':text-[#c338c3]']"/>
      </div>
      <h1 class=" text-[#5000a0] text-lg font-semibold text-center ">{{message}}</h1>
  </div>
</template>

