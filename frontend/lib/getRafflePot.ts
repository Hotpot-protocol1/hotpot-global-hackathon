export interface PotData {
  NumOfTickets: number
  wallet_address: string
  pot_id: number
  tickets: { ticket_id: number; is_winner: boolean }[]
}

export const getRafflePot = async (user: string): Promise<PotData | null> => {
  try {
    const latestRaffleResponse = await fetch(
      'http://3.1.203.189/pot/latest_raffle'
    )
    const potIdResponse = await latestRaffleResponse.json()
    const potId = potIdResponse.pot_id

    const userPotResponse = await fetch(
      `http://3.1.203.189/user/${user}/pot/${potId}`
    )
    const potDetails = await userPotResponse.json()

    return potDetails
  } catch (error) {
    console.error('Error fetching pot data:', error)
    return null
  }
}
