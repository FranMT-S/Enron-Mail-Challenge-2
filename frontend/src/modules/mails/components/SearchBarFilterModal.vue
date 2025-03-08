<script setup lang="ts">
import { OffSetFromDate } from '@/helpers/formatDate';
import type { IFilterFormData } from '@/models/FormFilter';
import CrossIcon from '@/modules/components/icons/CrossIcon.vue';
import { computed, reactive, ref, watch } from 'vue';
import { useClickOutside } from '../composables/useClickOutside';
import SearchBarFilterFieldInput from './SearchBarFilterFieldInput.vue';
import SelectCustom from './SelectCustom.vue';
import Slider from './Slider.vue';

interface IEvents {
  (e: 'on-close',isActive:boolean): void
  (e: 'on-submit',data:IFilterFormData): void
}

interface Props {
  isActive:boolean
}

const {isActive} = defineProps<Props>()

const emit = defineEmits<IEvents>()
const formData = reactive<IFilterFormData>({
  from:null,
  to:null,
  subject:null,
  after:null,
  before:null,
  since:null,
  until:null,
  exclude:null,
})


const dateSinceMax = computed(() => OffSetFromDate(formData.until,0))
const dateUntilMin = computed(() => OffSetFromDate(formData.since,0))

const errormsg = ref("")
const isBetween = ref(false)
const isPickerOpen = ref(isActive)
const selectedDate = ref<Date | undefined | null>(null)
const selectedType = ref('after')
const selectOptions = ref(['after','before'])

const handleClickOutside = (event: MouseEvent) => {
  if (isPickerOpen.value && formRef.value && !formRef.value.contains(event.target as Node)) {
    close()
  }
};

const onSelectAfterOrBefore = (type:string | number) =>{
  selectedType.value = type.toString()
}

const sendData = () =>{
  if(!validateErrors())
    return

  emit('on-submit',formData)
  close();
}

const validateErrors = () =>{
  const {since,until} = formData

  if(since && until && (since > until || until < since)){
    errormsg.value = "Select a valid dates range "
    return false
  }

  if((since && !until) || (!since && until)){
    errormsg.value = "Select a valid dates range "
    return false
  }

  errormsg.value = ""
  return true
}


const close = () => {
  emit('on-close',false)
}

watch(isBetween, () =>{
  formData.after = null
  formData.before = null
  formData.since = null
  formData.until = null
  errormsg.value = ""
})

watch(selectedType, () =>{
  formData.after = null
  formData.before = null
  errormsg.value = ""
})

watch(selectedDate, (newValue) =>{
  if(selectedType.value == 'after')
    formData.after = newValue
  else
    formData.before = newValue
})


const {formRef} = useClickOutside(handleClickOutside)

</script>

<template>
  <form ref="formRef"  @submit.prevent="sendData" @keypress.enter.prevent="sendData" v-if="isActive" class="min-h-fit w-full max-w-[600px] right-0  flex flex-col justify-center top-full absolute w-100 z-[300] ">
      <div class="relative  px-4 py-6  max-h-[86vh] md:max-h-fit overflow-auto bg-white md:ms-8 ms-0 shadow rounded-lg shadow-modal-1">
        <div class="max-w-md mx-auto">
          <div class="divide-y divide-gray-200">
            <section class=" text-base leading-6 space-y-4 text-gray-700 sm:text-lg sm:leading-7">
              <SearchBarFilterFieldInput type="text" v-model="formData.from" label="From" placeholder="emails"/>
              <SearchBarFilterFieldInput type="text" v-model="formData.to" label="To" placeholder="emails"/>
              <SearchBarFilterFieldInput type="text" v-model="formData.subject" label="Subject" />
              <SearchBarFilterFieldInput type="text" v-model="formData.exclude" label="Exclude" />
              <div class="flex flex-col  gap-[4px] border-gray-200 border p-[10px] rounded-[5px] py-[12px]">
                <div class="flex flex-row row gap-[10px] mb-[8px] ">
                  <Slider v-model="isBetween"
                    :switch-style=" isBetween ? {'background-color':'#2563eb'} : undefined"
                  />
                  <span :class="{'font-semibold text-[#2049d5]': isBetween}">Between</span>
                </div>
                <div v-if="!isBetween" class="flex flex-col mobile:flex-row items-end gap-[4px] ">
                  <SelectCustom class="max-mobile:w-full max-mobile:h-full h-[35px] w-[83px] text-center"
                    :items="selectOptions"
                    @onSelect="onSelectAfterOrBefore"
                  />
                  <SearchBarFilterFieldInput
                    class="max-mobile:w-full"
                    type="date"
                    v-model="selectedDate"
                  />
                </div>
                <div v-else class="flex items-start flex-wrap gap-[4px] ">
                  <SearchBarFilterFieldInput
                    type="date"
                    v-model="formData.since"
                    label="Since"
                    :max="dateSinceMax"
                  />
                  <SearchBarFilterFieldInput
                    type="date"
                    v-model="formData.until"
                    label="Until"
                    :min="dateUntilMin"
                  />
                </div>
                <p v-if="errormsg.length > 0" class="text-red-600 text-[15px] first-letter:capitalize "> {{ errormsg }}</p>
              </div>
            </section>
            <footer class="pt-4 flex items-center gap-4 mb-[5px]">
                <button @click="close" class="min-w-fit flex shrink justify-center items-center border opacity-70 hover:opacity-100 border-red-500 text-red-500 hover:text-white hover:bg-red-500 w-full px-[2px] py-[2px]  mobile-sm:px-4 mobile-sm:py-3 rounded-md focus:outline-none"
                >
                  <CrossIcon class="hidden mobile:block"/>Cancel
                </button>
                <button type="submit"
                  class="bg-blue-500 flex hover:bg-blue-600  justify-center items-center w-full text-white px-[2px] py-[2px] mobile-sm:px-4 mobile-sm:py-3 rounded-md focus:outline-none"
                >
                  Create
                </button>
            </footer>
          </div>
        </div>
      </div>
    </form>
</template>


<style>
.flatpickr-calendar {
    top: 50px !important; /* Mueve el picker hacia abajo */
    left: 100px !important; /* Mueve el picker a la derecha */
}
</style>
