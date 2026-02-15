import  { useState } from 'react'

export const StatusButton = () => {
const [status, setSattus] = useState('')

    const getStatus = () => {
        fetch('http://localhost:8080').then((data)=>
           data.text()
        ).then(data => setSattus(data))
    }

  return (<>
    <button onClick={getStatus}>statusButton</button>
    <p>{status}</p>
  </>
  )
}
 