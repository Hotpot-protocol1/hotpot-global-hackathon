import ConnectWalletButton from 'components/ConnectWalletButton'
import Layout from 'components/Layout'
import { NextPage } from 'next'
import { useAccount } from 'wagmi'
import Toast from 'components/Toast'
import toast from 'react-hot-toast'
import { ComponentProps } from 'react'
import Hero from 'components/Hero'
import useMounted from 'hooks/useMounted'

const Rewards: NextPage = () => {
  const { address, isConnected } = useAccount()
  const isMounted = useMounted()

  if (!isMounted) {
    return null
  }

  const setToast: (data: ComponentProps<typeof Toast>['data']) => any = (
    data
  ) => toast.custom((t) => <Toast t={t} toast={toast} data={data} />)

  return (
    <Layout navbar={{}}>
      <div className="col-span-full mt-4 mb-12 px-2 md:mt-5 lg:px-12">
        <Hero variant="rewards" />
        {isConnected ? (
          <></>
        ) : (
          <div className="flex flex-col items-center gap-y-6 py-16 text-center md:py-32">
            <h3 className="mb-2 text-2xl font-semibold dark:text-white sm:text-3xl">
              Rewards
            </h3>

            <ConnectWalletButton>
              <span className="w-40">Connect Wallet</span>
            </ConnectWalletButton>
          </div>
        )}
      </div>
    </Layout>
  )
}

export default Rewards
