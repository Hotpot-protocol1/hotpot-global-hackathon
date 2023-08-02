const useTix = (amount: number | string): number => {
  const parsedAmount = typeof amount === 'string' ? parseFloat(amount) : amount

  if (isNaN(parsedAmount)) {
    console.warn('Invalid amount provided to useTicketCalculator')
    return 0
  }

  return Math.floor(parsedAmount / 0.1)
}

export default useTix
