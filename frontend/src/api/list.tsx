import axios from "axios";
import { Resp, ApiHost } from "./consts";

type Graph = OSPFGraph | BGPGraph

type OSPFGraph = {
    type:'ospf'
    asn:number
}

type BGPGraph = {
    type:'bgp'
}

export async function getList() {
    const res = await axios.get(`${ApiHost}/api/list`);
    const data = res.data as Resp<Array<Graph>>;
    if (data.code!==0){
        throw new Error(data.msg)
    }
    return data.data;
}