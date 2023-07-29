import { ethers } from 'ethers'
import { hotpotAbi, Hotpot_CONTRACT_SEP } from '../contracts/index'

const alchemyKey = process.env.NEXT_PUBLIC_ALCHEMY_ID

export type Item = {
  currentPotSize?: string
  potLimit?: string
}

const getPrizePool = async (): Promise<Item | null> => {
  const provider = new ethers.providers.JsonRpcProvider(
    `https://eth-sepolia.g.alchemy.com/v2/${alchemyKey}`
  )
  const contract = new ethers.Contract(Hotpot_CONTRACT_SEP, hotpotAbi, provider)

  try {
    const currentPotSizeWei = await contract.currentPotSize()
    const potLimitWei = await contract.potLimit()

    const currentPotSize = ethers.utils.formatEther(currentPotSizeWei)
    const potLimit = ethers.utils.formatEther(potLimitWei)

    const formattedPrizePool: Item = {
      currentPotSize,
      potLimit,
    }

    return formattedPrizePool
  } catch (err) {
    console.error('Error:', err)
    return null
  }
}

export default getPrizePool