import getLeaderboardData, { GetLeaderboardDataRequest } from "lib/getLeaderboardData";
import useSWR from "swr";

const useLeaderboardData = ({ potId, chain }: GetLeaderboardDataRequest) => {
  const { data, error, mutate } = useSWR(
    ["getLeaderboardData", potId, chain],
    () => getLeaderboardData({ potId, chain })
  );

  return { data, error, mutate };
}

export default useLeaderboardData;