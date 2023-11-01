import axios from "axios";
import { MetadataHost,Resp,ApiHost } from "./consts";

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
  return res.data as Resp<BGP>;
}

export interface ASMetaData {
    display: string
    appendix?: {
        [key: string]: string | string[]  
    }
    customNode?: Object
}

export async function getASMetaData(asn: number) {
    const res = await axios.get(`${MetadataHost}/as/${asn}.json`);
    return res.data as ASMetaData
}
