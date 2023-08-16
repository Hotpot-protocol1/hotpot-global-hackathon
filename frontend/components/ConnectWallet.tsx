import { FC } from 'react'
import {
  useAccount,
  useBalance,
  useConnect,
  useDisconnect,
  useEnsAvatar,
  useEnsName,
  Address,
} from 'wagmi'
import * as DropdownMenu from '@radix-ui/react-dropdown-menu'
import { HiOutlineLogout } from 'react-icons/hi'
import { FaWallet } from 'react-icons/fa'
import { truncateAddress, truncateEns } from 'lib/truncateText'
import Link from 'next/link'
import FormatNativeCrypto from './FormatNativeCrypto'
import ConnectWalletButton from 'components/ConnectWalletButton'
import CopyClipboard from './CopyClipboard'
import useMounted from 'hooks/useMounted'
import Avatar from './Avatar'

const ConnectWallet: FC = () => {
  const account = useAccount()
  const { data: ensAvatar } = useEnsAvatar({ address: account?.address })
  const { data: ensName } = useEnsName({ address: account?.address })

  const { connectors } = useConnect()
  const { disconnect } = useDisconnect()
  const wallet = connectors[0]
  const isMounted = useMounted()

  if (!isMounted) {
    return null
  }

  if (!account.isConnected)
    return (
      <ConnectWalletButton>
        <img src="/icons/wallet.svg" alt="Wallet Icon" />
      </ConnectWalletButton>
    )

  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger className="btn-primary-outline ml-auto rounded-full border-transparent p-0 normal-case dark:border-neutral-600 dark:bg-neutral-900 dark:ring-primary-900 dark:focus:ring-4">
        <Avatar address={account.address} avatar={ensAvatar} size={40} />
      </DropdownMenu.Trigger>

      <DropdownMenu.Content align="end" sideOffset={6}>
        <div className="w-48 space-y-1 rounded rounded-t bg-white px-1.5 py-2 shadow-md radix-side-bottom:animate-slide-down dark:bg-neutral-900 md:w-56">
          <div className="gap-0 rounded border-b">
            <div className="items-left group flex w-full flex-col justify-between rounded px-4 py-3 font-medium outline-none transition dark:bg-blue-900">
              {ensName ? (
                <span className="font-normal text-gray-500 dark:text-gray-300">
                  {truncateEns(ensName)}
                  <CopyClipboard content={ensName as string} />
                </span>
              ) : (
                <div className="flex flex-row items-center gap-2 text-gray-500">
                  {truncateAddress(account.address || '')}
                  <CopyClipboard content={account.address as string} />
                </div>
              )}
              <div className="mt-2 flex flex-row items-center gap-2">
                <FaWallet className="ml-1 h-4 w-4 text-gray-500" />
                <div className="flex flex-row items-center rounded border border-violet-400 px-3 pr-5 text-sm text-gray-800 dark:text-gray-300">
                  {account.address && <Balance address={account.address} />}
                </div>
              </div>
            </div>
          </div>

          <Link href={`/address/${account.address}`} legacyBehavior={true}>
            <DropdownMenu.Item asChild>
              <a className="group flex w-full cursor-pointer items-center justify-between rounded px-4 py-3 outline-none transition hover:bg-neutral-100 focus:bg-neutral-100 dark:hover:bg-neutral-800 dark:focus:bg-neutral-800">
                My Profile
              </a>
            </DropdownMenu.Item>
          </Link>

          <DropdownMenu.Item asChild>
            <button
              key={wallet.id}
              onClick={() => {
                disconnect()
              }}
              className="group flex w-full cursor-pointer items-center justify-between gap-3 rounded px-4 py-3 outline-none transition hover:bg-neutral-100 focus:bg-neutral-100 dark:hover:bg-neutral-800 dark:focus:bg-neutral-800"
            >
              <span>Disconnect</span>
              <HiOutlineLogout className="h-6 w-7" />
            </button>
          </DropdownMenu.Item>
        </div>
      </DropdownMenu.Content>
    </DropdownMenu.Root>
  )
}

export default ConnectWallet

type Props = {
  address: string
}

export const Balance: FC<Props> = ({ address }) => {
  const { data: balance } = useBalance({ address: address as Address })
  return <FormatNativeCrypto amount={balance?.value} />
}
