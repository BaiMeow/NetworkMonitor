import { MetadataHost } from "./consts";
import axios from "axios";

const metaMap = new Map<number, ASMetaData>();

export interface ASMetaData {
  display: string;
  appendix?: {
    [key: string]: string | string[];
  };
  customNode?: Object;
  announce: string[];
}

export async function getASMetaData(asn: number) {
  const res = await axios.get(`${MetadataHost}/as/${asn}.json`);
  metaMap.set(asn, res.data as ASMetaData);
  return res.data as ASMetaData;
}

export function getASMetaDataFromCache(asn: number) {
  return metaMap.get(asn);
}
