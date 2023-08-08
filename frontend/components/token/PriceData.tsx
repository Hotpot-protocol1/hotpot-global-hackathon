import BuyNow from 'components/BuyNow'
import { TokenDetails } from 'types/reservoir'
import {
  ListModal,
  useReservoirClient,
  useTokens,
} from '@reservoir0x/reservoir-kit-ui'
import React, {
  ComponentPropsWithoutRef,
  FC,
  ReactNode,
  useEffect,
  useState,
} from 'react'
import { useRecoilState, useRecoilValue } from 'recoil'
import { useAccount, useNetwork, useSigner } from 'wagmi'
import { setToast } from './setToast'
import recoilCartTokens, {
  getCartCurrency,
  getPricingPools,
  getTokensMap,
} from 'recoil/cart'
import FormatCrypto from 'components/FormatCrypto'
import { Collection } from 'types/reservoir'
import { formatDollar } from 'lib/numbers'
import useCoinConversion from 'hooks/useCoinConversion'
import SwapCartModal from 'components/SwapCartModal'
import { FaShoppingCart } from 'react-icons/fa'
import ConnectWalletButton from 'components/ConnectWalletButton'
import useMounted from 'hooks/useMounted'
import { useRouter } from 'next/router'
import { getPricing } from 'lib/token/pricing'
import { useContract } from 'wagmi'
import { abi, NFTMarketplace_CONTRACT_SEP } from '../../contracts/index'
import { CgSpinner } from 'react-icons/cg'
import BuyModal from 'components/modal/BuyModal'
import ListModal2 from '../../components/modal/ListModal'
import useTix from 'lib/tix'
import { useHotpotContext } from 'context/HotpotContext'

const CHAIN_ID = process.env.NEXT_PUBLIC_CHAIN_ID
const SOURCE_ID = process.env.NEXT_PUBLIC_SOURCE_ID
const SOURCE_ICON = process.env.NEXT_PUBLIC_SOURCE_ICON
const API_BASE =
  process.env.NEXT_PUBLIC_RESERVOIR_API_BASE || 'https://api.reservoir.tools'
const CURRENCIES = process.env.NEXT_PUBLIC_LISTING_CURRENCIES
const HOTPOT_CONTRACT = process.env.NEXT_HOTPOT_MARKETPLACE_CONTRACT_SEP

type Props = {
  details: ReturnType<typeof useTokens>
  collection?: Collection
  isOwner: boolean
  tokenDetails?: TokenDetails
}

type ItemInfo = {
  itemId: number
  price: string
}

type ListingCurrencies = ComponentPropsWithoutRef<
  typeof ListModal
>['currencies']
let listingCurrencies: ListingCurrencies = undefined

if (CURRENCIES) {
  listingCurrencies = JSON.parse(CURRENCIES)
}

