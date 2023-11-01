import axios from "axios";
import { Resp,ApiHost } from "./consts";

interface BGP {
    as: AS[]
    link: Link[]
}

interface AS {
    asn: number
    network: string[]
    metadata?: object
}

interface Link {
    src: number;
    dst: number;
}

export async function getBGP() {
  const res = await axios.get(`${ApiHost}/api/bgp`);
  const data = res.data as Resp<BGP>;
  if (data.code!==0){
    throw new Error(data.msg)
  }
  return data.data;
}

