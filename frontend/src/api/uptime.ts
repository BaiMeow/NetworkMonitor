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

interface OSPFUptimeLinksResp {
  in: UptimeLinks[]
  out: UptimeLinks[]
}

export async function getBGPUptimeRecent(grName: string, asn: number) {
  const res = await axios.get(
    `${ApiHost}/api/bgp/${grName}/uptime/${asn}/recent`,
  )
  const data = res.data as Resp<UptimeRecent>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data
}

export async function getBGPUptimeLinks(
  grName: string,
  asn: number,
  time: string,
  window: string,
) {
  const res = await axios.get(
    `${ApiHost}/api/bgp/${grName}/uptime/${asn}/links`,
    {
      params: {
        time,
        window,
      },
    },
  )
  const data = res.data as Resp<Array<UptimeLinksResp>>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data.map((item) => {
    return {
      time: new Date(item.time),
      links: item.links,
    }
  })
}

export async function getOSPFUptimeRecent(asn: number, routerId: string) {
  const res = await axios.get(
    `${ApiHost}/api/ospf/${asn}/uptime/${routerId}/recent`,
  )
  const data = res.data as Resp<UptimeRecent>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data
}

export async function getOSPFUptimeLinks(
  asn: number,
  routerId: string,
  time: string,
  window: string,
) {
    const res = await axios.get(
      `${ApiHost}/api/ospf/${asn}/uptime/${routerId}/links`,
      {
        params: {
          time,
          window,
        },
      },
    )
    const data = res.data as Resp<OSPFUptimeLinksResp>
    if (data.code !== 0) {
      throw new Error(data.msg)
    }
    return {
      in: data.data.in.map((item) => {
        return {
          time: new Date(item.time),
          links: item.links,
        }
      }),
      out: data.data.out.map((item) => {
        return {
          time: new Date(item.time),
          links: item.links,
        }
      }),
    }
}
