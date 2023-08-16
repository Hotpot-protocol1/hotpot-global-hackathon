import { getRafflePot } from 'lib/getRafflePot'
import { useEffect, useState } from 'react'
import { useAccount } from 'wagmi'

import ResultsModal from './modal/ResultsModal'

const PotResultBanner = () => {
  const { isConnected, address } = useAccount()
  const [isRaffleDrawn, setIsRaffleDrawn] = useState(false)

  useEffect(() => {
    if (isConnected && address) {
      getRafflePot(address).then((res) => {
        if (!res) {
          setIsRaffleDrawn(true)
        }
      })
    }
  }, [isConnected, address])

  if (isRaffleDrawn) {
    return (
      <div className="col-span-full">
        <div className="flex w-full flex-row items-center justify-center gap-1 bg-[#FFD43C] py-2">
          <h2 className="text-sm font-normal">
            Pot O’ Gold Raffle has Drawn,{' '}
          </h2>
          <ResultsModal
            trigger={
              <a className="cursor-pointer text-sm font-semibold text-[#6A3CF5] hover:text-purple-500">
                {' '}
                CHECK RESULTS!
              </a>
            }
          />
        </div>
      </div>
    )
  }
  return null
}

export default PotResultBanner