const PriceData: FC<Props> = ({
  details,
  collection,
  isOwner,
  tokenDetails,
}) => {
  const router = useRouter()
  const isMounted = useMounted()
  const [cartTokens, setCartTokens] = useRecoilState(recoilCartTokens)
  const tokensMap = useRecoilValue(getTokensMap)
  const cartCurrency = useRecoilValue(getCartCurrency)
  const cartPools = useRecoilValue(getPricingPools)
  const accountData = useAccount()
  const { data: signer } = useSigner()
  const { chain: activeChain } = useNetwork()
  const reservoirClient = useReservoirClient()
  const [clearCartOpen, setClearCartOpen] = useState(false)
  const [cartToSwap, setCartToSwap] = useState<undefined | typeof cartTokens>()
  const account = useAccount()
  const { listedNFTs, isLoadingNFTs } = useHotpotContext()
  const [currentNFT, setCurrentNFT] = useState<ItemInfo | null>(null)
  const token = details.data ? details.data[0] : undefined
  const tokenId = token?.token?.tokenId
  const contract = token?.token?.contract

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
      const currentNFT = findItem(contract, tokenId)
      setCurrentNFT(currentNFT)
    }
  }, [listedNFTs, contract, tokenId])

  let floorAskPrice = getPricing(cartPools, token)
  let canAddToCart = true

  // Disabling the rules of hooks here due to erroneous error message,
  //  the linter is likely confused due to two custom hook calls of the same name
  // eslint-disable-next-line react-hooks/rules-of-hooks
  const topBidUsdConversion = useCoinConversion(
    token?.market?.topBid?.price?.currency?.symbol ? 'usd' : undefined,
    token?.market?.topBid?.price?.currency?.symbol
  )

  // eslint-disable-next-line react-hooks/rules-of-hooks
  const floorAskUsdConversion = useCoinConversion(
    floorAskPrice?.currency?.symbol ? 'usd' : undefined,
    floorAskPrice?.currency?.symbol
  )

  if (!isMounted) {
    return null
  }

  const topBidUsdPrice =
    topBidUsdConversion && token?.market?.topBid?.price?.amount?.decimal
      ? topBidUsdConversion * token?.market?.topBid?.price?.amount?.decimal
      : null

  const floorAskUsdPrice =
    floorAskUsdConversion && floorAskPrice?.amount?.decimal
      ? floorAskUsdConversion * floorAskPrice?.amount?.decimal
      : null

  const listSourceName = token?.market?.floorAsk?.source?.name as
    | string
    | undefined
  const listSourceDomain = token?.market?.floorAsk?.source?.domain as
    | string
    | undefined

  const offerSourceName = token?.market?.topBid?.source?.name as
    | string
    | undefined
  const offerSourceDomain = token?.market?.topBid?.source?.domain as
    | string
    | undefined

  let isLocalListed = false

  if (
    reservoirClient?.source &&
    listSourceDomain &&
    reservoirClient.source === listSourceDomain
  ) {
    isLocalListed = true
  } else if (SOURCE_ID && listSourceName && SOURCE_ID === listSourceName) {
    isLocalListed = true
  }

  const listSourceLogo =
    isLocalListed && SOURCE_ICON
      ? SOURCE_ICON
      : `${API_BASE}/redirect/sources/${
          listSourceDomain || listSourceName
        }/logo/v2`

  if (!CHAIN_ID) return null

  const isTopBidder =
    accountData.isConnected &&
    token?.market?.topBid?.maker?.toLowerCase() ===
      accountData?.address?.toLowerCase()
  const isListed = token
    ? floorAskPrice !== null && token?.token?.kind !== 'erc1155'
    : false
  const isInTheWrongNetwork = Boolean(signer && activeChain?.id !== +CHAIN_ID)

  const offerSourceLogo = `${API_BASE}/redirect/sources/${
    offerSourceDomain || offerSourceName
  }/logo/v2`

  const listSourceRedirect = `${API_BASE}/redirect/sources/${
    listSourceDomain || listSourceName
  }/tokens/${contract}:${tokenId}/link/v2`

  const isInCart = Boolean(tokensMap[`${contract}:${tokenId}`])

  const isHotpot =
    tokenDetails?.owner == '0x4cfef2903d920069984d30e39eb5d9a1c6e08fc0'

  const tix = useTix(currentNFT?.price ?? '0')

  return (
    <div className="col-span-full md:col-span-4 lg:col-span-5 lg:col-start-2">
      <article className="col-span-full rounded-2xl border border-gray-300 bg-white p-6 dark:border-neutral-600 dark:bg-black">
        <div className="grid grid-cols-1 gap-6">
          {isHotpot && currentNFT ? (
            <div className="flex flex-row">
              <div className="flex-grow">
                <div className="reservoir-h5 font-headings dark:text-white">
                  List Price
                </div>
                <div className="justify-left my-1 flex flex-row items-center gap-2">
                  <img
                    src="/hotpot.png"
                    alt="hotpot-marketplace"
                    className="h-5 w-5"
                  />
                  <p className="text-xs font-light"> Hotpot Marketplace</p>
                </div>
              </div>
              <div className="reservoir-h3 font-headings dark:text-white">
                <div className="flex flex-row items-center">
                  {tix > 0 && (
                    <div className="z-10 mx-2 flex items-center justify-center truncate rounded border border-[#0FA46E] bg-[#DBF1E4] px-2 text-sm font-normal text-[#0FA46E]">
                      +{tix} TIX
                    </div>
                  )}
                  <img
                    src="/eth.svg"
                    alt="hotpot-marketplace"
                    className="mr-1 h-5 w-5"
                  />{' '}
                  {currentNFT?.price}
                </div>
                <div className="text-sm text-neutral-600 dark:text-neutral-300">
                  {/* {formatDollar(usdPrice)} */}
                </div>
              </div>
            </div>
          ) : (
            <Price
              title="List Price"
              source={
                listSourceName && (
                  <a
                    target="_blank"
                    rel="noopener noreferrer"
                    href={listSourceRedirect}
                    className="reservoir-body flex items-center gap-2 dark:text-white"
                  >
                    on {listSourceName}
                    <img
                      className="h-6 w-6"
                      src={listSourceLogo}
                      alt="Source Logo"
                    />
                  </a>
                )
              }
              price={
                <FormatCrypto
                  amount={floorAskPrice?.amount?.decimal}
                  address={floorAskPrice?.currency?.contract}
                  decimals={floorAskPrice?.currency?.decimals}
                  logoWidth={30}
                  maximumFractionDigits={8}
                />
              }
              usdPrice={floorAskUsdPrice}
            />
          )}
        </div>

        <div className="mt-6 grid grid-cols-1 gap-3 md:grid-cols-2">
          {account.isDisconnected ? (
            <ConnectWalletButton className="w-full">
              <span>Connect Wallet</span>
            </ConnectWalletButton>
          ) : (
            <>
              {isOwner && (
                <ListModal2
                  trigger={
                    <button className="btn-primary-fill w-full dark:ring-primary-900 dark:focus:ring-4">
                      {floorAskPrice?.amount?.decimal
                        ? 'Create New Listing'
                        : 'List for Sale'}
                    </button>
                  }
                  collectionId={contract}
                  tokenId={tokenId}
                  onListingComplete={() => {
                    details && details.mutate()
                  }}
                  onListingError={(err: any) => {
                    if (err?.code === 4001) {
                      setToast({
                        kind: 'error',
                        message: 'You have canceled the transaction.',
                        title: 'User canceled transaction',
                      })
                      return
                    }
                    setToast({
                      kind: 'error',
                      message: 'The transaction was not completed.',
                      title: 'Could not list token',
                    })
                  }}
                />
              )}
              {!isOwner && (
                <BuyNow
                  buttonClassName="btn-primary-fill col-span-1"
                  data={{
                    details: details,
                  }}
                  signer={signer}
                  isInTheWrongNetwork={isInTheWrongNetwork}
                  mutate={details.mutate}
                />
              )}
              {isHotpot && !isOwner && (
                <BuyModal
                  trigger={
                    <button className="btn-primary-fill col-span-1">
                      Buy Now
                    </button>
                  }
                  itemId={currentNFT?.itemId}
                  price={currentNFT?.price}
                  tokenDetails={tokenDetails}
                />
              )}

              {isInCart && !isOwner && (
                <button
                  onClick={() => {
                    const newCartTokens = [...cartTokens]
                    const index = newCartTokens.findIndex(
                      (cartToken) =>
                        cartToken?.token?.contract === contract &&
                        cartToken?.token?.tokenId === tokenId
                    )
                    newCartTokens.splice(index, 1)
                    setCartTokens(newCartTokens)
                  }}
                  className="btn-primary-outline w-full dark:border-neutral-600 dark:text-white dark:ring-primary-900 dark:focus:ring-4"
                >
                  Remove
                  <FaShoppingCart className="ml-[10px] h-[18px] w-[18px] text-[#FF3B3B] dark:text-[#FF9A9A]" />
                </button>
              )}

              {!isInCart && !isOwner && isHotpot && canAddToCart && (
                <button
                  disabled={isInTheWrongNetwork}
                  onClick={() => {
                    if (token?.token && token.market) {
                      if (
                        !cartCurrency ||
                        floorAskPrice?.currency?.contract ===
                          cartCurrency?.contract
                      ) {
                        setCartTokens([
                          ...cartTokens,
                          {
                            token: token.token,
                            market: token.market,
                            itemId: currentNFT?.itemId ?? 0,
                            hotpotPrice: currentNFT?.price ?? '0',
                            tix: tix ?? 0,
                          },
                        ])
                      } else {
                        setCartToSwap([
                          {
                            token: token.token,
                            market: token.market,
                            itemId: currentNFT?.itemId ?? 0,
                            hotpotPrice: currentNFT?.price ?? '0',
                            tix: tix ?? 0,
                          },
                        ])
                        setClearCartOpen(true)
                      }
                    }
                  }}
                  className="btn-primary-outline w-full dark:border-neutral-600 dark:text-white dark:ring-primary-900 dark:focus:ring-4"
                >
                  Add to Cart
                  <FaShoppingCart className="ml-[10px] h-[18px] w-[18px] text-primary-700 dark:text-primary-100" />
                </button>
              )}

              {/* {isHotpot && canAddToCart && (
                <button
                  disabled={!floorAskPrice || !currentNFT}
                  onClick={() => {
                    if (token?.token && token.market) {
                      if (
                        !cartCurrency ||
                        floorAskPrice?.currency?.contract ===
                          cartCurrency?.contract
                      ) {
                        setCartTokens([
                          ...cartTokens,
                          {
                            token: token.token,
                            market: token.market,
                            itemId: currentNFT?.itemId ?? 0,
                            hotpotPrice: currentNFT?.price ?? '0',
                            tix: tix ?? 0,
                          },
                        ])
                      } else {
                        setCartToSwap([
                          {
                            token: token.token,
                            market: token.market,
                            itemId: currentNFT?.itemId ?? 0,
                            hotpotPrice: currentNFT?.price ?? '0',
                            tix: tix ?? 0,
                          },
                        ])
                        setClearCartOpen(true)
                      }
                    }
                  }}
                  className="w-full btn-primary-outline dark:border-neutral-600 dark:text-white dark:ring-primary-900 dark:focus:ring-4"
                >
                  Add to Cart
                  <FaShoppingCart className="ml-[10px] h-[18px] w-[18px] text-primary-700 dark:text-primary-100" />
                </button>
              )} */}
            </>
          )}
        </div>
      </article>

      <SwapCartModal
        open={clearCartOpen}
        setOpen={setClearCartOpen}
        cart={cartToSwap}
      />
    </div>
  )
}

export default PriceData

const Price: FC<{
  title: string
  price: ReactNode
  source?: ReactNode
  usdPrice: number | null
}> = ({ title, price, usdPrice, source }) => (
  <div className="flex flex-row">
    <div className="flex-grow">
      <div className="reservoir-h5 font-headings dark:text-white">{title}</div>
      {source}
    </div>
    <div className="reservoir-h3 font-headings dark:text-white">
      {price}
      <div className="text-sm text-neutral-600 dark:text-neutral-300">
        {formatDollar(usdPrice)}
      </div>
    </div>
  </div>
)
