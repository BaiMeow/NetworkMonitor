import { getASMetaDataFromCache } from "../api/meta";
import { Netmask } from "netmask";

export function prettierNet(net: string[], asn: string) {
  const meta = getASMetaDataFromCache(parseInt(asn));
  const record = meta?.announce.map((n) => new Netmask(n));
  const real = net.map((n) => new Netmask(n));
  let str = "";

  // 错误宣告
  real
    .filter((n) => record?.every((r) => !r.contains(n)))
    .forEach(
      (n) => (str += `<div class="overannounced-net">${n.toString()}</div>`)
    );

  // 未宣告
  record
    ?.filter((r) => real.every((n) => !n.contains(r)))
    ?.forEach(
      (n) => (str += `<div class="unannounced-net">${n.toString()}</div>`)
    );

  // 正确宣告
  real
    .filter((n) => record?.some((r) => r.contains(n)))
    .forEach((n) => (str += `<div>${n.toString()}</div>`));

  return str;
}
