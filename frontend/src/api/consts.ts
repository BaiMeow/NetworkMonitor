const ApiHost = ''
const MetadataHost = ''

export { ApiHost, MetadataHost }

export interface Resp<T> {
    code: number
    msg: string
    data: T
}