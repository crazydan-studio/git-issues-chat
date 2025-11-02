// Utility functions

/**
 * Format epoch milliseconds to a readable date string
 * @param millis - Epoch time in milliseconds
 * @param pattern - Date format pattern (e.g., 'yyyy-MM-dd HH:mm:ss.S', 'yyyy-MM-dd', 'HH:mm:ss')
 * @returns Formatted date string
 */
export function formatEpocMillis(millis: number, pattern: string): string {
  const date = new Date(millis);
  
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');
  const milliseconds = String(date.getMilliseconds()).padStart(3, '0');
  
  return pattern
    .replace('yyyy', String(year))
    .replace('MM', month)
    .replace('dd', day)
    .replace('HH', hours)
    .replace('mm', minutes)
    .replace('ss', seconds)
    .replace('S', milliseconds);
}