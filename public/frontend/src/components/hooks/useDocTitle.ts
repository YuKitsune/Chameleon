import { useEffect, useState } from 'react'

const useDocTitle = (title: string) => {
  const [doctitle, setDocTitle] = useState(title)

  useEffect(() => {
    document.title = doctitle
  }, [doctitle])

  return [doctitle, setDocTitle]
}

export default useDocTitle
