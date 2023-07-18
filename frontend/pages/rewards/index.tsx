import ConnectWalletButton from 'components/ConnectWalletButton'
import Layout from 'components/Layout'
import { NextPage } from 'next'
import { useAccount } from 'wagmi'
import Toast from 'components/Toast'
import toast from 'react-hot-toast'
import { ComponentProps } from 'react'
import Hero from 'components/Hero'
import useMounted from 'hooks/useMounted'
import TicketsGrid from 'components/TicketsGrid'
import Leaderboard from 'components/Leaderboard'
import Faq from 'components/Faq'
import Footer from 'components/Footer'

const Rewards: NextPage = () => {
  const { address, isConnected } = useAccount()
  const isMounted = useMounted()

  if (!isMounted) {
    return null
  }

  const setToast: (data: ComponentProps<typeof Toast>['data']) => any = (
    data
  ) => toast.custom((t) => <Toast t={t} toast={toast} data={data} />)

  const leaderboardData = [
    {
      rank: 133,
      name: 'You',
      boost: '1x',
      tickets24h: 628,
      totalTickets: 8080,
    },
    { rank: 1, name: 'John', boost: 'Yes', tickets24h: 10, totalTickets: 50 },
  ]

  return (
    <Layout navbar={{}}>
      <div className="col-span-full mt-4 mb-12 px-2 md:mt-5 lg:px-12">
        <Hero variant="rewards" />
        <TicketsGrid />
        <Leaderboard />
        {isConnected ? <></> : <div className=""></div>}
        <Faq />
      </div>
      <Footer />
    </Layout>
  )
}

export default Rewards
