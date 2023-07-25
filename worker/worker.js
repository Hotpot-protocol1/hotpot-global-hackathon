// alchemy setup
// import { Network, Alchemy } from "alchemy-sdk";
//commonjs module setup of alchemy
const { Network, Alchemy } = require("alchemy-sdk");

const settings = {
    apiKey: process.env.ALCHEMY_API_KEY, // Replace with your Alchemy API Key.
    network: Network.ETH_SEPOLIA, // Replace with your network.
};

const alchemy = new Alchemy(settings);
// end alchemy setup
const Web3 = require('web3');
require('dotenv').config();
console.log('hotpot contract: ' + process.env.HOTPOT_CONTRACT_ADDRESS);

const marketplaceContractAddress = process.env.HOTPOT_CONTRACT_ADDRESS;
//
const eventName = 'GenerateRaffleTickets(address,address,uint32,uint32,uint32,uint32,uint256,uint256)';
const generateRaffleTicketsTopic = Web3.utils.keccak256(eventName);
console.log('generated raffletickets topic: ' + generateRaffleTicketsTopic);
//


const pgp = require('pg-promise')();
console.log(process.env.DATABASE_URL);
const db = pgp(process.env.DATABASE_URL);

// const web3 = new Web3(new Web3.providers.WebsocketProvider(process.env.INFURA_WS_URL));
// const contract = new web3.eth.Contract(ABI, CONTRACT_ADDRESS);


// Create the log options object.
const hotpotGenerateRaffleTicketsEvents = {
    address: marketplaceContractAddress,
    topics: [generateRaffleTicketsTopic],
};

// TODO: Add whatever logic you want to run upon mint events.
const doSomethingWithTxn = (txn) => console.log(txn);

// Open the websocket and listen for events!
alchemy.ws.on(hotpotGenerateRaffleTicketsEvents, doSomethingWithTxn);

// how to get event signature:
// const Web3 = require('web3');
// const web3 = new Web3(); // No provider necessary in this case

// const eventName = 'Approval(address,address,uint256)'; 
// const eventSignature = web3.utils.keccak256(eventName);
// console.log(eventSignature);
//

// contract.events.EventName({ fromBlock: 'latest' })
//     .on('data', async event => {
//         const query = 'INSERT INTO tablename(event_data_column) VALUES($1)';
//         try {
//             await db.none(query, [JSON.stringify(event.returnValues)]);
//             console.log('Event data saved!');
//         } catch (err) {
//             console.log(err);
//         }
//     })
//     .on('error', console.error);