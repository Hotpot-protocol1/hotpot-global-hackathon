import { FC } from 'react'
import { useAccount, useConnect, useDisconnect } from 'wagmi'
import useMounted from 'hooks/useMounted'

const DARK_MODE = process.env.NEXT_PUBLIC_DARK_MODE

const Ticket: FC = () => {
  const account = useAccount()
  const { connectors } = useConnect()
  const { disconnect } = useDisconnect()
  const wallet = connectors[0]
  const isMounted = useMounted()

  if (!isMounted) {
    return null
  }

  if (account.isConnected) {
    return (
      <div className=" flex items-center rounded-lg border border-solid border-[#CFD8E1] px-6 py-2">
        <div className="text-base font-normal">
          {' '}
          <span className="text-purple-800">0</span> TIX
        </div>
      </div>
    )
  }

  return null
}

export default Ticket
