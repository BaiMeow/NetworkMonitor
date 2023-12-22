import { InjectionKey } from "vue";
import { ASData } from "../api/meta";
import { Ref }from "vue";

export const ASDataKey = Symbol("ASDataKey") as InjectionKey<Ref<ASData>>;