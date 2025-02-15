export const formatDate = (date:string | Date) => {
  try {
    if (!(date instanceof Date))
      date = new Date(date);

    return date.toLocaleDateString('es-ES', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });

  } catch (error) {
    return ""
  }
};

