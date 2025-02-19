import { API_URL } from "@/constants/varEnviroment";
import { fetchCustom } from "@/helpers/fetch";


console.log(API_URL);
const baseURl = new URL(API_URL + "/")

export function fetchGetMails(query:string = "", page:number = 1, size:number=  30){
  // Cancela la petición anterior si existe
  const controller = new AbortController()

  const url = new URL("mails",baseURl)
  url.searchParams.set('query',query)
  url.searchParams.set('page',page.toString())
  url.searchParams.set('size',size.toString())

  const response = fetchCustom(url, {
    signal:controller.signal,
    method:"Get",
  });

  return {response,controller}
};

export function fetchGetMailByID(id:string){
  const controller = new AbortController()
  const url = new URL(`mails/${id}`, baseURl);

  const response = fetchCustom(url, {
    signal:controller.signal,
    method:"Get",
  });

  return {response,controller}
};
