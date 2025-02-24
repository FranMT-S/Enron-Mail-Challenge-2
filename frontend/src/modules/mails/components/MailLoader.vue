<script setup lang="ts">
  interface Props{
    showText:boolean
  }

  const {showText} = withDefaults(defineProps<Props>(), {
    showText: true,
  });

</script>

<template>
  <div class="flex justify-center items-center flex-col">
    <span class="loader"></span>
    <h1 v-if="showText" class="w-full text-center pt-3 font-bold text-md flex justify-center item-center  ">
      Loading<span class="loading-text w-3 text-left"></span>
    </h1>
  </div>
</template>

<style>
 .loader {
  position: relative;
  border-style: solid;
  box-sizing: border-box;
  border-width: 40px 60px 30px 60px;
  border-color: #3760C9 #96DDFC #96DDFC #36BBF7;
  animation: envFloating 1s ease-in infinite alternate;
}

.loader:after{
  content:"";
  position: absolute;
  right: 62px;
  top: -40px;
  height: 70px;
  width: 50px;
  background-image:
            linear-gradient(#3b73ff 45px, transparent 0),
            linear-gradient(#3b73ff 45px, transparent 0),
            linear-gradient(#3b73ff 45px, transparent 0);
  background-repeat: no-repeat;
  background-size: 30px 4px;
  background-position: 0px 11px , 8px 35px, 0px 60px;
  animation: envDropping 0.75s linear infinite;
}

@keyframes envFloating {
  0% { transform: translate(-2px, -5px)}
  100% { transform: translate(0, 5px)}
}

@keyframes envDropping {
  0% {background-position: 100px 11px , 115px 35px, 105px 60px; opacity: 1;}
  50% {background-position: 0px 11px , 20px 35px, 5px 60px; }
  60% {background-position: -30px 11px , 0px 35px, -10px 60px; }
  75%, 100% {background-position: -30px 11px , -30px 35px, -30px 60px; opacity: 0;}
}

</style>
