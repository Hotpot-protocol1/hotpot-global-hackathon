import { Chain } from 'wagmi'

const apothem_rpcs: string[] = ['https://apothem.xdcrpc.com/'];
const apothemChain: Chain = {
	id: 51,
	name: 'Apothem',
	network: 'apothem',
	nativeCurrency: {
			decimals: 18,
			name: 'TXDC',
			symbol: 'TXDC',
	},
	rpcUrls: {
			public: { http: apothem_rpcs },
			default: { http: apothem_rpcs },
	},
	blockExplorers: {
			etherscan: { name: 'Blocksscan', url: 'https://explorer.apothem.network/' },
			default: { name: 'Blocksscan', url: 'https://explorer.apothem.network/' },
	},
	testnet: true,
};
export {apothemChain};
