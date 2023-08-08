import React, {
  ReactNode,
  createContext,
  useContext,
  useEffect,
  useState,
} from 'react'
import getPrizePool, { Item } from '../lib/getPrizePool'
import getAllListedNFTs, { Item as NFT } from 'lib/getAllListedNFTs'
type HotpotProviderProps = {
  children: ReactNode
}

type HotpotContextType = {
  prizePool: Item | null
  listedNFTs: NFT[] | null
  isLoadingPrizePool: boolean
  isLoadingNFTs: boolean
}

const HotpotContext = createContext<HotpotContextType | undefined>(undefined)

export const HotpotProvider: React.FC<HotpotProviderProps> = ({ children }) => {
  const [prizePool, setPrizePool] = useState<Item | null>(null)
  const [isLoadingPrizePool, setIsLoadingPrizePool] = useState(true)
  const [listedNFTs, setListedNFTs] = useState<NFT[] | null>(null)
  const [isLoadingNFTs, setIsLoadingNFTs] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      try {
        const prizePool = await getPrizePool()
        const nfts = await getAllListedNFTs()
        setPrizePool(prizePool)
        setListedNFTs(nfts)
      } catch (error) {
        console.error('An error occurred while fetching data:', error)
      }
      setIsLoadingPrizePool(false)
      setIsLoadingNFTs(false)
    }

    fetchData()
  }, [])

  return (
    <HotpotContext.Provider
      value={{ prizePool, listedNFTs, isLoadingPrizePool, isLoadingNFTs }}
    >
      {children}
    </HotpotContext.Provider>
  )
}
export const useHotpotContext = () => {
  const context = useContext(HotpotContext)
  if (context === undefined) {
    throw new Error('useHotpotContext must be used within a HotpotProvider')
  }
  return context
}
