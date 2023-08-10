import { ethers } from 'ethers'
import { abi, NFTMarketplace_CONTRACT_SEP } from '../contracts/index'

const alchemyKey = process.env.NEXT_PUBLIC_ALCHEMY_ID

const getTotalPrice = async (itemId: number): Promise<string | null> => {
  const provider = new ethers.providers.JsonRpcProvider(
    `https://eth-goerli.g.alchemy.com/v2/${alchemyKey}`
  )
  const contract = new ethers.Contract(
    NFTMarketplace_CONTRACT_SEP,
    abi,
    provider
  )

  try {
    if (itemId) {
      const price = await contract.getTotalPrice(itemId)
      return ethers.utils.formatEther(price)
    }
  } catch (err) {
    console.error('Error:', err)
  }

  return null
}

export default getTotalPrice
