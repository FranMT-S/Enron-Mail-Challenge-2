export const API_URL = import.meta.env.VITE_API_URL
export const TIME_ZONE = Intl.DateTimeFormat().resolvedOptions().timeZone
export const URL_ENDPOINTS = {
  GetMail:`${API_URL}/mails?`,
  GetMails:`${API_URL}/mails/`
}
