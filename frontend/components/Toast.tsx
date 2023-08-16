import React, { FC, Fragment } from 'react'
import { Transition } from '@headlessui/react'
import {
  HiCheckCircle,
  HiOutlineCheckCircle,
  HiOutlineExclamationCircle,
  HiOutlineInformationCircle,
  HiOutlineXCircle,
  HiStar,
  HiX,
} from 'react-icons/hi'
import toast, { Toast } from 'react-hot-toast'

type Props = {
  t: Toast
  toast: typeof toast
  data: {
    kind: 'error' | 'success' | 'warning' | 'info' | 'tickets' | 'complete'
    title: string
    message: string
  }
}

const toastStyles = {
  error: {
    bgColor: 'bg-white',
    borderColor: 'ring-black',
    textColor: 'text-gray-900',
    messageColor: 'text-gray-500',
  },
  success: {
    bgColor: 'bg-white',
    borderColor: 'ring-green-500',
    textColor: 'text-gray-900',
    messageColor: 'text-gray-500',
  },
  warning: {
    bgColor: 'bg-white',
    borderColor: 'ring-yellow-500',
    textColor: 'text-gray-900',
    messageColor: 'text-gray-500',
  },
  info: {
    bgColor: 'bg-white',
    borderColor: 'ring-blue-500',
    textColor: 'text-gray-900',
    messageColor: 'text-gray-600',
  },
  tickets: {
    bgColor: 'bg-[#FFF3C9]',
    borderColor: 'ring-[#FFC700;]',
    textColor: 'text-[#FF991C]',
    messageColor: 'text-white',
  },
  complete: {
    bgColor: 'bg-[#B5E3E3]',
    borderColor: 'ring-[#0FA4A4]',
    textColor: 'text-[#0C8383]',
    messageColor: 'text-white',
  },
}

const Toast: FC<Props> = ({ t, toast, data: { kind, message, title } }) => {
  const { bgColor, borderColor, textColor, messageColor } =
    toastStyles[kind] || {}
  return (
    <div
      className={`mx-4 flex w-full max-w-sm flex-col items-center space-y-4 rounded-lg sm:items-end ${bgColor}`}
    >
      <Transition
        show={t.visible}
        as={Fragment}
        enter="transform ease-out duration-300 transition"
        enterFrom="translate-y-4 opacity-0 sm:translate-y-0 sm:translate-x-2"
        enterTo="translate-y-0 opacity-100 sm:translate-x-0"
        leave="transition ease-in duration-100"
        leaveFrom="opacity-100"
        leaveTo="opacity-0"
      >
        <div
          className={`pointer-events-auto w-full max-w-sm overflow-hidden rounded-lg shadow-lg ring-1 ${borderColor} ring-opacity-5 dark:bg-black dark:ring-neutral-600 ${bgColor}`}
        >
          <div className="p-4">
            <div className="flex items-start">
              <div className="flex-shrink-0">{icons[kind]}</div>
              <div className="ml-3 w-0 flex-1">
                <p className={`reservoir-p font-medium ${textColor} `}>
                  {title}
                </p>
                <p className={`reservoir-p mt-1 ${messageColor}`}>{message}</p>
              </div>
              <div className="ml-4 flex items-center">
                <button
                  className={`focus:offset-2 inline-flex rounded-full bg-white bg-opacity-30 p-1 text-gray-400 hover:text-gray-500 focus:outline-none focus:outline-none `}
                  onClick={() => toast.dismiss(t.id)}
                >
                  <span className="sr-only">Close</span>
                  <HiX className="h-5 w-5" aria-hidden="true" />
                </button>
              </div>
            </div>
          </div>
        </div>
      </Transition>
    </div>
  )
}

export default Toast

const icons = {
  error: <HiOutlineXCircle className="h-6 w-6 rounded-full text-red-400" />,
  success: (
    <HiOutlineCheckCircle className="h-6 w-6 rounded-full text-green-400" />
  ),
  warning: (
    <HiOutlineExclamationCircle className="h-6 w-6 rounded-full text-yellow-400" />
  ),
  info: (
    <HiOutlineInformationCircle className="h-6 w-6 rounded-full text-blue-400" />
  ),
  tickets: <HiStar className="h-6 w-6 rounded-full text-[#FFA800]" />,
  complete: <HiCheckCircle className="h-6 w-6 rounded-full text-[#0FA4A4]" />,
}
