import React from 'react'

interface HeroProps {
  variant?: string
}

const Hero: React.FC<HeroProps> = ({ variant }) => {
  const backgroundImageUrl =
    variant === 'rewards' ? '/banner-rewards.svg' : '/banner-home.svg'
  const bottomImageUrl =
    variant === 'rewards' ? '/gold-pot.svg' : '/gold-chest.svg'
  const textColor = variant === 'rewards' ? 'text-white' : 'text-[#101828]'

  return (
    <div
      className="grid w-full grid-cols-1 overflow-hidden rounded-lg bg-cover bg-center md:max-h-[428px] md:grid-cols-2"
      style={{ backgroundImage: `url('${backgroundImageUrl}')` }}
    >
      <div>
        <div className="flex flex-col items-center justify-center p-12">
          <img src="/pot-o-gold.svg" alt="pot-o-gold" />
          <h2
            className={`mt-[-12px] text-base font-normal ${textColor} md:text-xl`}
          >
            Earn 1 raffle ticket for every 0.10 ETH bought or sold
          </h2>
          <div className="mt-11 flex flex-col items-center justify-center gap-2 rounded-2xl border-2 border-solid border-[#FFF06A] bg-gradient-to-b from-[#FFE179] to-[#FFB52E] px-10 py-4 text-black md:px-16">
            <div className="text-lg font-normal md:text-xl">
              The Prize Pool:
            </div>
            <div className="text-2xl md:text-3xl">
              27.1 ETH / <span className="text-purple-600">100 ETH</span>
            </div>
          </div>
        </div>
      </div>

      <div className="hidden flex-row items-end justify-center md:flex">
        <img
          src={bottomImageUrl}
          alt="rewards-image"
          className="h-[500px] object-cover object-bottom"
        />
      </div>
    </div>
  )
}

export default Hero
