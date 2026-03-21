import { useState } from 'react'

export const PostButton = () => {
  const [data, setData] = useState('')
  const [status, setStatus] = useState('')

  const postStatus = async () => {
    const res = await fetch('http://localhost:8080/post', {
      method: 'POST',
      headers: {
        'Content-Type': 'text/plain',
      },
      body: data,
    })

    const text = await res.text()
    setStatus(text)
  }

  return (
    <>
      <input
        value={data}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
          setData(e.target.value)
        }
      />

      <button onClick={postStatus}>Post status</button>

      <p>{status}</p>
    </>
  )
}