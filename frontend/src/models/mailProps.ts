import type { MailSummary } from "./Mails";

export interface MailsProps{
  mails:MailSummary[],
  error?:Error,
  isLoading?:boolean
}

export interface MailsEmit{
  (e:'onselectmail',mail:MailSummary): void
}
