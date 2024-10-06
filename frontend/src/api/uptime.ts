import axios from 'axios'
import { Resp, ApiHost } from './consts'

type UptimeRecent = Array<boolean>

export async function getUptimeRecent(asn: number) {
  const res = await axios.get(`${ApiHost}/api/bgp/uptime/${asn}/recent`)
  const data = res.data as Resp<UptimeRecent>
  if (data.code !== 0) {
    throw new Error(data.msg)
  }
  return data.data
}