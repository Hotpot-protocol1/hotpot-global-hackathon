import { FC, useState, useEffect } from 'react'
import ConnectWalletButton from './ConnectWalletButton'
import { useAccount } from 'wagmi'
import { getLatestPot, PotData, Pots } from 'lib/getLatestPot'
import { CgSpinner } from 'react-icons/cg'
import { getPotById } from 'lib/getPotById'
import { Item } from '../lib/getPrizePool'
import useMounted from 'hooks/useMounted'
import ResultsModal from './modal/ResultsModal'
import { useHotpotContext } from 'context/HotpotContext'

interface Ticket {
  number: number
  date: Date
}

interface TicketsGridProps {
  prizePool?: Item | null
}

const TicketsGrid: FC<TicketsGridProps> = () => {
  const [tab, setTab] = useState<string>('current')
  const [dummyData, setDummyData] = useState<Ticket[]>([])
  const [data, setData] = useState<PotData | null>(null)
  const [previousData, setPreviousData] = useState<PotData | null>(null)
  const { prizePool, isLoadingPrizePool } = useHotpotContext()
  const currentPotSize = parseFloat(prizePool?.currentPotSize ?? '0')
  const potLimit = parseFloat(prizePool?.potLimit ?? '0')
  const potFill = Math.round((currentPotSize / potLimit) * 100)
  const [claimModalOpen, setClaimModalOpen] = useState(false)
  const [tabs, setTabs] = useState<Pots[] | null>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<Error | null>(null)
  const account = useAccount()
  const { address } = useAccount()
  const isMounted = useMounted()

  useEffect(() => {
    if (account?.isDisconnected) {
      setDummyData(
        Array.from({ length: 50 }, (_, i) => ({
          number: Math.floor(Math.random() * 90000 + 10000),
          date: new Date(Date.now() - (i % 200) * 24 * 60 * 60 * 1000),
        }))
      )
    }
    const fetchLatestPotData = async () => {
      if (address) {
        const { currentPot, pots } = await getLatestPot(address)
        if (pots) {
          console.log(pots)
          const potsWithRaffle = pots.filter(
            (pot) => pot.raffle_timestamp !== null
          )
          const lastTwoPots = potsWithRaffle.slice(-2)
          setData(currentPot)

          setTabs(lastTwoPots)
        }
        setData(currentPot)
      }
    }

    fetchLatestPotData()
  }, [address])

  const handleTabClick = async (potId: number) => {
    setLoading(true)
    setError(null)

    if (potId && address) {
      const potData = await getPotById(address, potId)
      if (potData) {
        setPreviousData(potData)
        setLoading(false)
      } else {
        setError(new Error('Failed to fetch pot data.'))
        setLoading(false)
      }
    }
  }

  const filteredTickets =
    tab === 'current' ? data?.tickets : previousData?.tickets

  return (
    <>
      {account.isDisconnected || !isMounted || !address ? (
        <div className="relative">
          <div className="absolute inset-x-0 top-1/3 z-20 mx-auto flex h-[150px] w-[380px] items-center justify-center rounded bg-slate-100">
            <div className="flex flex-col items-center justify-center gap-4 px-4">
              <div className="reservoir-subtitle text-center">
                Connect your wallet to see your tickets
              </div>
              <ConnectWalletButton className="w-full px-10">
                <span>Connect Wallet</span>
              </ConnectWalletButton>
            </div>
          </div>
          <div className="relative z-0 rounded-md bg-opacity-50 p-2 blur-lg">
            <div className="my-4 mx-4 flex border-b border-b-[#E1D8FD] text-sm font-normal text-[#98A2B3]">
              <button className="${'border-b-2 border-[#6A3CF5] p-4 font-medium  text-[#6A3CF5]">
                Current
              </button>
              <button className="p-4 font-medium">7 July 2023</button>
              <button className="p-4 font-medium">7 June 2023</button>
            </div>

            <h2 className="m-4 font-medium text-[#FF60D5]">My Tickets: 1445</h2>
            <div className="grid-rows-10 md:grid-rows-19 m-4 grid max-h-[500px] grid-cols-5 gap-4 overflow-auto md:grid-cols-10">
              <div className="rounded-lg border border-dashed border-black">
                <div
                  className="rounded-l-lg bg-[#9270FF] py-2 text-center text-sm text-[#FAF9FE]"
                  style={{ width: '80%' }}
                >
                  <div>80%</div>
                </div>
              </div>
              {dummyData.map((ticket, i) => (
                <div
                  key={i}
                  className="rounded-lg border bg-[#9270FF] py-2 text-center text-sm text-[#FAF9FE]"
                >
                  <div>#{ticket.number}</div>
                </div>
              ))}
            </div>
          </div>
        </div>
      ) : (
        <div className="w-full">
          <div className="my-4 mx-4 flex border-b border-b-[#E1D8FD] text-sm font-normal text-[#98A2B3]">
            {}
            <button
              className={`p-4 font-medium ${
                tab === 'current'
                  ? 'border-b-2 border-[#6A3CF5]  text-[#6A3CF5] '
                  : ''
              }`}
              onClick={() => setTab('current')}
            >
              Current
            </button>
            {tabs?.map((pot, index) => (
              <button
                className={`p-4 font-medium ${
                  tab === `tab${index}`
                    ? 'border-b-2 border-[#6A3CF5] text-[#6A3CF5]'
                    : ''
                }`}
                onClick={() => {
                  setTab(`tab${index}`)
                  handleTabClick(pot.pot_id)
                }}
              >
                {new Date(pot.raffle_timestamp!).toLocaleDateString(undefined, {
                  day: 'numeric',
                  month: 'long',
                  year: 'numeric',
                })}
              </button>
            ))}
          </div>

          <h2 className="m-4 font-medium text-[#FF60D5]">
            My Tickets: {data?.NumOfTickets}
          </h2>
          {data?.NumOfTickets === 0 ? (
            <div className="grid-rows-10 md:grid-rows-19 m-4 grid max-h-[500px] grid-cols-5 gap-4 overflow-auto md:grid-cols-10">
              <div>No tickets</div>
            </div>
          ) : (
            <div className="grid-rows-10 md:grid-rows-19 m-4 grid max-h-[500px] grid-cols-5 gap-4 overflow-auto md:grid-cols-10">
              {tab === 'current' && (
                <div className="flex rounded-lg border border-dashed border-black">
                  <div
                    className={`flex items-center justify-center text-center text-xs text-[#FAF9FE] ${
                      potFill === 100 ? 'rounded-lg' : 'rounded-l-lg'
                    } bg-[#9270FF]`}
                    style={{ width: `${potFill}%` }}
                  >
                    {potFill >= 25 ? `${potFill}%` : ''}
                  </div>
                </div>
              )}

              {filteredTickets?.map((ticket, i) => (
                <div
                  key={i}
                  className={`rounded-lg border py-2 text-center text-sm ${
                    tab === 'current' || ticket.is_winner
                      ? 'cursor-default bg-[#9270FF] text-[#FAF9FE]'
                      : 'cursor-pointer bg-[#E1D8FD] text-[#A58AF9]'
                  }`}
                  style={{ cursor: ticket.is_winner ? 'pointer' : 'default' }}
                >
                  {tab !== 'current' && ticket.is_winner ? (
                    <>
                      <ResultsModal trigger={<div>#{ticket.ticket_id}</div>} />
                    </>
                  ) : (
                    <div>#{ticket.ticket_id}</div>
                  )}
                </div>
              ))}
            </div>
          )}
        </div>
      )}
    </>
  )
}

export default TicketsGrid
