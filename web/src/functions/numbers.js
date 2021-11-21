export let formatPrice = (amount) => {
  if (amount > 1) {
    return amount.toLocaleString('en-US', { maximumFractionDigits: 2 })
  }
  return amount.toLocaleString('en-US', { maximumFractionDigits: 10 })
}