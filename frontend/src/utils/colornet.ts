import { ASData } from "../api/meta";
import { Netmask } from "netmask";

export function prettierNet(
  net: string[],
  asn: string,
  all: ASData["announcements"]
) {
  const record = all.assigned
    ?.filter((a) => a.asn === asn)
    .map((a) => new Netmask(a.prefix));
  const real = net.map((n) => new Netmask(n));
  let str = "";

  // 宣告他人
  real.forEach((n) => {
    if (!record?.every((r) => !r.contains(n))) return;
    if (
      all.assigned?.some((a) => new Netmask(a.prefix).contains(n)) ||
      all.reserved?.some((r) => new Netmask(r).contains(n))
    ) {
      str += `<div class="overannounced-net">${n.toString()}</div>`;
    }
  });

  // 宣告未分配网段
  real.forEach((n) => {
    if (all.assigned?.some((a) => new Netmask(a.prefix).contains(n))) return;
    if (all.reserved?.some((r) => new Netmask(r).contains(n))) return;
    if (all.public?.some((p) => new Netmask(p.prefix).contains(n))) return;
    str += `<div class="unassigned-net">${n.toString()}</div>`;
  })

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

  return str == "" ? "</br>" : str;
}
