import { FC, useState, useEffect } from 'react'

interface Ticket {
  number: number
  date: Date
}

interface TicketsGridProps {
  address?: string
}

const TicketsGrid: FC<TicketsGridProps> = ({ address }) => {
  const [tab, setTab] = useState<string>('current')

  const [tickets, setTickets] = useState<Ticket[]>([])

  useEffect(() => {
    setTickets(
      Array.from({ length: 1551 }, (_, i) => ({
        number: Math.floor(Math.random() * 90000 + 10000),
        date: new Date(Date.now() - (i % 200) * 24 * 60 * 60 * 1000),
      }))
    )
  }, [])

  let filteredTickets: Ticket[] = []

  const date = new Date()
  const currentMonth = date.getMonth()

  const lastMonth = new Date(
    date.getFullYear(),
    currentMonth - 1,
    date.getDate()
  )

  const twoMonthsAgo = new Date(
    date.getFullYear(),
    currentMonth - 2,
    date.getDate()
  )

  switch (tab) {
    case 'current':
      filteredTickets = tickets.filter(
        (ticket) => ticket.date.getMonth() === currentMonth
      )
      break
    case 'lastMonth':
      filteredTickets = tickets.filter(
        (ticket) => ticket.date.getMonth() === lastMonth.getMonth()
      )
      break
    case 'twoMonthsAgo':
      filteredTickets = tickets.filter(
        (ticket) => ticket.date.getMonth() === twoMonthsAgo.getMonth()
      )
      break
    default:
      filteredTickets = tickets
  }

  return (
    <div className="w-full">
      <div className="my-4 mx-4 flex border-b border-b-[#E1D8FD] text-sm font-normal text-[#98A2B3]">
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
        <button
          className={`p-4 font-medium ${
            tab === 'lastMonth'
              ? 'border-b-2 border-[#6A3CF5] text-[#6A3CF5]'
              : ''
          }`}
          onClick={() => setTab('lastMonth')}
        >
          {`${lastMonth.toLocaleDateString(undefined, {
            day: 'numeric',
            month: 'long',
            year: 'numeric',
          })}`}
        </button>
        <button
          className={`p-4 font-medium ${
            tab === 'twoMonthsAgo'
              ? 'border-b-2 border-[#6A3CF5] text-[#6A3CF5]'
              : ''
          }`}
          onClick={() => setTab('twoMonthsAgo')}
        >
          {`${twoMonthsAgo.toLocaleDateString(undefined, {
            day: 'numeric',
            month: 'long',
            year: 'numeric',
          })}`}
        </button>
      </div>

      <h2 className="m-4 font-medium text-[#FF60D5]">
        My Tickets: {tickets.length}{' '}
      </h2>
      <div className="grid-rows-10 md:grid-rows-19 m-4 grid max-h-[500px] grid-cols-5 gap-4 overflow-auto md:grid-cols-10">
        <div className="rounded-lg border border-dashed border-black">
          <div
            className="rounded-l-lg bg-[#9270FF] py-2 text-center text-sm text-[#FAF9FE]"
            style={{ width: '80%' }}
          >
            <div>80%</div>
          </div>
        </div>
        {filteredTickets.map((ticket, i) => (
          <div
            key={i}
            className="rounded-lg border bg-[#9270FF] py-2 text-center text-sm text-[#FAF9FE]"
          >
            <div>#{ticket.number}</div>
          </div>
        ))}
      </div>
    </div>
  )
}

export default TicketsGrid
