console.log("URL:",import.meta.url)
console.log("base URL:",import.meta.env.BASE_URL)
console.log("VITE URL:",import.meta.env.VITE_API_URL)
console.log("VITE URL:",import.meta.env.VITE_INVERSED_URL_API)
export const API_URL = import.meta.env.VITE_API_URL
export const INVERSED_API_URL = import.meta.env.VITE_INVERSED_URL_API
export const TIME_ZONE = Intl.DateTimeFormat().resolvedOptions().timeZone
export const URL_ENDPOINTS = {
  GetMail:`${API_URL}/mails?`,
  GetMails:`${API_URL}/mails/`
}

export const IS_DEV =  import.meta.env.PROD
export const IS_PROD =  import.meta.env.DEV
export const PORT =  import.meta.env.VITE_PORT

