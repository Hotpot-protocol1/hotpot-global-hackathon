<!-- GETTING STARTED -->

## Getting Started (Self-Hosted)

### Prerequisites

1. Install [Node.js and NPM](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
2. Install [Yarn](https://classic.yarnpkg.com/en/docs/install)
3. Request free [Reservoir API key](https://api.reservoir.tools/#/0.%20Auth/postApikeys)

### Built With

- [ReservoirKit](https://docs.reservoir.tools/docs/reservoir-kit)
- [Reservoir Protocol and API](https://reservoirprotocol.github.io/)
- [Next.js](https://nextjs.org/)
- [React.js](https://reactjs.org/)
- [Ethers.io](https://ethers.io/)
- [WAGMI](https://wagmi.sh/)
- [Tailwind CSS](https://tailwindcss.com/)

### Installation

Fork this repo and follow these instructions to install dependancies.

With yarn:

```bash
$ yarn install
```

With NPM:

```bash
$ npm install
```

### Configuration

To preview your configuration locally you can copy the values you want to use from `env.development` or `env.production` into a new file called `.env.local`.

Note: Environment variables can also be added during deployment via deployment platforms like [vercel](https://vercel.com/).

**Required Configuration**
| Environment Variable | Required | Description | Example |
|--------------------------------|----------|-------------------------------------------------------------------------------------|---------------------|
| NEXT_PUBLIC_RESERVOIR_API_BASE | `true` | The Reservoir API base URL. Available on Mainnet, Rinkeby, Goerli, and Optimism. | https://api-rinkeby.reservoir.tools/ https://api.reservoir.tools/ |
| NEXT_PUBLIC_CHAIN_ID | `true` | The Ethereum network to be used. 1 for Etherem Mainnet and 4 for Rinkeby Testnet, etc. | 1 4 |
| NEXT_PUBLIC_PROXY_API_BASE | `true` | The proxy API used to pass the Reservoir API key without exposing it to the client. | /api/reservoir |
| NEXT_PUBLIC_RESERVOIR_API_KEY | `true` | Reservoir API key provided by the Reservoir Protocol. [Get your own API key](https://api.reservoir.tools/#/0.%20Auth/postApikeys). | 123e4567-e89b-12d3-a456-426614174000 |
| NEXT_PUBLIC_ALCHEMY_ID | `true` | Alchemy API key required for buying items on mobile. [Get your own API key here](https://docs.alchemy.com/alchemy/introduction/getting-started#1.create-an-alchemy-key). | 123e4567-e89b-12d3-a456-426614174000

**General Configuration**
| Environment Variable | Required | Description | Example |
| :---------------------------------- | :------- | :--------------------------------------------------------------------------------------------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| NEXT_PUBLIC_SOURCE_ID | `false` | **DEPRECATED**: The source ID used to attribute listings and offers from your marketplace. | Reservoir Market |
| NEXT_PUBLIC_SOURCE_DOMAIN | `false` | The source ID used to attribute listings and offers from your marketplace. Must be in a domain format. | reservoir.market |
| NEXT_PUBLIC_SOURCE_NAME | `false` | The source name used to attribute listings and offers from your marketplace, falls back to the domain. | Reservoir Market |
| NEXT_PUBLIC_ORDER_KIND | `false` | The order kind to be used when listing or making offers on your marketplace. Default will use `seaport` if not set. | zeroex-v4. wyvern-v2.3, seaport |
| NEXT_PUBLIC_FEE_BPS | `false` | The marketplace fee that will be earned from native listings and offers on your marketplace. Configured as <<glossary:bps>>. | 100 |
| NEXT_PUBLIC_FEE_RECIPIENT | `false` | The address that will receive marketplace fee. | 0xF296178d553C8Ec21A2fBD2c5dDa8CA9ac905A00 |
| NEXT_PUBLIC_COLLECTION | `false` | Use this to configure a single collection marketplace or set the default collection for your community marketplace. | 0x8d04a8c79ceb0889bdd12acdf3fa9d207ed3ff63 |
| NEXT_PUBLIC_COLLECTION_SET_ID | `false` | Use this to configure a community marketplace. [Generate your collection set ID here](https://docs.reservoir.tools/reference/postcollectionssetsv1). | f566ba09c14f56aedeed3f77e3ae7f5ff28b9177714d3827a87b7a182f8f90ff |
| NEXT_PUBLIC_COMMUNITY | `false` | Use this to configure a community marketplace. Note: Community IDs are only available for certain communities. | artblocks |
| NEXT_PUBLIC_REDIRECT_HOMEPAGE | `false` | If enabled, homepage will automatically redirect to collection page set in NEXT_PUBLIC_COLLECTION. | true |
| NEXT_PUBLIC_EXTERNAL_LINKS | `false` | External links to be displayed in the top navigation bar. | Blog::[\<\<\<\<https://blog.com,Docs::https://docs.com>>>>](https://blog.com,Docs::https://docs.com) |
| NEXT_PUBLIC_COLLECTION_DESCRIPTIONS | `false` | Customize descriptions on a per collection basis. | 0xb74bf94049d2c01f8805b8b15db0909168cabf46::`test description`,0xc751c84678d8e229e361f9b04c080256516f4a0a::`another description` |
| NEXT_PUBLIC_NAVBAR_LOGO_LINK | `false` | Customize the marketplace navbar logo's link. | <https://blog.com> |
| NEXT_PUBLIC_DEFAULT_TO_SEARCH | `false` | If enabled, search bar will be displayed instead of collection switcher. | true |
| NEXT_PUBLIC_LISTING_CURRENCIES | `false` | Customize which currency users can list in. | [{"symbol": "ETH", "decimals": 18, "contract": "0x0000000000000000000000000000000000000000"},{"symbol": "USDC", "decimals": 6, "contract": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"}] |

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.js`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/basic-features/font-optimization) to automatically optimize and load Inter, a custom Google Font.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.
