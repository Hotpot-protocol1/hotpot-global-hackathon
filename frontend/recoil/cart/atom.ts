import { useTokens } from '@reservoir0x/reservoir-kit-ui'
import { atom } from 'recoil'

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

export default atom<Token[]>({
  key: 'cartTokens',
  default: [],
})
