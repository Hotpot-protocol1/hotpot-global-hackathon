import React, { useState, useEffect } from 'react'

type LoadingNumberProps = {
  loading: boolean
  value: string
}

const LoadingNumber: React.FC<LoadingNumberProps> = ({ loading, value }) => {
  const [currentValues, setCurrentValues] = useState<Array<string>>(
    value.split('')
  )
  const [animation, setAnimation] = useState(true)

  useEffect(() => {
    let intervals: NodeJS.Timeout[] = []

    if (loading && animation) {
      currentValues.forEach((_, index) => {
        if (currentValues[index] === '.' || isNaN(Number(currentValues[index])))
          return // Skip non-numeric characters

        intervals[index] = setInterval(() => {
          setCurrentValues((prevValues) => {
            const newValues = [...prevValues]
            newValues[index] = Math.floor(Math.random() * 10).toString()
            return newValues
          })
        }, 50) // Increase this value to slow down the animation
      })
    }

    return () => {
      intervals.forEach((interval) => {
        clearInterval(interval)
      })
    }
  }, [loading, animation, currentValues.length])

  useEffect(() => {
    if (!loading) {
      setTimeout(() => setAnimation(false), 1000)
    }
  }, [loading])

  useEffect(() => {
    if (!loading && !animation) {
      setCurrentValues(value.split(''))
    }
  }, [loading, value, animation])

  return (
    <>
      {currentValues.map((val, index) => (
        <span key={index}>{val}</span>
      ))}
    </>
  )
}

export default LoadingNumber
