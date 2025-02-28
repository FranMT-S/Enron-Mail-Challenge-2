import '@mdi/font/css/materialdesignicons.css';
import { createVuetify } from 'vuetify'
import { VCol, VContainer, VDatePicker, VRow } from 'vuetify/components'
import { VDateInput } from 'vuetify/labs/components'


export const vuetify =createVuetify({
  components:{
    VDatePicker,
    VDateInput,
    VRow,
    VCol,
    VContainer,
  },
  icons:{
    defaultSet:'mdi'
  }
})
