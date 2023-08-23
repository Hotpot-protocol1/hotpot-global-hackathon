# Marketplace Deployment Guide

This guide explains how to clone and deploy the Marketplace project.

## Prerequisites

Before you begin, make sure you have the following tools installed:

- Git
- Node.js and npm

## Clone the Repository

1. Open your terminal.

2. Change the current working directory to where you want to clone the project.

3. Run the following command to clone the repository:

   ```bash
   git clone https://github.com/your-username/your-marketplace-repo.git
   ```

## Install Dependencies

1. Change your working directory to the cloned project's directory:

   ```bash
   cd your-marketplace-repo
   ```

2. Install project dependencies using npm:

   ```bash
   npm install
   ```

## Start the Application

1. After installing dependencies, you can start the application:

   ```bash
   npm start
   ```

The application will start, and you can access it in your web browser at [http://localhost:3000](http://localhost:3000).

## Additional Configuration

You may need to configure environment variables or other settings based on your deployment environment. Refer to the project's documentation for more details.

## Project Description

<br>**Background**</br>
Hotpot brings a new dimension of engagement for web3 dapps. It transforms any transaction producing activity into a lottery game. With every trade, transaction, swap, users are awarded with a raffle ticket. A % of the transaction fee is sent to the raffle pot. When the pot is full, random winners will be chosen on-chain and can claim the ultimate reward. It also allows any one create their own lottery game on any DApp. 
</br>
<br>
**How it works**
<br> 
We’ve developed a custom cross chain NFT marketplace using Axelar, deployed contracts on Avalanche Fuji, Base Georli, XDC, ZkSync and ETH Goerli. The raffle contract is deployed on Base which manages generating raffle tickets, selecting the winners, holding the jackpot, and allowing the winners to claim the prize. 
</br>
<br>
A) Alice lists an NFT on Avalanche, Bob buys it. A transaction is sent through axelar via a cross chain msg & token function to call ‘executeTrade’  on the raffle contract deployed on Base. This custom function generates onchain events that have a range of ticketIDs (We use a range in order to optimize for gas, we wouldn’t want users to pay for gas for every single ticket!). 
</br>
<br>
B) 1% of the trading fees is also sent to the raffle contract, growing the jackpot. 
</br>
<br>
C) When the pot is full based on a predetermined parameter, the executeRaffle function in the raffle contract will trigger, calling a function called fulfillRandomness. Because chainlink VRF isn’t available on Base or XDC, we’ve developed our own VRF solution by hashing a salt with a timestamp of the current block. This returns randomnessFulfilled and we use that to derive the winning tickets numbers to ensure provable fairness. 
</br>
<br>
D) Once the winner has been determined, the raffle contract will make an attestation through a custom schema containing the address of the winner, the raffle address, the size of the prize, and a timestamp. 
</br>
<br>
E) The frontend lets users check if they’ve won, and will verify that the user addresses matches those with the winning tickets. ExecuteRaffle also starts a new pot, resetting all tickets. 
</br>
<br>
F) Because of Axelar, NFT trades that happen on Base, XDC, or any other EVM chains supported by Axelar can 1) send trading fees to the raffle 2) generate tickets 3) trigger the raffle contract, making this an interoperable, fast growing, multichain raffle system. 
</br>
<br>
G) Deploying a raffle through the factory requires certain parameters: 1) size of the pot 2) price per ticket 3) claim window etc… that can be set by anyone wanting to hold their own lottery. 
</br>

<br>**Axelar Bounty Requirement**</br>
Our 2 positive experiences with Axelar:
</br>
The documentation was very easy to follow.
It is easy to implement, easier than hyperlane. 
 <br>
Our 2 negative experiences:
</br>
The finalization time from Base to Avax was over 10 minutes long. While the finalization from Avax to Base was much faster.
</br>
The Axelar testnet explorer UI is cluttered and a bit hard to read. 
</br>

</br>

<br> **GH Repos** </br> 
- https://github.com/Hotpot-protocol1/base-hotpot
- https://github.com/Hotpot-protocol1/avax-hotpot
- https://github.com/Hotpot-protocol1/xdc-hotpot
- https://github.com/Hotpot-protocol1/goerli-hotpot
- https://github.com/Hotpot-protocol1/zksync-hotpot
- https://github.com/Hotpot-protocol1/hotpot-ethtoronto-contracts 
</br>
<br> **Deployments**</br>
<br>
*Base*
</br>
<br>
Hotpot Factory
https://goerli.basescan.org/address/0x6bccd1AB0F8f6278c19C3eF8C02934233c8b1D32
</br>
<br>
Deployed Hotpot
https://goerli.basescan.org/address/0x30b5db47421Fc8Db08d1d2a5CD2fC1437378f66b
</br>*
<br>
Hotpot Implementation
https://goerli.basescan.org/address/0x231F762d3B4DC761f5f9cd398ef6A843A76D8B9D
</br>
<br>
Marketplace (Base)
https://goerli.basescan.org/address/0x5AB54B696216deF02811B04700A4d23a8a0a3793
</br>
<br>
*Avalanche*
The cross-chain marketplace is deployed on Avax 
https://testnet.snowtrace.io/address/0xBE38c5eefE99a622cA5F36b6306E5C61470B3974
</br>
<br>
*Ethereum Attestation Service*
</br>
https://base-goerli.easscan.org/schema/view/0x2c5ca635f20506a5963e5cb00b3b77f99f40b42bee9a705a0b68d57346030405 
</br>
<br>
*Axelar*
https://testnet.axelarscan.io/address/0xE2172c474F2E95419754140d4ac19045A70Bc93F
</br>
<br>
*XDC (Apothem testnet)* </br>
<br>
Marketplace: xdcBE786989E194433365f34B45BdC9760c246c3f35 </br>
<br>
Hotpot implementation: xdc6554bc9bcBa2bB6E4F957A6425B9038C70E29085
</br>
<br>
Hotpot Factory: xdc4b5F283ecc3e609252caf1980D534Cf3779206a2
</br>
<br>
Hotpot: xdc4B20b152C338Fa96F639BBc170270f544523b307
</br>
<br>
Operator: xdcaE789C29DF5de1d2B36D7ACD695fcBdb55be85ab
</br>
</br><br>
*zkSync Era (testnet)*
</br>
<br>
Hotpot implementation: 0xC0b3DE5C1A951eA86F72Cf5AaD68bD63fD8d2BBD
</br><br>
Marketplace: 0x91bff59A8493f93DE8D996Cfe906f3a180f73506
</br><br>
Hotpot Factory: 0xcb5bEF15Ff527882726fef68168EA8Ac8b941E76
</br><br>
Hotpot: 0xc03Bfd96322FF03eD955AF8d6c71abFb209036b0
</br><br>
Hotpot (created without a factory): 0x2A3CeFa9D21d1e79b57b3d2858009F51722b041C
</br>
