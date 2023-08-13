import { FaTwitter, FaDiscord } from 'react-icons/fa'
import Link from 'next/link'

const Footer = () => {
  return (
    <footer className="col-span-full flex flex-col items-center justify-between px-6 pb-4 sm:flex-row md:px-16">
      <div className="flex flex-row flex-wrap items-center justify-between gap-x-6  text-xs sm:mb-0 sm:gap-x-8 sm:text-sm">
        <Link href="/" legacyBehavior={true}>
          <img src="/hotpot.png" alt="hotpot-logo" className="h-8 w-8" />
        </Link>
      </div>
      <div className="py-2 text-center text-xs text-[#98A2B3]">
        All rights reserved ©2023
      </div>
      <div className="flex flex-row items-center gap-x-6">
        <Link href="https://discord.gg/" className="ml-5" legacyBehavior={true}>
          <a className="" target="_blank" rel="noreferrer">
            <FaDiscord className="h-[19px] w-[25px] text-[#845CFF]" />
          </a>
        </Link>
        <Link href="https://twitter.com/hotpot_gg" legacyBehavior={true}>
          <a className="" target="_blank" rel="noreferrer">
            <FaTwitter className="h-[20px] w-[25px] text-[#845CFF]" />
          </a>
        </Link>
      </div>
    </footer>
  )
}

export default Footer
