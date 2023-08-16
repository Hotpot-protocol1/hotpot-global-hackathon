import React, { useState, useEffect, Dispatch, SetStateAction } from 'react'
import { ethers } from 'ethers'
import { useAccount, useContract, useProvider, useSigner } from 'wagmi'
import * as Dialog from '@radix-ui/react-dialog'
import { CgSpinner } from 'react-icons/cg'
import { HiCheckCircle, HiExclamationCircle, HiX } from 'react-icons/hi'
import { abi, NFTMarketplace_CONTRACT_SEP } from '../../contracts/index'
import { useTokens } from '@reservoir0x/reservoir-kit-ui'
import { mutate } from 'swr'
import { setToast } from 'components/token/setToast'
import Modal from './Modal'
import useTix from 'lib/tix'

type BuyCallbackData = {
  tokenId?: string
  collectionId?: string
}

type UseTokensReturnType = ReturnType<typeof useTokens>

export type Token = {
  token: NonNullable<
    NonNullable<NonNullable<UseTokensReturnType['data']>[0]>['token']
  >
  market: NonNullable<
    NonNullable<NonNullable<UseTokensReturnType['data']>[0]>['market']
  >
  itemId: number
  hotpotPrice: string
  tix: number
}

type Props = Pick<Parameters<typeof Modal>['0'], 'trigger'> & {
  openState?: [boolean, Dispatch<SetStateAction<boolean>>]
  itemIds?: number[]
  totalPrice?: string
  loading?: boolean
  onGoToToken?: () => any
  onBuyComplete?: (data: BuyCallbackData) => void
  onBuyError?: (error: Error, data: BuyCallbackData) => void
  onClose?: () => void
  setWaitingTx: (isWaiting: boolean) => void
  cartTokens: Token[]
  handleSuccess: () => void
  cartCount?: number
}

