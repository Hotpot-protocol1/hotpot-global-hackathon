import ConnectWalletButton from 'components/ConnectWalletButton'
import Layout from 'components/Layout'
import type { GetStaticProps, InferGetStaticPropsType, NextPage } from 'next'
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
import getPrizePool, { Item } from '../../lib/getPrizePool'
import getTicketCost from 'lib/getTicketCost'
import PotResultBanner from 'components/PotResultBanner'

type Props = InferGetStaticPropsType<typeof getStaticProps> & {
  prizePool: Item | null
  ticketCost: string | null
}

const Rewards: NextPage<Props> = ({ prizePool, ticketCost }) => {
  const { isConnected } = useAccount()
  const isMounted = useMounted()

  if (!isMounted) {
    return null
  }

  const setToast: (data: ComponentProps<typeof Toast>['data']) => any = (
    data
  ) => toast.custom((t) => <Toast t={t} toast={toast} data={data} />)

  return (
    <Layout navbar={{}}>
      <PotResultBanner />
      <div className="col-span-full mt-4 mb-12 px-2 md:mt-5 lg:px-12">
        <Hero variant="rewards" />
        <TicketsGrid prizePool={prizePool} />
        <Leaderboard />
        {isConnected ? <></> : <div className=""></div>}
        <Faq />
      </div>
      <Footer />
    </Layout>
  )
}

export default Rewards

export const getStaticProps: GetStaticProps<{}> = async () => {
  const prizePool = await getPrizePool()
  const ticketCost = await getTicketCost()

  return {
    props: { prizePool, ticketCost },
    revalidate: 20,
  }
}
