export function formatDateTime(date: string): string {
  const d = new Date(date)
  return new Intl.DateTimeFormat(undefined, {
    dateStyle: 'long',
    timeStyle: 'short',
    hour12: false,
  }).format(d)
}
