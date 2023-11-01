const ApiHost = 'https://monitor.dn11.baimeow.cn'
const MetadataHost = 'https://metadata.dn11.baimeow.cn'

export { ApiHost, MetadataHost }

export interface Resp<T> {
    code: number
    msg: string
    data: T
}