import { useUsersActivity } from '@reservoir0x/reservoir-kit-ui'
import ActivityTable from 'components/tables/ActivityTable'
import { FC, useEffect, useState } from 'react'
import { Item } from '../../lib/getAllListedNFTs'

type Props = {
  user?: string
  listedNFTs: Item[] | null
}

type ActivityQuery = NonNullable<
  Exclude<Parameters<typeof useUsersActivity>['1'], boolean>
>
type ActivityTypes = Exclude<ActivityQuery['types'], string>

const UserActivityTab: FC<Props> = ({ user, listedNFTs }) => {
  const [activityTypes, setActivityTypes] = useState<ActivityTypes>([])
  const query: ActivityQuery = {
    limit: 20,
    types: activityTypes,
  }
  const data = useUsersActivity(user ? [user] : undefined, query, {
    revalidateOnMount: false,
    fallbackData: [],
    revalidateFirstPage: true,
  })

  useEffect(() => {
    data.mutate()
    return () => {
      data.setSize(1)
    }
  }, [])

  return (
    <ActivityTable
      data={data}
      listedNFTs={listedNFTs}
      types={activityTypes}
      onTypesChange={(types) => {
        setActivityTypes(types)
      }}
      emptyPlaceholder={
        <div className="reservoir-body mt-14 grid justify-center dark:text-white">
          There hasn&apos;t been any activity yet.
        </div>
      }
    />
  )
}

export default UserActivityTab
