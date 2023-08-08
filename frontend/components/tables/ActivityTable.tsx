import { optimizeImage } from 'lib/optmizeImage'
import { truncateAddress } from 'lib/truncateText'
import { DateTime } from 'luxon'
import Link from 'next/link'
import { FC, ReactElement, useEffect, useState } from 'react'
import Image from 'next/legacy/image'
import { useMediaQuery } from '@react-hookz/web'
import LoadingIcon from 'components/LoadingIcon'
import { FiExternalLink, FiRepeat, FiTrash2, FiXSquare } from 'react-icons/fi'
import useEnvChain from 'hooks/useEnvChain'
import { useAccount } from 'wagmi'
import { constants } from 'ethers'
import { FaSeedling } from 'react-icons/fa'
import FormatNativeCrypto from 'components/FormatNativeCrypto'
import {
  useCollectionActivity,
  useUsersActivity,
} from '@reservoir0x/reservoir-kit-ui'
import { useInView } from 'react-intersection-observer'
import MobileActivityFilter from 'components/filter/MobileActivityFilter'
import { Item } from '../../lib/getAllListedNFTs'
import useTix from 'lib/tix'
import { useHotpotContext } from 'context/HotpotContext'

const RESERVOIR_API_BASE = process.env.NEXT_PUBLIC_RESERVOIR_API_BASE
const MARKET_CONTRACT = '0x4cfef2903d920069984d30e39eb5d9a1c6e08fc0'
type CollectionActivityResponse = ReturnType<typeof useCollectionActivity>
type CollectionActivity = CollectionActivityResponse['data'][0]
export type CollectionActivityTypes = NonNullable<
  Exclude<Parameters<typeof useCollectionActivity>['0'], boolean>
>['types']

type UsersActivityResponse = ReturnType<typeof useCollectionActivity>
type UsersActivity = UsersActivityResponse['data'][0]
type ActivityResponse = CollectionActivityResponse | UsersActivityResponse
export type UserActivityTypes = NonNullable<
  Exclude<Parameters<typeof useUsersActivity>['1'], boolean>
>['types']

type Activity = CollectionActivity | UsersActivity
type ActivityTypes = Exclude<
  CollectionActivityTypes | UserActivityTypes,
  string
>

type Props = {
  data: ActivityResponse
  types: ActivityTypes
  onTypesChange: (types: ActivityTypes) => void
  emptyPlaceholder: ReactElement
}

const ActivityTable: FC<Props> = ({
  data,
  types,
  onTypesChange,
  emptyPlaceholder,
}) => {
  const headings = ['Event', 'Item', 'Amount', 'From', 'To', 'Rewards', 'Time']
  const isMobile = useMediaQuery('only screen and (max-width : 730px)')
  const filters = ['Sales', 'Listings', 'Transfer', 'Mints']
  const { listedNFTs, isLoadingNFTs } = useHotpotContext()
  const enabledFilters: typeof filters = []

  if (types?.includes('sale')) {
    enabledFilters.push('Sales')
  }
  if (types?.includes('ask')) {
    enabledFilters.push('Listings')
  }
  if (types?.includes('transfer')) {
    enabledFilters.push('Transfer')
  }
  if (types?.includes('mint')) {
    enabledFilters.push('Mints')
  }

  const { ref, inView } = useInView()

  const activities = data.data

  useEffect(() => {
    if (inView) data.fetchNextPage()
  }, [inView])

  return (
    <>
      {isMobile ? (
        <MobileActivityFilter
          filters={filters}
          enabledFilters={enabledFilters}
          data={data}
          onTypesChange={onTypesChange}
          types={types}
        />
      ) : (
        <div className="mt-2 flex flex-wrap gap-2 md:m-5 md:gap-4">
          {filters.map((filter, i) => {
            const isSelected = enabledFilters.includes(filter)
            return (
              <button
                disabled={data.isFetchingPage || data.isValidating}
                key={i}
                className={`flex gap-3 rounded-full px-4 py-3 md:hover:bg-primary-100 dark:md:hover:bg-neutral-600 ${
                  isSelected
                    ? 'border-[1px] border-transparent bg-primary-100 dark:bg-neutral-600'
                    : 'border-[1px] border-neutral-300 bg-white dark:bg-black'
                }`}
                onClick={() => {
                  let updatedTypes: Props['types'] = types?.slice() || []
                  let activityType:
                    | 'sale'
                    | 'ask'
                    | 'transfer'
                    | 'mint'
                    | undefined = undefined

                  if (filter === 'Sales') {
                    activityType = 'sale'
                  } else if (filter === 'Listings') {
                    activityType = 'ask'
                  } else if (filter === 'Transfer') {
                    activityType = 'transfer'
                  } else if (filter === 'Mints') {
                    activityType = 'mint'
                  }

                  if (!activityType) {
                    return
                  }

                  if (enabledFilters.includes(filter)) {
                    updatedTypes = updatedTypes.filter(
                      (type) => activityType !== type
                    )
                  } else {
                    updatedTypes.push(activityType)
                  }
                  onTypesChange(updatedTypes)
                }}
              >
                {filter}
              </button>
            )
          })}
        </div>
      )}
      {!data.isValidating && (!activities || activities.length === 0) ? (
        emptyPlaceholder
      ) : (
        <table className="w-full">
          {!isMobile && (
            <thead>
              <tr className="text-left">
                {headings.map((name, i) => (
                  <th
                    key={i}
                    className="px-6 py-3 text-left text-sm font-medium text-neutral-600 dark:text-white"
                  >
                    {name}
                  </th>
                ))}
              </tr>
            </thead>
          )}

          <tbody>
            {activities.map((activity, i) => {
              if (!activity) return null

              return (
                <ActivityTableRow
                  key={`${activity?.txHash}-${i}`}
                  activity={activity}
                  listedNFTs={listedNFTs}
                />
              )
            })}
            <tr ref={ref}></tr>
          </tbody>
        </table>
      )}

      {data.isValidating && (
        <div className="my-20 flex justify-center">
          <LoadingIcon />
        </div>
      )}
    </>
  )
}

