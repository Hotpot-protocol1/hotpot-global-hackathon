import ResultsModal from './modal/ResultsModal'

import { useAccount } from 'wagmi'

const PotResultBanner = () => {
  const { isConnected } = useAccount()
  if (isConnected) {
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
