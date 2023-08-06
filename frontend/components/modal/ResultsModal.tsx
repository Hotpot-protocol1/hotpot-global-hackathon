import React, { useState, useEffect, Dispatch, SetStateAction } from 'react'
import { ethers } from 'ethers'
import { useContract, useProvider, useSigner } from 'wagmi'
import * as Dialog from '@radix-ui/react-dialog'
import { CgSpinner } from 'react-icons/cg'
import { HiCheckCircle, HiExclamationCircle, HiX } from 'react-icons/hi'
import { TokenDetails } from 'types/reservoir'
import Modal from './Modal'
import { Hotpot_CONTRACT_SEP, hotpotAbi } from '../../contracts/index'
import getTotalPrice from 'lib/getTotalPrice'
import useTix from 'lib/tix'
import { optimizeImage } from 'lib/optmizeImage'
import Image from 'next/legacy/image'
import { useMediaQuery } from '@react-hookz/web'

type ClaimCallbackData = {
  ticketId?: string
}

type Props = Pick<Parameters<typeof Modal>['0'], 'trigger'> & {
  openState?: [boolean, Dispatch<SetStateAction<boolean>>]
  ticketId?: number
  amount?: string
  loading?: boolean
  onGoToToken?: () => any
  onClaimComplete?: (data: ClaimCallbackData) => void
  onClaimError?: (error: Error, data: ClaimCallbackData) => void
  onClose?: () => void
}

const ResultsModal: React.FC<Props> = ({ trigger, ticketId }) => {
  const [isLoading, setIsLoading] = useState(false)
  const [priceLoading, setPriceLoading] = useState(false)
  const provider = useProvider()
  const { data: signer } = useSigner()
  const [isMounted, setIsMounted] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)
  const [toast, setToast] = useState(null)
  const [isSuccess, setIsSuccess] = useState<boolean>(false)
  const [totalPrice, setTotalPrice] = useState<string | null>(null)
  const singleColumnBreakpoint = useMediaQuery('(max-width: 640px)')
  const imageSize = singleColumnBreakpoint ? 533 : 250

  useEffect(() => {
    setIsMounted(true)
  }, [])

  const Hotpot = useContract({
    address: Hotpot_CONTRACT_SEP,
    abi: hotpotAbi,
    signerOrProvider: signer,
  })

  const handleSubmit = async () => {
    setIsLoading(true)
    setError(null)

    try {
      if (!Hotpot) {
        console.log('NftMarketplace contract instance is not available.')
        return
      }
      if (!totalPrice) {
        console.log('Wait for total price to load')
        return
      }
      const claimPrize = await Hotpot.claim
      setIsLoading(false)
      console.log('Listing Transaction Hash:', claimPrize.hash)
      setIsSuccess(true)
    } catch (error) {
      setIsLoading(false)
      console.log(error)
      setError('Oops, something went wrong. Please try again')
    }
  }

  const onClose = () => {
    setError(null)
    setIsSuccess(false)
  }

  if (!isMounted) {
    return null
  }

  let win = (
    <div className="rounded-xl bg-gradient-to-r from-[#FEF0D6] to-[#FFECAC] pb-4">
      <div className="flex flex-row justify-between">
        {' '}
        <Dialog.Title></Dialog.Title>
        {/* <Dialog.Close asChild>
          <button
            onClick={onClose}
            className="m-3 inline-flex h-[20px] w-[20px] items-center justify-center text-gray-600 focus:outline-none"
            aria-label="Close"
          >
            <HiX />
          </button>
        </Dialog.Close> */}
      </div>
      <div className="mx-4 mt-6 flex flex-col items-center justify-center">
        <h1 className="reservoir-h1 mb-2 text-[40px] font-semibold text-[#101828]">
          Congratulations!
        </h1>
        <p className="reservoir-subtitle font-medium">
          You won 5 ETH with Golden Ticket #2354
        </p>
        <img src="/gold-chest.svg" className="my-4 w-[260px]" />
        <button className="mb-4 rounded-full border border-[#FFF06A] bg-gradient-to-b from-[#FFE179] to-[#FFB52E] px-16  py-3 text-sm font-medium text-[#CD7100] hover:from-[#FFC269] hover:to-[#FFB82E] focus:outline-none">
          CLAIM
        </button>
      </div>
    </div>
  )

  let result = (
    <div className="rounded-xl bg-[#FFF5F5] pb-4">
      <div className="flex flex-row justify-between ">
        {' '}
        <Dialog.Title></Dialog.Title>
      </div>
      <div className="mx-4 mt-6 flex flex-col items-center justify-center">
        <h1 className="reservoir-h1 mx-2 mb-2 text-center text-[40px] font-semibold text-[#101828]">
          Sorry, your ticket(s) did not win
        </h1>
        <p className="reservoir-subtitle font-medium">
          Don't sweat! Try your luck on the next draw!
        </p>
        <img src="/sad.svg" className="w-[200px]" />

        <Dialog.Close asChild>
          <button
            onClick={onClose}
            aria-label="Close"
            className="mb-4 rounded-full bg-[#6A3CF5] px-16 py-3 text-sm font-medium text-white hover:bg-[#7C4CF5]"
          >
            Close
          </button>
        </Dialog.Close>
      </div>
    </div>
  )
  return (
    <Modal trigger={trigger}>
      <Dialog.Content className=" rounded-4xl fixed top-[50%] left-[50%] mt-10 w-[90vw] max-w-[500px] translate-x-[-50%] translate-y-[-50%] rounded shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] backdrop-blur-md focus:outline-none data-[state=open]:animate-contentShow">
        {win}
      </Dialog.Content>
    </Modal>
  )
}

export default ResultsModal
