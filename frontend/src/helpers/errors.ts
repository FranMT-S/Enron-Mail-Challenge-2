import { ResponseError } from "@/models/Response"

export const isAbortError = (err:Error) =>{
  return err instanceof DOMException && err.name === "AbortError"
}

export const ValidateError = (err:unknown):ResponseError | undefined =>{
  let error:ResponseError | undefined = undefined
  if(err instanceof Error){
    if (isAbortError(err))
      error = undefined

   if(err instanceof ResponseError)
    error = err
  }else{
    console.log(err)
    err = new ResponseError(500,'InternalServerError','There was a problem try again later.')
  }

  return error
}
