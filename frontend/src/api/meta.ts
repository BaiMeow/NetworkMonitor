import { MetadataHost } from './consts'
import axios from 'axios'

export interface ASMetaData {
  display: string
  appendix?: {
    [key: string]: string | string[]
  }
  customNode?: Object
  announce: string[]
}

export interface ASData {
  announcements: {
    assigned: {
      prefix: string
      asn: string
    }[]
    public: {
      prefix: string
      service: {
        prefix: string
        usage: string
        allowedASN: string[]
      }[]
    }[]
    reserved: string[]
  }
  metadata: {
    [key: string]: {
      display: string
      monitor?: {
        appendix?: {
          [key: string]: string | string[]
        }
        customNode?: Object
      }
    }
  }
}

export async function loadASData(): Promise<ASData | null> {
  try {
    const { data } = await axios.get<ASData>(
      `${MetadataHost}/monitor-metadata.json`,
    )
    return data
  } catch (e) {
    console.error(e)
    return null
  }
}
