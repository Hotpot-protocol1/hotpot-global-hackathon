import { FC, useEffect, useState } from 'react'
import { useAccount, useConnect, useDisconnect } from 'wagmi'
import useMounted from 'hooks/useMounted'
import { CgSpinner } from 'react-icons/cg'
import { PotData, getLatestPot } from 'lib/getLatestPot'

const Ticket: FC = () => {
  const account = useAccount()
  const { address } = useAccount()
  const { connectors } = useConnect()
  const { disconnect } = useDisconnect()
  const wallet = connectors[0]
  const isMounted = useMounted()
  const [data, setData] = useState<PotData | null>(null)

  useEffect(() => {
    const fetchLatestPotData = async () => {
      if (address) {
        const { currentPot } = await getLatestPot(address)
        setData(currentPot)
      }
    }

    if (account.isConnected && address) {
      fetchLatestPotData()
    }
  }, [address])

  if (!isMounted || !address) {
    return null
  }

  if (account.isConnected) {
    if (!data && account?.isConnected) {
      return (
        <div>
          <CgSpinner className="animate-spin" />
        </div>
      )
    }
    return (
      <div className="flex items-center rounded-lg border border-solid border-[#CFD8E1] px-6 py-2">
        <div className="text-base font-normal">
          {' '}
          <span className="text-purple-800">{data?.NumOfTickets}</span> TIX
        </div>
      </div>
    )
  }

  return null
}

export default Ticket
