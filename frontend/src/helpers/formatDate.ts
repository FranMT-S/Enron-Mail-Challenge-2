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

    let month = (date.getMonth() + 1).toString().padStart(2,'0');
    let day = '' + date.getDate().toString().padStart(2,'0');
    let year = date.getFullYear();

  return `${year}-${month}-${day}`;

  } catch (error) {
    return ""
  }
};

