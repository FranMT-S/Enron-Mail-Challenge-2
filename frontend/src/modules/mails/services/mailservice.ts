import { API_URL } from "@/constants/varEnviroment";
import { fetchCustom } from "@/helpers/fetch";
import { sanitizeInput } from "@/helpers/regex";
import { CreateSortString } from "@/helpers/sort";
import type { SortField } from "@/models/Sort";

const baseURL = new URL(API_URL + "/")

export function fetchGetMails(query:string = "", page:number = 1, size:number=  30,sortList:SortField[] = []){

  query = sanitizeInput(query);

  const controller = new AbortController()
  const url = new URL("mails",baseURL)
  url.searchParams.set('query',query)
  url.searchParams.set('page',page.toString())
  url.searchParams.set('size',size.toString())

  if(sortList.length > 0){
    let sortString = CreateSortString(sortList)
    url.searchParams.set('sort',sortString)
  }

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
