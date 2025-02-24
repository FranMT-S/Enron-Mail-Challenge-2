<script setup lang="ts">
import AlertTriangleIcon from '@/modules/components/icons/AlertTriangleIcon.vue';
import WifiErrorIcon from '@/modules/components/icons/WifiErrorIcon.vue';
import WifiNoConnectIcon from '@/modules/components/icons/WifiNoConnectIcon.vue';

import { ResponseError } from '@/models/Response';
import { onMounted, ref, shallowRef } from 'vue';

interface Props{
  error:Error
}

const icons = {
  wifiError: WifiErrorIcon,
  alert: AlertTriangleIcon,
  wifiNoConnectIcon: WifiNoConnectIcon
};

const props = defineProps<Props>();
const Image = shallowRef(icons.wifiError)
const message = ref("")

const validateError = (error:Error) =>{
  if(error instanceof ResponseError){
    message.value = error.message
    switch (error.status) {
      case 503:
        Image.value = icons.wifiNoConnectIcon
        break;
      case 408:
        Image.value = icons.wifiError
        break;

      default:
        Image.value = icons.alert
        break;
    }
  }else{
  }
}

onMounted(() => validateError(props.error))
</script>

<template>
  <div class="w-full flex flex-col justify-center items-center ">
      <div class="flex flex-col max-w-[200px]">
        <Image class="ico w-full text-[#4f108e]" :class="[':text-[#c338c3]']"/>
      </div>
      <h1 class=" text-[#5000a0] text-lg font-semibold text-center ">{{message}}</h1>
  </div>
</template>

