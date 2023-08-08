export interface PotData {
  NumOfTickets: number
  wallet_address: string
  pot_id: number
  tickets: { ticket_id: number; is_winner: boolean }[]
}

export const getPotById = async (
  user: string,
  potId: number
): Promise<PotData | null> => {
  try {
    const userPotResponse = await fetch(
      `https://api.metalistings.xyz/user/${user}/pot/${potId}`
    )
    const potDetails = await userPotResponse.json()
    return potDetails
  } catch (error) {
    console.error('Error fetching pot data:', error)
    return null
  }
}
