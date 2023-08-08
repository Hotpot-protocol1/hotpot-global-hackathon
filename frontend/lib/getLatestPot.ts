export interface PotData {
  NumOfTickets: number
  wallet_address: string
  pot_id: number
  tickets: { ticket_id: number; is_winner: boolean }[]
}

export interface Pots {
  pot_id: number
  raffle_timestamp: string | null
}

export const getLatestPot = async (
  user: string
): Promise<{ currentPot: PotData | null; pots: Pots[] | null }> => {
  try {
    const currentPotPromise = fetch(
      `https://api.metalistings.xyz/user/${user}/pot/current`
    ).then((res) => res.json())
    const potsPromise = fetch(
      `https://api.metalistings.xyz/user/${user}/pot`
    ).then((res) => res.json())

    const [currentPot, pots] = await Promise.all([
      currentPotPromise,
      potsPromise,
    ])

    return {
      currentPot,
      pots,
    }
  } catch (error) {
    console.error('Error fetching data:', error)
    return {
      currentPot: null,
      pots: null,
    }
  }
}
