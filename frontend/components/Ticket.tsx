import { FC } from 'react'
import { useAccount } from 'wagmi'
import { CgSpinner } from 'react-icons/cg'
import useSWR from 'swr'
import { getLatestPot } from 'lib/getLatestPot'
import Link from 'next/link'

const Ticket: FC = () => {
  const account = useAccount()
  const { address } = useAccount()
  const { data, error } = useSWR(address ? ['latestPot', address] : null, () =>
    address ? getLatestPot(address) : null
  )

  if (!address) {
    return null
  }

  if (account.isConnected) {
    if (!data && !error) {
      return (
        <div className="flex items-center rounded-lg border border-solid border-[#CFD8E1] px-6 py-2">
          {' '}
          <CgSpinner className="mr-2 animate-spin" />
          <div className="text-base font-normal">TIX</div>
        </div>
      )
    }

    if (data?.currentPot) {
      return (
        <>
          <Link href="/rewards" legacyBehavior={true}>
            <div className="flex items-center rounded-lg border border-solid border-[#CFD8E1] px-6 py-2">
              <div className="text-base font-normal">
                {' '}
                <span className="text-purple-800">
                  {data.currentPot.NumOfTickets}
                </span>{' '}
                TIX
              </div>
            </div>
          </Link>
        </>
      )
    }
  }

  return null
}

export default Ticket
