import React, { FC, useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { CgSpinner } from 'react-icons/cg'
import { getLatestPot } from 'lib/getLatestPot'
import { setToast } from './token/setToast'
import Link from 'next/link'
import useSWR from 'swr'

const Ticket: FC = () => {
  const account = useAccount()
  const { address } = useAccount()
  const { data, error } = useSWR(address ? ['latestPot', address] : null, () =>
    address ? getLatestPot(address) : null
  )

  const [prevTicketCount, setPrevTicketCount] = useState<number | null>(null)

  useEffect(() => {
    if (data?.currentPot) {
      if (
        prevTicketCount !== null &&
        data.currentPot.NumOfTickets > 0 &&
        data.currentPot.NumOfTickets !== prevTicketCount
      ) {
        setToast({
          kind: 'tickets',
          message: '',
          title: 'You have earned Golden Tickets!',
        })
      }

      setPrevTicketCount(data.currentPot.NumOfTickets)
    }
  }, [data])

  if (!address) {
    return null
  }

  if (account.isConnected) {
    if (!data && !error) {
      return (
        <div className="flex items-center rounded-lg border border-solid border-[#CFD8E1] px-6 py-2">
          <CgSpinner className="mr-2 animate-spin" />
          <div className="text-base font-normal">TIX</div>
        </div>
      )
    }

    if (data?.currentPot) {
      return (
        <>
          <Link href="/rewards" legacyBehavior={true}>
            <div className="btn-primary-outline flex cursor-pointer items-center rounded-lg border border-solid border-[#CFD8E1] px-6 py-2">
              <div className="text-base font-normal">
                <span className="mr-1 text-purple-800">
                  {data.currentPot.NumOfTickets}
                </span>
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
