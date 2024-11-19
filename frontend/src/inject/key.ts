import { InjectionKey } from 'vue'
import { ASData } from '../api/meta'
import { Ref } from 'vue'

export const ASDataKey = Symbol('ASDataKey') as InjectionKey<Ref<ASData | null>>
export const LoadingKey = Symbol('LoadingKey') as InjectionKey<{
  setLoading: () => void
  doneLoading: () => void
}>
