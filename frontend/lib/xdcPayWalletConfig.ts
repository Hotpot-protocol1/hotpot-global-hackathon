import {
	Chain,
	Wallet,
	getWalletConnectConnector,
} from '@rainbow-me/rainbowkit';
import { InjectedConnector } from 'wagmi/connectors/injected';

  
export interface XdcWalletOptions {
	projectId: string;
	chains: Chain[];
}

export const xdcPayWalletConfig = ({
    chains,
    projectId,
}: XdcWalletOptions): Wallet => ({
	id: 'xdc-pay',
	name: 'XDC Pay',
	iconUrl: 'https://xinfin.org/assets/images/brand-assets/xdc-icon.png',
	iconBackground: '#f7f7f7',
	installed: true,
	createConnector() {
		const connector = new InjectedConnector({chains});

		return {
			connector,
		}
	},
});