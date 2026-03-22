import  { useState } from 'react'

export const StatusButton = () => {
const [status, setSattus] = useState('')

    const getStatus = () => {
        fetch('/api/get').then((data)=>
           data.text()
        ).then(data => setSattus(data))
    }

  return (<>
    <button onClick={getStatus}>statusButton</button>
    <p>{status}</p>
  </>
  )
}
 