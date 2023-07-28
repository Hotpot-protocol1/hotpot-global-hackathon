import { useEffect, useState } from 'react'
import { abi, NFTMarketplace_CONTRACT_SEP } from '../contracts/index'
import { ethers } from 'ethers'

const alchemyKey = process.env.NEXT_PUBLIC_ALCHEMY_ID

const getTotalPrice = (itemId: number): string | null => {
  const [totalPrice, setTotalPrice] = useState<string | null>(null) // Updated type
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchTotalPrice = async () => {
      const provider = new ethers.providers.JsonRpcProvider(
        `https://eth-sepolia.g.alchemy.com/v2/${alchemyKey}`
      )
      const contract = new ethers.Contract(
        NFTMarketplace_CONTRACT_SEP,
        abi,
        provider
      )

      try {
        if (itemId) {
          const price = await contract.getTotalPrice(itemId)
          setTotalPrice(ethers.utils.formatEther(price))
          setLoading(false)
          console.log('Total Price', price)
        }
      } catch (err) {
        console.error('Error:', err)
      }
    }

    fetchTotalPrice()
  }, [itemId])

  return totalPrice
}

export default getTotalPrice
