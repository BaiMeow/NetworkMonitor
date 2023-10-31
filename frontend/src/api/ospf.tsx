import axios from "axios";
import { Resp } from "./consts";

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
  const res = await axios.get(`/api/ospf/${asn}`);
  return res.data as Resp<Array<Area>>;
}
