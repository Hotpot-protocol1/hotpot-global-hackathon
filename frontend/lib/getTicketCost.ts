import { ethers } from 'ethers'
import { hotpotAbi, Hotpot_CONTRACT_SEP } from '../contracts/index'

const alchemyKey = process.env.NEXT_PUBLIC_ALCHEMY_ID

const getTicketCost = async (): Promise<string | null> => {
  const provider = new ethers.providers.JsonRpcProvider(
    `https://eth-sepolia.g.alchemy.com/v2/${alchemyKey}`
  )
  const contract = new ethers.Contract(Hotpot_CONTRACT_SEP, hotpotAbi, provider)

  try {
    const ticketCost = await contract.raffleTicketCost()
    return ethers.utils.formatEther(ticketCost)
  } catch (err) {
    console.error('Error:', err)
  }

  return null
}

export default getTicketCost
