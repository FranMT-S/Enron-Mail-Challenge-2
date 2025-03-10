import type { SortField } from "@/models/Sort"

// Create a sort string used in query params
export const CreateSortString = (sortList:SortField[] = []) =>{
  if(sortList.length == 0)
    return ""

  let sortString = ""

  for(var s of sortList){
    sortString += `,${s.field}.${s.sort}`
  }
  sortString = sortString.substring(1)
  return sortString
}
