import axios from "axios";
import { Resp, ApiHost } from "./consts";

interface Router {
  router_id: string;
  metadata?: Object;
}

interface Link {
  src: string;
  dst: string;
  cost: number;
}

interface Area {
  area_id: string;
  router: Router[];
  links: Link[];
}

export async function getOSPF(asn: number) {
  const res = await axios.get(`${ApiHost}/api/ospf/${asn}`);
  const data = res.data as Resp<Array<Area>>;
  if (data.code!==0){
    throw new Error(data.msg)
  }  
  return data.data
}
