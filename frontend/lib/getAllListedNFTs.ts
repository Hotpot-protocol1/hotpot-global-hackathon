import { ethers } from 'ethers'
import { abi, NFTMarketplace_CONTRACT_SEP } from '../contracts/index'

const alchemyKey = process.env.NEXT_PUBLIC_ALCHEMY_ID

export type Item = {
  itemId: number
  nft: string
  tokenId: string
  price: string
  seller: string
  sold: boolean
}

const getAllListedNFTs = async (): Promise<Item[] | null> => {
  const provider = new ethers.providers.JsonRpcProvider(
    `https://eth-goerli.g.alchemy.com/v2/${alchemyKey}`
  )
  const contract = new ethers.Contract(
    NFTMarketplace_CONTRACT_SEP,
    abi,
    provider
  )

  try {
    const rawNFTs = await contract.getAllListedNfts()
    const listedNFTs: Item[] = rawNFTs.map((rawNFT: any) => {
      const [itemId, nft, tokenId, price, seller, sold] = rawNFT
      return {
        itemId: itemId.toNumber(),
        nft,
        tokenId: tokenId.toString(),
        price: ethers.utils.formatEther(price),
        seller,
        sold,
      }
    })
    console.log('ListedNFTs:', listedNFTs)
    return listedNFTs
  } catch (err) {
    console.error('Error:', err)
    return null
  }
}

export default getAllListedNFTs
