import { MetadataHost } from "./consts";
import axios from "axios";

export interface ASMetaData {
  display: string;
  appendix?: {
    [key: string]: string | string[];
  };
  customNode?: Object;
}

export async function getASMetaData(asn: number) {
  const res = await axios.get(`${MetadataHost}/as/${asn}.json`);
  return res.data;
}
