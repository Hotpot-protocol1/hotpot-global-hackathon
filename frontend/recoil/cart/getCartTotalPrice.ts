import { selector } from 'recoil'
import recoilCartTokens, { getPricingPools } from 'recoil/cart'
import getTotalPrice from 'lib/getTotalPrice'

export default selector({
  key: 'cartTotal',
  get: async ({ get }) => {
    const cartTokens = get(recoilCartTokens)

    const totalPricesPromises = cartTokens.map(async (token) => {
      const { itemId } = token
      const price = await getTotalPrice(itemId)
      return Number(price) || 0
    })

    const totalPrices = await Promise.all(totalPricesPromises)

    const cartTotal = totalPrices.reduce((total, price) => total + price, 0)
    return cartTotal
  },
})
