import { FC } from 'react'
import { FiAlertCircle } from 'react-icons/fi'

interface LeaderboardItem {
  rank: number
  name: string
  boost: string
  tickets24h: number
  totalTickets: number
}

const Faq: FC = () => {
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
        Frequently asked questions
      </h2>
      <div className="w-full text-xs">
        <div className="mb-4 rounded-lg border bg-white p-6">
          <h3 className="mb-2 text-base font-medium">What is a Pot O’ Gold?</h3>
          <p className="text-gray-700">
            Flotapay is an online bill payment app that allows users in Nigeria
            to pay their bills and make other payments easily and securely from
            their smartphones or other devices.
          </p>
        </div>
        <div className="mb-4 rounded-lg border bg-white p-6">
          <h3 className="mb-2 text-base font-medium">What is a Pot O’ Gold?</h3>
          <p className="text-gray-700">
            Flotapay is an online bill payment app that allows users in Nigeria
            to pay their bills and <br /> make other payments easily and
            securely from their smartphones or other devices.
          </p>
        </div>
        <div className="mb-4 rounded-lg border bg-white p-6">
          <h3 className="mb-2 text-base font-medium">What is a Pot O’ Gold?</h3>
          <p className="text-gray-700">
            Flotapay is an online bill payment app that allows users in Nigeria
            to pay their bills and make other payments easily and securely from
            their smartphones or other devices.
          </p>
        </div>
        <div className="mb-4 rounded-lg bg-white p-6 shadow">
          <h3 className="mb-2 text-base font-medium">What is a Pot O’ Gold?</h3>
          <p className="text-gray-700">
            Flotapay is an online bill payment app that allows users in Nigeria
            to pay their bills and make other payments easily and securely from
            their smartphones or other devices.
          </p>
        </div>
        <div className="mb-4 rounded-lg border bg-white p-6">
          <h3 className="mb-2 text-base font-medium">What is a Pot O’ Gold?</h3>
          <p className="text-gray-700">
            Flotapay is an online bill payment app that allows users in Nigeria
            to pay their bills and make other payments easily and securely from
            their smartphones or other devices.
          </p>
        </div>
      </div>
    </div>
  )
}

export default Faq
