import { API_URL } from "@/constants/varEnviroment";
import { fetchCustom } from "@/helpers/fetch";
import { sanitizeInput } from "@/helpers/regex";

const baseURL = new URL(API_URL + "/")

export function fetchGetMails(query:string = "", page:number = 1, size:number=  30){

  query = sanitizeInput(query);

  const controller = new AbortController()
  const url = new URL("mails",baseURL)
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
  const url = new URL(`mails/${id}`, baseURL);

  const response = fetchCustom(url, {
    signal:controller.signal,
    method:"Get",
  });

  return {response,controller}
};
