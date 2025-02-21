import type { Mail } from "@/models/Mails";
import { ref } from "vue";
import { fetchGetMailByID, fetchGetMails } from "../services/mailservice";
import type { IResponseError, ResponseMailSummary } from "@/models/Response";


export const useMailService = () => {
  const controller = ref<AbortController | null>(null);

  const abortPreviousRequest = () =>{
    if (controller.value) {
      controller.value.abort();
    }
  }

  const getMails = async (query:string = "", page:number = 1, size:number=  30, isAbortPreviousRequest = true) => {
    if(isAbortPreviousRequest)
      abortPreviousRequest()

    const {response,controller:co} =  fetchGetMails(query,page,size);
    controller.value = co;
    const res  = await response
    if(!res.ok){
      const error: IResponseError = await res.json()
      throw error
    }

    const mails:ResponseMailSummary = await res.json()
    mails.emails = mails.emails.map(em => {
      if (!(em.date instanceof Date)){
        em.date = new Date(em.date)
      }

      return em
    })

    return mails
  };

  const getMailByID = async (id:string, isAbortPreviousRequest:boolean = true) => {
    if(isAbortPreviousRequest)
      abortPreviousRequest()

    const {response,controller:co} =  fetchGetMailByID(id);
    controller.value = co;
    const res  = await response

    if(!res.ok){
      const error: IResponseError = await res.json()
      throw error
    }

    const mail:Mail = await res.json()
    if (!(mail.date instanceof Date)){
      mail.date = new Date(mail.date)
    }

    return mail
  }

  return {
    controller,
    getMailByID,
    getMails,
    abortPreviousRequest
  }
}
