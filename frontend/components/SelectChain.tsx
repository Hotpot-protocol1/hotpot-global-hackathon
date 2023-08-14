import { FC, useState } from 'react'
import useMounted from 'hooks/useMounted'
import { CgChevronDown, CgChevronUp } from 'react-icons/cg'
import * as Select from '@radix-ui/react-select'
import { useRouter } from 'next/router'

const SelectChain: FC = () => {
  const isMounted = useMounted()
  const router = useRouter()

  type Chains = 'mainnet' | 'goerli' | 'sepolia'

  const chains: Record<Chains, string> = {
    mainnet: 'Ethereum',
    goerli: 'Goerli',
    sepolia: 'Sepolia',
  }

  const links: Record<Chains, string> = {
    mainnet: 'https://hotpot-vercel.vercel.app',
    goerli: 'https://goerli-hotpot.vercel.app',
    sepolia: 'https://hotpot-vercel.vercel.app',
  }

  const logo: Record<Chains, string> = {
    mainnet: '/eth-logo.svg',
    goerli: '/eth.svg',
    sepolia: '/eth.svg',
  }

  const [selectedValue, setSelectedValue] = useState<Chains>('goerli')
  const [isOpen, setIsOpen] = useState(false)

  if (!isMounted) {
    return null
  }

  const handleValueChange = (value: Chains) => {
    setSelectedValue(value)
    router.push(links[value])
  }

  return (
    <Select.Root
      value={selectedValue}
      onValueChange={handleValueChange}
      onOpenChange={(open) => setIsOpen(open)}
    >
      <Select.Trigger className="btn-primary-outline ml-auto rounded-lg border-transparent p-0 normal-case">
        <div className="flex items-center gap-1 rounded-lg border border-solid border-[#CFD8E1] px-4 py-2">
          <img
            src={logo[selectedValue]}
            alt="Ethereum Logo"
            className="h-4 w-4"
          />
          <div className="text-base font-normal">{chains[selectedValue]}</div>
          {isOpen ? (
            <CgChevronUp className="text-gray-500" />
          ) : (
            <CgChevronDown className="text-gray-500" />
          )}
        </div>
      </Select.Trigger>

      <Select.Content
        position="popper"
        className="rounded-lg bg-white shadow-md"
      >
        <Select.Group>
          {Object.entries(chains).map(([value, name]) => (
            <Select.Item
              key={value}
              value={value}
              className="m-2 outline-none focus:outline-none"
            >
              <div className="group flex w-full cursor-pointer items-center gap-2 rounded-md px-6 py-3 text-sm transition hover:bg-gray-100 focus:bg-gray-100 focus:outline-none">
                <img
                  src={logo[value as Chains]}
                  alt="Ethereum Logo"
                  className="h-4 w-4"
                />
                <span className="text-gray-800 dark:text-white">{name}</span>
              </div>
            </Select.Item>
          ))}
        </Select.Group>
      </Select.Content>
    </Select.Root>
  )
}

export default SelectChain
