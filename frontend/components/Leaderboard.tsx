import useLeaderboardData from 'hooks/useLeaderboardData'
import { FC, useMemo } from 'react'
import { FiAlertCircle } from 'react-icons/fi'

interface LeaderboardItem {
  rank: number
  name: string
  boost: string
  tickets24h: number
  totalTickets: number
}

const Leaderboard: FC = () => {
  const { data } = useLeaderboardData({ potId: 2, chain: "goerli" });

  const leaderboardItems: Array<LeaderboardItem> = useMemo(() => {
    return data
      ?.sort((prev, next) => next.num_of_tickets - prev.num_of_tickets)
      .map((item, item_idx) => ({
        rank: item_idx + 1,
        name: item.wallet_address as string,
        boost: "1x",
        tickets24h: NaN,
        totalTickets: item.num_of_tickets,
      })) || [];
  }, [data]);

  return (
    <div className="m-4 mt-11">
      <h2 className="mb-4 font-medium text-[#FF60D5]">
        Ticket holder Leaderboard
      </h2>
      <div className="w-full rounded-lg border p-4">
        <table className="w-full">
          <thead>
            <tr className="text-center">
              <th className="p-2 font-medium">Rank</th>
              <th className="w-2/4 p-2 text-left font-medium">Name</th>
              <th className="p-2 text-right font-medium">
                Boost{' '}
                <FiAlertCircle className="ml-1 inline-block text-[#98A2B3]" />
              </th>
              <th className="p-2 text-right font-medium">24h Tickets</th>
              <th className="p-2 text-right font-medium">Total Tickets</th>
            </tr>
          </thead>
          <tbody>
            {leaderboardItems.map((item) => (
              <tr key={item.rank} className="m-2 border-b text-center text-sm">
                <td className="p-2">{item.rank}</td>
                <td className="w-1/3 p-2 text-left">{item.name}</td>
                <td className="p-2 text-right text-[#0FA46E]">{item.boost}</td>
                <td className="p-2 text-right">{item.tickets24h}</td>
                <td className="p-2 text-right">{item.totalTickets}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default Leaderboard
