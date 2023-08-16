import React from 'react'
import { FiCopy } from 'react-icons/fi'

interface CopyClipboardProps {
  content: string
}

const CopyClipboard: React.FC<CopyClipboardProps> = ({ content }) => {
  return (
    <FiCopy
      className="h-4 w-4 hover:cursor-pointer hover:text-gray-800 focus:text-gray-800 active:translate-y-1"
      onClick={() => navigator.clipboard.writeText(content)}
    />
  )
}

export default CopyClipboard
