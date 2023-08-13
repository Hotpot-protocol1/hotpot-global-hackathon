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
      <Select.Trigger className="btn-primary-outline ml-auto rounded-lg border-transparent p-0 normal-case outline-none">
        <div className="flex items-center gap-1 rounded-lg border border-solid border-[#CFD8E1] px-4 py-2">
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
        className="rounded-lg bg-gray-100 outline-none"
      >
        <Select.Group>
          {Object.entries(chains).map(([value, name]) => (
            <Select.Item key={value} value={value}>
              <div className="group flex w-full cursor-pointer items-center justify-between rounded px-6 py-3 outline-none transition hover:bg-gray-200 hover:outline-none focus:bg-neutral-100 dark:hover:bg-neutral-700 dark:focus:bg-neutral-800">
                {name}
              </div>
            </Select.Item>
          ))}
        </Select.Group>
      </Select.Content>
    </Select.Root>
  )
}

export default SelectChain
