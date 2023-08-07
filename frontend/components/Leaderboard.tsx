import { FC } from 'react'
import { FiAlertCircle } from 'react-icons/fi'

interface LeaderboardItem {
  rank: number
  name: string
  boost: string
  tickets24h: number
  totalTickets: number
}

const Leaderboard: FC = () => {
  const generateData = (): LeaderboardItem[] => {
    const Data: LeaderboardItem[] = []
    for (let i = 1; i <= 15; i++) {
      Data.push({
        rank: i,
        name: 'Mochi',
        boost: '1x',
        tickets24h: 628,
        totalTickets: 8080,
      })
    }
    return Data
  }

  const data = generateData()

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
            <tr className="rounded-lg border-b bg-[#F0EBFE] text-center text-sm">
              <td className="p-2">133</td>
              <td className="p-2 text-left">You</td>
              <td className="p-2 text-right text-[#0FA46E]">1x</td>
              <td className="p-2 text-right">628</td>
              <td className="p-2 text-right">8080</td>
            </tr>
            {data.map((item) => (
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
