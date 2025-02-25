import type { Mail } from '@/models/Mails';
import DOMPurify from 'dompurify';

export const formatXField = (value:string) =>{
  return value.split(/<[^>]+>,?\s?/).filter(Boolean)
}


export const sanatizeBodyMail = (mail:Mail) =>{
  return DOMPurify.sanitize(mail.body,{
      FORBID_ATTR: ['href']
  })
}
