import { ethers } from 'ethers'
import { hotpotAbi, Hotpot_CONTRACT_SEP } from '../contracts/index'

const alchemyKey = process.env.NEXT_PUBLIC_ALCHEMY_ID

export type Item = {
  currentPotSize: string
  potLimit: string
}
const provider = new ethers.providers.JsonRpcProvider(
  `https://eth-goerli.g.alchemy.com/v2/${alchemyKey}`
)
const contract = new ethers.Contract(Hotpot_CONTRACT_SEP, hotpotAbi, provider)

const getPrizePool = async (): Promise<Item | null> => {
  try {
    const [currentPotSizeWei, potLimitWei] = await Promise.all([
      contract.currentPotSize(),
      contract.potLimit(),
    ])

    const currentPotSize = ethers.utils.formatEther(currentPotSizeWei)
    const potLimit = ethers.utils.formatEther(potLimitWei)

    const prizePool: Item = {
      currentPotSize,
      potLimit,
    }
    console.log('prizepool:', prizePool)
    return prizePool
  } catch (err) {
    console.error('Error:', err)
    return null
  }
}

export default getPrizePool
