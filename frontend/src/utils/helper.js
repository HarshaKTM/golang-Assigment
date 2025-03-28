/**
 * Debounce function to limit how often a function is called
 * @param {Function} func - The function to debounce
 * @param {number} wait - The debounce wait time in milliseconds
 * @returns {Function} - The debounced function
 */
export const debounce = (func, wait) => {
  let timeout;
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout);
      func(...args);
    };
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
  };
};

/**
 * Format a price value to currency display
 * @param {number} price - The price value
 * @param {string} currency - The currency symbol
 * @returns {string} - Formatted price string
 */
export const formatPrice = (price, currency = '$') => {
  if (isNaN(price)) return `${currency}0.00`;
  return `${currency}${parseFloat(price).toFixed(2)}`;
};

/**
 * Truncate a string if it exceeds a specified length
 * @param {string} str - The string to truncate
 * @param {number} maxLength - Maximum length before truncating
 * @returns {string} - Truncated string with ellipsis
 */
export const truncateString = (str, maxLength) => {
  if (!str) return '';
  if (str.length <= maxLength) return str;
  return `${str.substring(0, maxLength)}...`;
};

/**
 * Format a date string to a more readable format
 * @param {string} dateString - Date string in ISO format (YYYY-MM-DD)
 * @returns {string} - Formatted date string
 */
export const formatDate = (dateString) => {
  if (!dateString) return '';
  
  const options = { year: 'numeric', month: 'long', day: 'numeric' };
  try {
    const [year, month, day] = dateString.split('-');
    const date = new Date(year, month - 1, day);
    return date.toLocaleDateString('en-US', options);
  } catch (error) {
    return dateString;
  }
}; 