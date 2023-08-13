export interface PotData {
  NumOfTickets: number
  wallet_address: string
  pot_id: number
  tickets: { ticket_id: number; is_winner: boolean }[]
}

export const getRafflePot = async (user: string): Promise<PotData | null> => {
  try {
    const latestRaffleResponse = await fetch(
      'https://api.hotpot.gg/pot/latest_raffle?chain=goerli'
    )
    const potIdResponse = await latestRaffleResponse.json()
    const potId = potIdResponse.pot_id

    const userPotResponse = await fetch(
      `https://api.hotpot.gg/user/${user}/pot/${potId}?chain=goerli`
    )
    const potDetails: PotData = await userPotResponse.json()

    return potDetails
  } catch (error) {
    console.error('Error fetching pot data:', error)
    return null
  }
}
