import useEnvChain from 'hooks/useEnvChain'
import Link from 'next/link'
import { FC } from 'react'

type Props = {
  variant?: 'desktop' | 'mobile' | undefined
  className?: string
}

const NavbarLogo: FC<Props> = ({ variant, className }) => {
  const logo = '/hotpot.png'
  const desktopLogo = '/hotpot-desktop.png'
  const logoAlt = 'Logo'

  const mobileVariant = variant === 'mobile'
  const desktopVariant = variant === 'desktop'

  return (
    <Link href="/" legacyBehavior={true}>
      <a
        className={`relative inline-flex flex-none items-center gap-1 ${className}`}
      >
        <img
          src={logo}
          alt={logoAlt}
          className={`h-9 w-auto ${!variant ? 'md:hidden' : ''} ${
            desktopVariant ? 'hidden' : ''
          } ${mobileVariant ? 'block' : ''}`}
        />
        <img
          src={desktopLogo}
          alt={logoAlt}
          className={`h-9 w-auto md:block ${
            !variant ? 'hidden md:block' : ''
          } ${mobileVariant ? 'hidden' : ''} ${desktopVariant ? 'block' : ''}`}
        />
      </a>
    </Link>
  )
}

export default NavbarLogo
