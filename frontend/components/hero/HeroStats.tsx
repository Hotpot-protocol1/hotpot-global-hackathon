import { FC, ReactNode } from 'react'
import FormatNativeCrypto from 'components/FormatNativeCrypto'
import { formatNumber } from 'lib/numbers'
import FormatCrypto from 'components/FormatCrypto'
import { useCollections } from '@reservoir0x/reservoir-kit-ui'

const API_BASE =
  process.env.NEXT_PUBLIC_RESERVOIR_API_BASE || 'https://api.reservoir.tools'

type Currency = NonNullable<
  NonNullable<
    NonNullable<ReturnType<typeof useCollections>['data']>[0]['topBid']
  >['price']
>['currency']

type Props = {
  count: number
  floor: number | undefined
  allTime: number | undefined
  volumeChange: number | undefined
  floorChange: number | undefined
}

const HeroStats: FC<{ stats: Props }> = ({ stats }) => {
  return (
    <div className="grid grid-cols-3 items-center gap-2 overflow-hidden rounded-lg dark:bg-[#525252] md:m-0 md:h-[82px] md:gap-20 md:bg-white dark:md:bg-black">
      <Stat name="Items">
        <h3 className="text-base font-semibold dark:text-white">
          {formatNumber(stats.count)}
        </h3>
      </Stat>
      <Stat name="Floor price">
        <h3 className="flex items-center justify-center gap-1 text-base dark:text-white">
          <FormatNativeCrypto amount={stats.floor} maximumFractionDigits={2} />
          {/*<PercentageChange value={stats.floorChange} />*/}
        </h3>
      </Stat>
      <Stat name="Total Vol">
        <h3 className="flex items-center justify-center gap-1 text-base dark:text-white">
          <FormatNativeCrypto
            amount={stats.allTime}
            maximumFractionDigits={2}
          />
        </h3>
      </Stat>
    </div>
  )
}

const Stat: FC<{ name: string; children: ReactNode }> = ({
  name,
  children,
}) => (
  <div className="flex flex-col items-center justify-center bg-white dark:bg-black md:h-auto">
    <p className="mb-1 text-sm font-normal text-[#A3A3A3]">{name}</p>
    <div className="mt-1 text-base font-medium">{children}</div>
  </div>
)

export const PercentageChange: FC<{ value: number | undefined | null }> = ({
  value,
}) => {
  if (value === undefined || value === null) return null

  const percentage = (value - 1) * 100

  if (percentage > 100 || value === 0) {
    return null
  }

  if (value < 1) {
    return (
      <div className="text-sm text-[#FF3B3B]">{formatNumber(percentage)}%</div>
    )
  }

  if (value > 1) {
    return (
      <div className="text-sm text-[#06C270]">+{formatNumber(percentage)}%</div>
    )
  }

  return null
}

export default HeroStats