const BuyCartModal: React.FC<Props> = ({
  trigger,
  cartTokens,
  totalPrice,
  setWaitingTx,
  handleSuccess,
  cartCount,
}) => {
  const [isLoading, setIsLoading] = useState(false)
  const [priceLoading, setPriceLoading] = useState(false)
  const { address } = useAccount()
  const provider = useProvider()
  const { data: signer } = useSigner()
  const [isMounted, setIsMounted] = useState<boolean>(false)
  const [isApproved, setIsApproved] = useState<boolean>(false)

  const [error, setError] = useState<string | null>(null)
  const [isSuccess, setIsSuccess] = useState<boolean>(false)

  useEffect(() => {
    setIsMounted(true)
  }, [])

  const NftMarketplace = useContract({
    address: NFTMarketplace_CONTRACT_SEP,
    abi: abi,
    signerOrProvider: signer,
  })

  const handleSubmit = async () => {
    setWaitingTx(true)
    setIsLoading(true)
    setError(null)

    try {
      if (!NftMarketplace) {
        setError('NftMarketplace contract instance is not available.')
        setIsLoading(false)
        return
      }
      if (!totalPrice) {
        console.log('Wait for total price to load')
        return
      }

      const priceInWei = ethers.utils.parseEther(totalPrice.toString())

      const itemIds = cartTokens.map((tokenData) => tokenData.itemId)

      const allNFTs = itemIds.map(async (itemId) => {
        const buyNFT = await NftMarketplace.purchaseItem(itemId, {
          value: priceInWei,
        })
        console.log('Listing Transaction Hash:', buyNFT.hash)
        setIsApproved(true)
        await buyNFT.wait()
      })

      await Promise.all(allNFTs)

      setIsApproved(false)
      setIsLoading(false)
      setIsSuccess(true)
      setToast({
        kind: 'complete',
        message: '',
        title: 'Purchase Complete',
      })
    } catch (error) {
      setIsLoading(false)
      console.log(error)
      setError('Oops, something went wrong. Please try again')
      setWaitingTx(false)
      setToast({
        kind: 'error',
        message: 'The transaction was not completed.',
        title: 'Could not buy token',
      })
    }
    setWaitingTx(false)
  }

  const onClose = () => {
    setError(null)
    setIsSuccess(false)
    setWaitingTx(false)
    if (address) {
      mutate(['latestPot', address])
    }
  }

  if (!isMounted) {
    return null
  }
  const tix = useTix(totalPrice ?? 0)
  return (
    <Modal trigger={trigger}>
      <Dialog.Content className="fixed top-[50%] left-[50%] mt-10 w-[90vw] max-w-[500px] translate-x-[-50%] translate-y-[-50%] rounded-lg bg-white pb-4 shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] focus:outline-none data-[state=open]:animate-contentShow">
        <div className="flex flex-row justify-between rounded bg-[#F8F8F8] p-4">
          {' '}
          <Dialog.Title>
            <h2 className="text-md m-0 items-center justify-center font-semibold text-gray-900">
              {isSuccess
                ? 'Your purchase has been processed!'
                : isApproved
                ? 'Finalizing on the blockchain'
                : 'Complete Checkout'}
            </h2>
          </Dialog.Title>
          <Dialog.Close asChild>
            <button
              onClick={onClose}
              className="inline-flex h-[25px] w-[25px] items-center justify-center text-gray-600 focus:outline-none"
              aria-label="Close"
            >
              <HiX />
            </button>
          </Dialog.Close>
        </div>

        {priceLoading ? (
          <div className="flex h-40 items-center justify-center">
            <CgSpinner className="mr-3 h-6 w-6 animate-spin" />
            <p>Loading...</p>
          </div>
        ) : (
          <div className="m-2 flex min-h-[200px] flex-row md:flex-grow md:flex-col">
            {isSuccess ? (
              <div className="relative mt-4 flex flex-col items-center justify-center gap-2">
                <div className="absolute inset-0 z-10 mt-6 flex scale-125 transform items-center justify-center">
                  <img src="/success.gif" className="object-cover" />
                </div>
                <HiCheckCircle className=" h-[80px] w-[80px] items-center justify-center text-green-700" />
                <h1 className="text-xl font-semibold">
                  NFT Purchase Successful! 🎉
                </h1>
                <div className="text-sm font-light text-gray-500 ">
                  Your NFT has been sent to your wallet
                </div>
              </div>
            ) : (
              <>
                <div className="flex flex-col gap-1 p-2 ">
                  {error && (
                    <div className="flex flex-row items-center justify-center gap-2 rounded-sm border bg-gray-100 px-5 py-2 text-xs font-light text-gray-700">
                      <HiExclamationCircle className="h-4 w-4 text-red-500" />
                      {error}
                    </div>
                  )}
                  <div className="mb-2 text-sm text-gray-500">Item</div>
                  <div className="flex flex-row justify-between">
                    <div className="flex flex-row items-center justify-center gap-2">
                      {cartTokens.length > 1 ? (
                        <>
                          <img
                            src={cartTokens[0]?.token.collection?.image}
                            className="w-[50px] rounded-sm object-fill"
                          />
                          <div className="flex flex-col justify-center">
                            <p className="truncate text-sm">
                              {cartCount} items
                            </p>
                          </div>
                        </>
                      ) : (
                        <>
                          <img
                            src={cartTokens[0]?.token?.image}
                            className="w-[50px] rounded-sm object-fill"
                          />
                          <div className="flex flex-col justify-center">
                            <h1 className="truncate font-semibold">
                              {cartTokens[0]?.token?.tokenId}
                            </h1>
                            <p className="truncate text-sm">
                              {cartTokens[0]?.token?.collection?.name}
                            </p>
                          </div>
                        </>
                      )}
                    </div>
                    <div className="flex flex-row items-center justify-center gap-1 ">
                      {tix > 0 && (
                        <div className=" z-10 rounded border border-[#0FA46E] bg-[#DBF1E4] px-2 text-sm font-normal text-[#0FA46E]">
                          +{tix} TIX
                        </div>
                      )}
                      <img src="/eth.svg" alt="eth" className="h-4 w-4" />
                      <div className="text-sm font-semibold">{totalPrice}</div>
                    </div>
                  </div>
                </div>
                <div className="my-2 border-t"></div>
                <main className="mt-4 flex flex-col justify-between px-4">
                  <div className="flex flex-row items-center justify-between">
                    <div className="text-base font-medium text-gray-900">
                      Total
                    </div>

                    <div className="flex flex-row items-center justify-center gap-1">
                      <img src="/eth.svg" className="h-5 w-5" />
                      <div className="text-md font-semibold text-black">
                        {totalPrice}
                      </div>
                    </div>
                  </div>

                  <div className="mt-[25px] flex justify-end gap-4">
                    <button
                      onClick={handleSubmit}
                      disabled={isLoading}
                      className="w-full rounded bg-[#7000FF] py-2 text-white hover:bg-[#430099]"
                    >
                      {isLoading ? (
                        <>
                          <CgSpinner className="mr-2 inline-block h-6 w-6 animate-spin" />
                          {isApproved
                            ? 'Waiting to be Validated'
                            : 'Waiting for Approval'}
                        </>
                      ) : error ? (
                        'Retry'
                      ) : (
                        'Checkout'
                      )}
                    </button>
                  </div>
                </main>
              </>
            )}
          </div>
        )}
      </Dialog.Content>
    </Modal>
  )
}

export default BuyCartModal
