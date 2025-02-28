const NameFieldWithSpecialCharacters = `[-*]?\\w+`
const valueInnerParentheses = `\\([^)]+\\)`

export const fieldRegex = new RegExp(`${NameFieldWithSpecialCharacters}:(${valueInnerParentheses}|\\S+)`)
export const specialCharactersRegex = /[<>'";&|$\\[\\]{}#%=]/

export const sanitizeInput = (query:string) =>{
  return query.replace(specialCharactersRegex,"")
}
