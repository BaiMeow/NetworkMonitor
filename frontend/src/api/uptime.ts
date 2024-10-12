import axios from 'axios'
import { Resp, ApiHost } from './consts'

type UptimeRecent = Array<boolean>

export interface UptimeLinks {
  time: Date
  links: number
}

interface UptimeLinksResp {
  time: string
  links: number
}

export async function getUptimeRecent(asn: number) {
  const res = await axios.get(`${ApiHost}/api/bgp/uptime/${asn}/recent`)
  const data = res.data as Resp<UptimeRecent>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data
}

export async function getUptimeLinks(asn: number,time: string,window: string) {
  const res = await axios.get(`${ApiHost}/api/bgp/uptime/${asn}/links`,{
    params: {
      time,
      window
    }
  })
  const data = res.data as Resp<Array<UptimeLinksResp>>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data.map((item) => {
    return {
      time: new Date(item.time),
      links: item.links
    }
  })
}