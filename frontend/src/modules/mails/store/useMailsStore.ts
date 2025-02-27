import { defineStore } from "pinia";
import { useMailService } from "../composables/useMailService";
import type { ResponseMailSummary } from "@/models/Response";
import type { Mail, MailSummary } from "@/models/Mails";
import { reactive, ref } from "vue";

export const useMailsStore = defineStore('mails',() =>{

  const {getMails,getMailByID} = useMailService()
  const emailsSummary = ref<MailSummary[]>([])
  const emailDict = reactive(new Map<string,Mail>)
  const emailSummarySelected = ref<MailSummary | undefined>(undefined)
  const total = ref(0)

  const fillEmailsSummary = async (query:string = "", page:number = 1, size:number=  30) => {
    try {
      emailsSummary.value = []
      const mails:ResponseMailSummary = await getMails(query,page,size)
      total.value = mails.total
      emailsSummary.value = mails.emails
    } catch (err) {
      resetEmailsSummary()
      throw err
    }
  };

  const getEmail = async (id:string) =>{
    if(emailDict.has(id))
      return emailDict.get(id)!

    const email = await getMailByID(id)
    emailDict.set(id,email)
    return email
  }

  const resetEmailsSummary = () =>{
    emailsSummary.value = []
  }

  const resetEmails = () =>{
    emailDict.clear()
  }

  const addEmail = (email:Mail) =>{
    if(email.id.length > 0){
      emailDict.set(email.id,email)
      return true
    }

    return false
  }

  const resetEmailSummarySelected = () => {
    emailSummarySelected.value = undefined
  }

  const resetAll = () => {
    resetEmails()
    resetEmailSummarySelected()
    resetEmailsSummary
  }

  return {
    fillEmailsSummary,
    resetEmailsSummary,
    getEmail,
    resetEmails,
    resetEmailSummarySelected,
    resetAll,
    addEmail,

    emailSummarySelected,
    emailsSummary,
    emailDict,
    total,

  }
})
