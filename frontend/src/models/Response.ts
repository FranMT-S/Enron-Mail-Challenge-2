import type { MailSummary } from "./Mails";

export interface ResponseMailSummary{
  total:number,
  emails:MailSummary[]
}
