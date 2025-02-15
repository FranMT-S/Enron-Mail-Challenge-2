import { TIME_ZONE } from "@/constants/varEnviroment";

export const fetchCustom = (input: string | URL | globalThis.Request, init?: RequestInit) => {
  const headers = new Headers({
    ...init?.headers,
  });


  headers.set('Time-Zone',TIME_ZONE)

  return fetch(input,{...init,headers})
};
