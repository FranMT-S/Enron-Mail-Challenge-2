import type { Mail } from '@/models/Mails';
import DOMPurify from 'dompurify';

export const formatXField = (value:string) =>{
  let users = value.split(/<[^>]+>,?\s?|',/)
  if(users.length == 1){
    // X_to format only names without metadata
    users = value.split(",")
  }

  return users.map(w => w.replace("'","").trim()).filter(Boolean)
}


export const sanitizeBodyMail = (mail:Mail) =>{
  return DOMPurify.sanitize(mail.body,{
      FORBID_ATTR: ['href']
  })
}
