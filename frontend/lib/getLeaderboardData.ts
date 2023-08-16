import { Address } from "wagmi";
import fetcher from "./fetcher";

export interface GetLeaderboardDataRequest {
  potId: number;
  chain: "goerli"
}

export interface GetLeaderboardDataResponse extends Array<{
  num_of_tickets: number,
  wallet_address: Address,
  pot_id: number
}> { }

const getLeaderboardData = async ({ potId, chain }: GetLeaderboardDataRequest)
  : Promise<GetLeaderboardDataResponse | undefined> => {
  try {
    const response: GetLeaderboardDataResponse =
      await fetcher(`https://api.hotpot.gg/pot/${potId}/leaderboard?chain=${chain}`);

    return response;
  }
  catch (e) {
    console.error("Error from getLeaderboardData: ", e);
    return;
  }
}

export default getLeaderboardData;