// PrizePoolContext.tsx
import React, {
  ReactNode,
  createContext,
  useContext,
  useEffect,
  useState,
} from 'react'
import getPrizePool, { Item } from '../lib/getPrizePool'
type PrizePoolProviderProps = {
  children: ReactNode
}

type PrizePoolContextType = {
  prizePoolData: Item | null
  isLoading: boolean
}

const PrizePoolContext = createContext<PrizePoolContextType | undefined>(
  undefined
)

export const PrizePoolProvider: React.FC<PrizePoolProviderProps> = ({
  children,
}) => {
  const [prizePoolData, setPrizePoolData] = useState<Item | null>(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getPrizePool()
        if (data) {
          setPrizePoolData(data)
        }
      } catch (error) {
        console.error(
          'An error occurred while fetching the prize pool data:',
          error
        )
      }
      setIsLoading(false)
    }

    fetchData()
  }, [])

  return (
    <PrizePoolContext.Provider value={{ prizePoolData, isLoading }}>
      {children}
    </PrizePoolContext.Provider>
  )
}

export const usePrizePoolContext = () => {
  const context = useContext(PrizePoolContext)
  if (context === undefined) {
    throw new Error(
      'usePrizePoolContext must be used within a PrizePoolProvider'
    )
  }
  return context
}
