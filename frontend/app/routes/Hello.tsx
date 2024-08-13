import { json, useLoaderData } from '@remix-run/react'
import React from 'react'
import { Hellos } from '~/types/types'

export const loader = async () => {
  const hellosResponse = await fetch('http://localhost:8080/hello')
  const hellos: Hellos[] = await hellosResponse.json()

  return json({ hellos })
}

const Hello: React.FC = () => {
  const { hellos } = useLoaderData<typeof loader>()
  return (
    <div>
      {hellos.map((hello) => (
        <p key={hello.id}>{hello.content}</p>
      ))}
    </div>
  )
}

export default Hello