import React, { ReactNode } from 'react'
import * as Dialog from '@radix-ui/react-dialog'

interface ModalProps {
  trigger: ReactNode
  children: ReactNode
}

const Modal: React.FC<ModalProps> = ({ trigger, children }) => (
  <Dialog.Root>
    <Dialog.Trigger asChild>{trigger}</Dialog.Trigger>
    <Dialog.Portal>
      <Dialog.Overlay className="fixed inset-0 bg-blackA9 data-[state=open]:animate-overlayShow" />
      {children}
    </Dialog.Portal>
  </Dialog.Root>
)

export default Modal
