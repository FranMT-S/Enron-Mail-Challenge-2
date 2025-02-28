<script setup lang="ts">
import { formatDate } from '@/helpers/formatDate';
import type { MailSummary } from '@/models/Mails';
import CalendarIconTwo from '@/modules/components/icons/CalendarIconTwo.vue';
import PageIcon from '@/modules/components/icons/PageIcon.vue';
import UserIcon from '@/modules/components/icons/UserIcon.vue';
import UsersIcon from '@/modules/components/icons/UsersIcon.vue';
import router from '@/routes';
import { NamesRouter } from '@/routes/routesNames';
import { computed } from 'vue';
import MailsTableColumn from './MailsTableColumn.vue';
import MailsTableRow from './MailsTableRow.vue';
import MailsError from './errors/MailsError.vue';
import MailLoader from './MailLoader.vue';
import {  type MailsEmit, type MailsProps } from '@/models/mailProps';
import Highlight from './Highlight.vue';

const props = defineProps<MailsProps>();

const emit = defineEmits<MailsEmit>();

const mailList = computed(() => {
  return props.mails.map(mail => {
    return {
      ...mail,
      date: formatDate(mail.date)
    }
  })
});

const onSelectMail = (index:number) =>{
  const mail = props.mails[index];
  if(mail)
    emit('onselectmail',mail)
}

</script>

<template>
<div class="overflow-x-auto h-full mb-5 px-[5px] rounded-t-[10px]">
        <table class="min-w-max w-full table-auto " :class="{'h-full':mailList.length == 0 || error || isLoading}">
            <thead class="sticky top-0 bg-deg-purple-2 text-[#fff]">
                <tr class=" uppercase text-sm leading-normal">
                    <MailsTableColumn>
                      <UserIcon style="height: 16px;"/><p>From</p>
                    </MailsTableColumn>
                    <MailsTableColumn>
                      <UsersIcon style="height: 20px;"/><p>To</p>
                    </MailsTableColumn>
                    <MailsTableColumn>
                      <PageIcon style="height: 16px;"/><p>Subject</p>
                    </MailsTableColumn>
                    <MailsTableColumn>
                      <CalendarIconTwo style="height: 16px;"/><p>Date</p>
                    </MailsTableColumn>
                </tr>
            </thead>
            <tbody class="text-gray-600  text-sm font-light text-left">
              <template v-if="isLoading">
                <tr>
                  <td colspan="4">
                      <MailLoader showText/>
                  </td>
                </tr>
              </template>
              <template v-else-if="mailList.length > 0 && !error">
                <tr v-for="(mail,index) in mailList" :key="mail.id"  @click="() => onSelectMail(index)"
                  class=" h-[34px] border-b border-gray-200 shadow-purple cursor-pointer  hover:text-[#fff] bg-deg-purple-2_hover"
                >
                    <MailsTableRow>
                      <Highlight class="font-medium" >{{ mail.from }}</Highlight>
                    </MailsTableRow>
                    <MailsTableRow>
                      <Highlight>{{ mail.to }}</Highlight>
                    </MailsTableRow>
                    <MailsTableRow>
                      <Highlight>{{ mail.subject }}</Highlight>
                    </MailsTableRow>
                    <MailsTableRow>
                      <Highlight> {{ mail.date }}</Highlight>
                    </MailsTableRow>
                  </tr>
              </template>
              <template v-else>
                <tr>
                  <td colspan="4">
                    <MailsError :error="error" :onlyTitle="false"/>
                  </td>
                </tr>
              </template>
          </tbody>
        </table>
    </div>
</template>

<style>
.shadow-purple:hover {
    box-shadow: 0px 0px 7px 0px #200576;
}

</style>