type ActivityTableRowProps = {
  activity: Activity
  listedNFTs: Item[] | null
}
type ItemInfo = {
  itemId: number
  price: string
}
const ActivityTableRow: FC<ActivityTableRowProps> = ({
  activity,
  listedNFTs,
}) => {
  const isMobile = useMediaQuery('only screen and (max-width : 730px)')
  const { address } = useAccount()
  const [currentActivity, setCurrentActivity] = useState<ItemInfo | null>(null)
  const [toShortAddress, setToShortAddress] = useState<string>(
    activity?.toAddress || ''
  )
  const [fromShortAddress, setFromShortAddress] = useState<string>(
    activity?.fromAddress || ''
  )
  const [imageSrc, setImageSrc] = useState(
    activity?.token?.tokenImage ||
      `${RESERVOIR_API_BASE}/redirect/collections/${activity?.collection?.collectionImage}/image/v1`
  )
  const [timeAgo, setTimeAgo] = useState(activity?.timestamp || '')
  const envChain = useEnvChain()
  const tix = useTix(currentActivity?.price ?? '0')
  const blockExplorerBaseUrl =
    envChain?.blockExplorers?.default?.url || 'https://etherscan.io'
  const href = activity?.token?.tokenId
    ? `/${activity?.collection?.collectionId}/${activity?.token?.tokenId}`
    : `/collections/${activity?.collection?.collectionId}`

  useEffect(() => {
    let toShortAddress = truncateAddress(activity?.toAddress || '')
    let fromShortAddress = truncateAddress(activity?.fromAddress || '')
    if (!!address) {
      if (address?.toLowerCase() === activity?.toAddress?.toLowerCase()) {
        toShortAddress = 'You'
      }
      if (address?.toLowerCase() === activity?.fromAddress?.toLowerCase()) {
        fromShortAddress = 'You'
      }
      if (
        MARKET_CONTRACT?.toLowerCase() === activity?.toAddress?.toLowerCase()
      ) {
        toShortAddress = 'Hotpot'
      }
      if (
        MARKET_CONTRACT?.toLowerCase() === activity?.fromAddress?.toLowerCase()
      ) {
        fromShortAddress = 'Hotpot'
      }
    }
    setToShortAddress(toShortAddress)
    setFromShortAddress(fromShortAddress)
    setTimeAgo(
      activity?.timestamp
        ? DateTime.fromSeconds(activity.timestamp).toRelative() || ''
        : ''
    )
  }, [activity, address])

  useEffect(() => {
    if (activity?.token?.tokenImage) {
      setImageSrc(optimizeImage(activity?.token?.tokenImage, 48))
    } else if (activity?.collection?.collectionImage) {
      setImageSrc(optimizeImage(activity?.collection?.collectionImage, 48))
    }
  }, [activity])

  if (!activity) {
    return null
  }

  let activityDescription = ''

  const logos = {
    transfer: (
      <FiRepeat className="w- mr-1 h-4 w-4 text-neutral-400 md:mr-[10px] md:h-5 md:w-5" />
    ),
    mint: (
      <FaSeedling className="mr-1 h-4 w-4 text-neutral-400 md:mr-[10px] md:h-5 md:w-5" />
    ),
    burned: (
      <FiTrash2 className="mr-1 h-4 w-4 text-neutral-400 md:mr-[10px] md:h-5 md:w-5" />
    ),
    listing_canceled: (
      <FiXSquare className="mr-1 h-4 w-4 text-neutral-400 md:mr-[10px] md:h-5 md:w-5" />
    ),
    ask: null,
    bid: null,
  }

  if (
    activity.fromAddress?.toLowerCase() ===
    '0x4cfef2903d920069984d30e39eb5d9a1c6e08fc0'
  ) {
    activityDescription = 'Sale'
    logos.transfer = <img src="/hotpot.png" className="mr-2 h-6 w-6" />
  } else if (
    activity.toAddress?.toLowerCase() ===
    '0x4cfef2903d920069984d30e39eb5d9a1c6e08fc0'
  ) {
    activityDescription = 'Listing'
    logos.transfer = <img src="/hotpot.png" className="mr-2 h-6 w-6" />
  } else {
    switch (activity?.type) {
      case 'ask_cancel': {
        activityDescription = 'Listing Canceled'
        break
      }
      case 'mint': {
        activityDescription = 'Mint'
        break
      }
      case 'ask': {
        activityDescription = 'Listing'
        break
      }
      case 'transfer': {
        activityDescription = 'Transfer'
        break
      }
      case 'sale': {
        activityDescription = 'Sale'
        break
      }
      default: {
        if (activity.type) activityDescription = activity.type
        break
      }
    }
  }
  const tokenId = activity?.token?.tokenId
  const contract = activity?.collection?.collectionId
  const findItem = (
    contractToFind: string,
    tokenIdToFind: string
  ): ItemInfo | null => {
    if (!listedNFTs) {
      return null
    }

    for (const item of listedNFTs) {
      const { itemId, nft, tokenId: tokenIdInArray, price } = item

      if (
        nft.toLowerCase() === contractToFind.toLowerCase() &&
        tokenIdInArray === tokenIdToFind
      ) {
        return { itemId, price }
      }
    }

    return null
  }

  useEffect(() => {
    if (listedNFTs && contract && tokenId) {
      const currentActivity = findItem(contract, tokenId)
      setCurrentActivity(currentActivity)
    }
  }, [listedNFTs, contract, tokenId])

  if (isMobile) {
    return (
      <tr
        key={activity.txHash}
        className="h-24 border-b border-gray-300 dark:border-[#525252]"
      >
        <td className="flex flex-col gap-3">
          <div className="mt-6 flex items-center">
            {/* @ts-ignore */}
            {activity.type && logos[activity.type]}
            {!!activity.order?.source?.icon && (
              <img
                className="mr-2 inline h-3 w-3"
                // @ts-ignore
                src={activity.order?.source?.icon || ''}
                alt={`${activity.order?.source?.name} Source`}
              />
            )}
            <span className="text-sm capitalize text-neutral-600 dark:text-neutral-300">
              {activityDescription}
            </span>
          </div>
          <div className="flex items-center justify-between">
            <Link href={href} passHref legacyBehavior={true}>
              <a className="flex items-center">
                <Image
                  className="rounded object-cover"
                  loader={({ src }) => src}
                  src={imageSrc}
                  alt={`${activity.token?.tokenName} Token Image`}
                  width={48}
                  height={48}
                />
                <div className="ml-2 grid truncate">
                  <div className="reservoir-h6 dark:text-white">
                    {activity.token?.tokenName ||
                      activity.token?.tokenId ||
                      activity.collection?.collectionName}
                  </div>
                </div>
              </a>
            </Link>
            {activity.price &&
            activity.price !== 0 &&
            activity.type &&
            !['transfer', 'mint'].includes(activity.type) ? (
              <FormatNativeCrypto amount={activity.price} />
            ) : null}
            {currentActivity?.price}
          </div>

          <div className="flex items-center justify-between">
            <div className="reservoir-small">
              <span className="mr-1 font-light text-neutral-600 dark:text-neutral-300">
                From
              </span>
              {activity.fromAddress &&
              activity.fromAddress !== constants.AddressZero ? (
                <Link
                  href={`/address/${activity.fromAddress}`}
                  legacyBehavior={true}
                >
                  <a className="font-light text-primary-700 dark:text-primary-300">
                    {fromShortAddress}
                  </a>
                </Link>
              ) : (
                <span className="font-light">-</span>
              )}
              <span className="mx-1 font-light text-neutral-600 dark:text-neutral-300">
                to{' '}
              </span>
              {activity.toAddress &&
              activity.toAddress !== constants.AddressZero ? (
                <Link
                  href={`/address/${activity.toAddress}`}
                  legacyBehavior={true}
                >
                  <a className="font-light text-primary-700 dark:text-primary-300">
                    {toShortAddress}
                  </a>
                </Link>
              ) : (
                <span className="font-light">-</span>
              )}
              <div className="mb-4 flex items-center justify-between gap-2 font-light text-neutral-600 dark:text-neutral-300 md:justify-start">
                {timeAgo}
              </div>
            </div>
            {activity.txHash && (
              <Link
                href={`${blockExplorerBaseUrl}/tx/${activity.txHash}`}
                legacyBehavior={true}
              >
                <a
                  target="_blank"
                  rel="noopener noreferrer"
                  className="mb-4 flex items-center justify-between gap-2 font-light text-neutral-600 dark:text-neutral-300 md:justify-start"
                >
                  <FiExternalLink className="h-4 w-4 text-primary-700 dark:text-primary-300" />
                </a>
              </Link>
            )}
          </div>
        </td>
      </tr>
    )
  }

  return (
    <tr
      key={activity.txHash}
      className="h-24 border-b border-gray-300 dark:border-[#525252]"
    >
      <td className="px-6 py-4">
        <div className="mr-2.5 flex items-center">
          {/* @ts-ignore */}
          {activity.type && logos[activity.type]}
          {!!activity.order?.source?.icon && (
            <img
              className="mr-2 h-6 w-6"
              // @ts-ignore
              src={activity.order?.source?.icon || ''}
              alt={`${activity.order?.source?.name} Source`}
            />
          )}
          <span className="text-sm capitalize text-neutral-600 dark:text-neutral-300">
            {activityDescription}
          </span>
        </div>
      </td>
      <td className="px-6 py-4">
        <Link href={href} passHref legacyBehavior={true}>
          <a className="mr-2.5 flex items-center">
            <Image
              className="rounded object-cover"
              loader={({ src }) => src}
              src={imageSrc}
              alt={`${activity.token?.tokenName} Token Image`}
              width={48}
              height={48}
            />
            <div className="ml-2 grid truncate">
              <div className="reservoir-h6 dark:text-white">
                {activity.token?.tokenName ||
                  activity.token?.tokenId ||
                  activity.collection?.collectionName}
              </div>
            </div>
          </a>
        </Link>
      </td>
      <td className="px-6 py-4">
        {activity.price &&
        activity.price !== 0 &&
        activity.type &&
        !['transfer', 'mint'].includes(activity.type) ? (
          <FormatNativeCrypto amount={activity.price} />
        ) : null}{' '}
        {currentActivity?.price ? (
          <div className="flex flex-row items-center justify-center gap-1 font-semibold">
            <img src="/eth.svg" alt="price" className="h-3 w-3" />
            {currentActivity.price}
          </div>
        ) : null}
      </td>
      <td className="px-6 py-4">
        {activity.fromAddress &&
        activity.fromAddress !== constants.AddressZero ? (
          <Link href={`/address/${activity.fromAddress}`} legacyBehavior={true}>
            <a className="ml-2.5 mr-2.5 font-light text-primary-700 dark:text-primary-300">
              {fromShortAddress}
            </a>
          </Link>
        ) : (
          <span className="ml-2.5 mr-2.5 font-light">-</span>
        )}
      </td>
      <td className="px-6 py-4">
        {activity.toAddress && activity.toAddress !== constants.AddressZero ? (
          <Link href={`/address/${activity.toAddress}`} legacyBehavior={true}>
            <a className="ml-2.5 mr-2.5 font-light text-primary-700 dark:text-primary-300">
              {toShortAddress}
            </a>
          </Link>
        ) : (
          <span className="ml-2.5 mr-2.5 font-light">-</span>
        )}
      </td>
      <td className="px-6 py-4">
        <div className="flex items-center gap-2 whitespace-nowrap text-right font-light text-neutral-600 dark:text-neutral-300">
          {currentActivity?.price ? (
            <>
              <div className="rounded border border-[#0FA46E] bg-[#DBF1E4] px-1 text-sm text-[#0FA46E]">
                +{tix} Tickets
              </div>
            </>
          ) : (
            '-'
          )}
        </div>
      </td>
      <td className="px-6 py-4">
        <div className="flex items-center gap-2 whitespace-nowrap font-light text-neutral-600 dark:text-neutral-300">
          {timeAgo}
          {activity.txHash && (
            <Link
              href={`${blockExplorerBaseUrl}/tx/${activity.txHash}`}
              legacyBehavior={true}
            >
              <a target="_blank" rel="noopener noreferrer">
                <FiExternalLink className="h-4 w-4 text-primary-700 dark:text-primary-300" />
              </a>
            </Link>
          )}
        </div>
      </td>
    </tr>
  )
}

export default ActivityTable
