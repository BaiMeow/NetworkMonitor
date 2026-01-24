import axios from 'axios'
import { Resp, ApiHost } from './consts'

export interface Router {
  router_id: string
  subnet: string[]
  metadata?: object
}

export interface Link {
  src: string
  dst: string
  cost: number
}

export interface Area {
  area_id: string
  router: Router[]
  links: Link[]
}

export interface OSPF {
  graph: Area[]
  updated_at: Date | null
}

export async function getOSPF(asn: number) {
  const { data } = await axios.get<Resp<OSPF>>(`${ApiHost}/api/ospf/${asn}`)
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  if (data.data.updated_at)
    data.data.updated_at = new Date(data.data.updated_at)
  return data.data
}

export async function getCloseness(asn: number) {
  const res = await axios.get(`${ApiHost}/api/ospf/${asn}/analysis/closeness`)
  const data = res.data as Resp<{
    [key: string]: number
  }>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data
}

export async function getBetweenness(asn: number) {
  const res = await axios.get(`${ApiHost}/api/ospf/${asn}/analysis/betweenness`)
  const data = res.data as Resp<{
    [key: string]: number
  }>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data
}

export async function getPathBetweenness(asn: number) {
  const res = await axios.get(
    `${ApiHost}/api/ospf/${asn}/analysis/pathBetweenness`,
  )
  const data = res.data as Resp<
    Array<{
      src: string
      dst: string
      cost: number
      betweenness: number
    }>
  >
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data
}
