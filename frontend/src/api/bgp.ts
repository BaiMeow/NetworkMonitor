import axios from 'axios'
import { Resp, ApiHost } from './consts'

export interface BGP {
  as: AS[]
  link: Link[]
  updated_at: Date
}

export interface AS {
  asn: number
  network: string[]
  metadata?: object
}

export interface Link {
  src: number
  dst: number
}

export async function getBGP(name: string) {
  const { data } = await axios.get<Resp<BGP>>(`${ApiHost}/api/bgp/${name}`)
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  data.data.updated_at = new Date(data.data.updated_at)
  return data.data
}

export async function getCloseness(name: string) {
  const res = await axios.get(`${ApiHost}/api/bgp/${name}/analysis/closeness`)
  const data = res.data as Resp<{
    [key: string]: number
  }>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data
}

export async function getBetweenness(name: string) {
  const res = await axios.get(`${ApiHost}/api/bgp/${name}/analysis/betweenness`)
  const data = res.data as Resp<{
    [key: string]: number
  }>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data
}
