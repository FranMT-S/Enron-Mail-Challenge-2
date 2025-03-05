export type SortType = "asc" | "desc"

export interface SortField{
  field:string,
  sort:SortType
}


export interface SortByItemList{
  field:string,
  label:string,
  ascLabel:string
  descLabel:string
}

export interface SortByItem{
  field:string,
  label:string,
  sortLabel:string
  sort:SortType
}
