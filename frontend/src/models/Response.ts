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

export class ResponseError extends Error implements IResponseError {
  status: number;
  code: string;

  constructor(status:number,code:string,message:string){
    super()
    this.status = status;
    this.code = code;
    this.message = message;
  }

}
