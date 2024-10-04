import axios from 'axios'
import { Resp, ApiHost } from './consts'

type Graph = OSPFGraph | BGPGraph

type OSPFGraph = {
  type: 'ospf'
  asn: number
}

type BGPGraph = {
  type: 'bgp'
}

export async function getList() {
  const {
    data: { code, msg, data },
  } = await axios.get<Resp<Array<Graph>>>(`${ApiHost}/api/list`)
  if (code !== 0) {
    throw new Error(msg)
  }
  return data
}
