export const formatDate = (date:string | Date) => {
  try {
    if (!(date instanceof Date))
      date = new Date(date);

    return date.toLocaleDateString(Intl.NumberFormat().resolvedOptions().locale, {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });

  } catch (error) {
    return ""
  }
};


export const formatDateYYYYMMDD = (date:string | Date) => {
  try {
    if (!(date instanceof Date))
      date = new Date(date);

    let month = (date.getUTCMonth() + 1).toString().padStart(2,'0');
    let day = '' + date.getUTCDate().toString().padStart(2,'0');
    let year = date.getUTCFullYear();

  return `${year}-${month}-${day}`;

  } catch (error) {
    return ""
  }
};

/**  Returns a date string in format YYYY-MM-DD with  offset in days. */
export const OffSetFromDate = (date:Date | undefined | null, daysOffset = 0) =>{
  if(!date)
    return undefined

  const offset = new Date(date)
  offset.setUTCDate(offset.getUTCDate() + daysOffset)

  return formatDateYYYYMMDD(offset)
}
