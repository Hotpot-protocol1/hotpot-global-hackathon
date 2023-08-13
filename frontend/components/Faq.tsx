import { FC } from 'react'
const Faq: FC = () => {
  return (
    <div className="m-4 mt-11">
      <h2 className="mb-4 font-medium text-[#FF60D5]">
        Frequently asked questions
      </h2>
      <div className="reservoir-subtitle w-full text-sm">
        <div className="mb-4 rounded-lg border bg-white p-6">
          <h3 className="reservoir-title mb-2 text-lg font-medium">
            What is a Pot O' Gold?
          </h3>
          <p className="text-gray-700 ">
            Pot O’ Gold is a provably-fair community jackpot where users earn
            golden tickets for an opportunity to win up to 100 ETH!
          </p>
        </div>
        <div className="mb-4 rounded-lg border bg-white p-6">
          <h3 className="reservoir-title mb-2 text-lg font-medium">
            How do I earn Golden Tickets?
          </h3>
          <p className="text-gray-700">
            Golden Tickets are automatically earned every time you trade 0.2 ETH
            Hotpot NFTs.
          </p>
        </div>
        <div className="mb-4 rounded-lg border bg-white p-6">
          <h3 className="reservoir-title mb-2 text-lg font-medium">
            Can I buy Golden Tickets?
          </h3>
          <p className="text-gray-700">
            No entry purchases are required, golden tickets can only be earned
            while trading.
          </p>
        </div>
        <div className="mb-4 rounded-lg bg-white p-6 shadow">
          <h3 className="reservoir-title mb-2 text-lg font-medium">
            How is the Pot O' Gold funded?
          </h3>
          <p className="text-gray-700">
            For every Hotpot transaction, the full 1% transaction fee will be
            directly sent to the pot.
          </p>
        </div>
        <div className="mb-4 rounded-lg border bg-white p-6">
          <h3 className="reservoir-title mb-2 text-lg font-medium">
            How can I claim my winnings?
          </h3>
          <p className="text-gray-700">
            Go to “My Profile” under my tickets, click on your golden ticket to
            claim! All raffle winners have 30 days to claim their prize.
          </p>
        </div>
        <div className="mb-4 rounded-lg border bg-white p-6">
          <h3 className="reservoir-title mb-2 text-lg font-medium">
            Can I win multiple times?
          </h3>
          <p className="text-gray-700">
            Yes, each ticket has a chance to win, meaning you can win more than
            once!
          </p>
        </div>
      </div>
    </div>
  )
}

export default Faq
