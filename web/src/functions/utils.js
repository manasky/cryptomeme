export let formatPrice = (amount) => {
  if (amount > 1) {
    return amount.toLocaleString('en-US', { maximumFractionDigits: 2 })
  }
  return amount.toLocaleString('en-US', { maximumFractionDigits: 10 })
}

export let normalizeArrayToMin = (arr) => {
  let priceMin = Math.min(...arr)
  for (var i = 0; i < arr.length; i++) {
    arr[i] = arr[i] - priceMin
  }
  return arr
}


export let extractObjectProperyToArray = (obj, prop) => {
  return Object.keys(obj).map((key) => {
    return obj[key][prop]
  })
}