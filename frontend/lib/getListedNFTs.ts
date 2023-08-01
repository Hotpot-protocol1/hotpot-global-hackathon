import { useEffect, useState } from 'react'
import { abi, NFTMarketplace_CONTRACT_SEP } from '../contracts/index'
import { ethers } from 'ethers'

const alchemyKey = process.env.NEXT_PUBLIC_ALCHEMY_ID

type Item = {
  itemId: number
  nft: string
  tokenId: string
  price: string
  seller: string
  sold: boolean
}

const getListedNFTs = () => {
  const [listedNFTs, setListedNfts] = useState<Item[] | null>(null)
  const [loading, setLoading] = useState(true)
  const provider = new ethers.providers.JsonRpcProvider(
    `https://eth-sepolia.g.alchemy.com/v2/${alchemyKey}`
  )
  const contract = new ethers.Contract(
    NFTMarketplace_CONTRACT_SEP,
    abi,
    provider
  )

  useEffect(() => {
    const fetchNFTs = async () => {
      try {
        const rawNFTs = await contract.getAllListedNfts()
        const parsedNFTs: Item[] = rawNFTs.map((rawNFT: any) => {
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
        console.log(parsedNFTs)
        setListedNfts(parsedNFTs)
        setLoading(false)
      } catch (err) {
        console.error('Error:', err)
        setLoading(false)
      }
    }

    fetchNFTs()
  }, [])

  return { listedNFTs }
}

export default getListedNFTs
