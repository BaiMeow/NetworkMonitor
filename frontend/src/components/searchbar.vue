<script lang="ts" setup>
import { Netmask } from 'netmask';
import { ref, Ref, watch } from 'vue';
export interface selectItem {
  asn?: string;
  name?: string;
  network?: Array<string>;
  value: string;
  label: string;
  onselected: () => void;
}

const props = defineProps<{
  data: Array<selectItem>
}>();

const options: Ref<Array<selectItem>> = ref([]);

const input: Ref<string> = ref('');

watch(
  () => props.data,
  () => search(input.value)
)

const onchange = (value: string) => {
  props.data.find((v) => {
    return v.value === value
  })?.onselected()
}

function search(val: string) {
  if (!val) {
    options.value = props.data
    console.log(options.value)
    return
  }
  options.value = props.data.filter((v) => {
    return v.label.toLowerCase().includes(val.toLowerCase()) || v.value.toLowerCase().includes(val.toLowerCase()) || v.network?.some((n) => n.includes(val)) || (() => {
      try {
        const sr = new Netmask(val)
        return v.network?.some((n) => new Netmask(n).contains(sr.base))
      } catch {
        return false
      }
    })()
  })
}

</script>

<template>
  <div style="display:flex;align-items: center">
    <el-icon color="#409EFF" style="margin-right:.5rem">
      <i-ep-search />
    </el-icon>
    <el-select v-model="input" filterable placeholder="Search" @change="onchange" :filter-method="search">
      <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
        <span style="float: left;width:40%;margin-right:1rem">{{ item.label }}</span>
        <span v-if="item.label !== item.value" style="
          float: right;
          color: var(--el-text-color-secondary);
          font-size: 13px;
        ">
          {{ item.value }}
        </span>
      </el-option>
    </el-select>
  </div>

</template>

<style scoped></style>