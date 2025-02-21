import type { MailSummary } from "./Mails";

export interface ResponseMailSummary{
  total:number,
  emails:MailSummary[]
}


export interface IResponseError{
  status: number,
  code: string,
  message: string
}
