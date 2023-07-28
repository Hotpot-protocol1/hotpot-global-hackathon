import React, { Dispatch, SetStateAction, useEffect, useState } from 'react'
import { ethers } from 'ethers'
import { useContract, useProvider, useSigner } from 'wagmi'
import * as Dialog from '@radix-ui/react-dialog'
import { CgSpinner } from 'react-icons/cg'
import { HiX } from 'react-icons/hi'
import { TokenDetails } from 'types/reservoir'
import Modal from './Modal'
import { abi, NFTMarketplace_CONTRACT_SEP } from '../../contracts/index'
import getTotalPrice from 'lib/getTotalPrice'

type BuyCallbackData = {
  tokenId?: string
  collectionId?: string
}

type Props = Pick<Parameters<typeof Modal>['0'], 'trigger'> & {
  openState?: [boolean, Dispatch<SetStateAction<boolean>>]
  itemId?: number
  price?: string
  totalPrice?: number
  tokenDetails?: TokenDetails
  onGoToToken?: () => any
  onBuyComplete?: (data: BuyCallbackData) => void
  onBuyError?: (error: Error, data: BuyCallbackData) => void
  onClose?: () => void
}

const BuyModal: React.FC<Props> = ({
  trigger,
  itemId,
  price,
  tokenDetails,
}) => {
  const [isLoading, setIsLoading] = useState(false)
  const provider = useProvider()
  const { data: signer } = useSigner()
  const [isMounted, setIsMounted] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)
  const [priceValue, setPriceValue] = useState<number>(0)
  const [toast, setToast] = useState(null)

  const totalPrice = itemId ? getTotalPrice(itemId) : null

  useEffect(() => {
    setIsMounted(true)
  }, [])

  const NftMarketplace = useContract({
    address: NFTMarketplace_CONTRACT_SEP,
    abi: abi,
    signerOrProvider: signer,
  })

  const handleSubmit = async () => {
    setIsLoading(true)
    setError(null)

    try {
      if (!NftMarketplace) {
        console.log('NftMarketplace contract instance is not available.')
        return
      }
      if (!totalPrice) {
        console.log('Wait for total price to load')
        return
      }
      const priceInWei = ethers.utils.parseEther(totalPrice)
      const buyNFT = await NftMarketplace.purchaseItem(itemId, {
        value: priceInWei,
      })
      setIsLoading(false)
      console.log('Listing Transaction Hash:', buyNFT.hash)
    } catch (error) {
      setIsLoading(false)
      console.log(error)
      setError('Oops, something went wrong. Please try again')
    }
  }

  if (!isMounted) {
    return null
  }

  return (
    <Modal trigger={trigger}>
      <Dialog.Content className="fixed top-[50%] left-[50%] mt-10 w-[90vw] max-w-[500px] translate-x-[-50%] translate-y-[-50%] rounded-lg bg-white pb-4 shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] focus:outline-none data-[state=open]:animate-contentShow">
        <div className="flex flex-row justify-between rounded bg-[#F8F8F8] p-4">
          {' '}
          <Dialog.Title>
            <h2 className="m-0 font-semibold text-gray-900 text-md">
              Complete Checkout
            </h2>
          </Dialog.Title>
          <Dialog.Close asChild>
            <button
              className="inline-flex h-[25px] w-[25px] items-center justify-center text-gray-600 focus:outline-none"
              aria-label="Close"
            >
              <HiX />
            </button>
          </Dialog.Close>
        </div>
        <div className="m-2 flex min-h-[200px] flex-row md:flex-grow md:flex-col">
          <div className="flex flex-col gap-1 p-2">
            <div className="mb-2 text-sm text-gray-500">Item</div>
            <div className="flex flex-row justify-between">
              <div className="flex flex-row items-center justify-center gap-2">
                {' '}
                <img
                  src="/hotpot.png"
                  className="w-[50px] rounded-sm object-fill"
                />
                <div className="flex flex-col justify-center">
                  <h1 className="font-semibold truncate">NFT Name</h1>
                  <p className="text-sm text-gray-500 truncate">Collection</p>
                </div>
              </div>
              <div className="flex flex-row items-center justify-center gap-1">
                <img src="/eth.svg" alt="eth" className="w-4 h-4" />
                <div className="text-sm font-semibold">{price}</div>
              </div>
            </div>
          </div>
          <main className="flex flex-col justify-between px-4 mt-8">
            <div className="flex flex-row items-center justify-between">
              <div className="text-base font-medium text-gray-900">Total</div>

              <div className="flex flex-row items-center justify-center gap-1">
                <img src="/eth.svg" className="w-5 h-5" />
                <div className="font-semibold text-black text-md">
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
                    <CgSpinner className="inline-block w-6 h-6 mr-2 animate-spin" />
                    Waiting for approval
                  </>
                ) : error ? (
                  'Retry'
                ) : (
                  'Checkout'
                )}
              </button>
            </div>
          </main>
        </div>
      </Dialog.Content>
    </Modal>
  )
}

export default BuyModal
