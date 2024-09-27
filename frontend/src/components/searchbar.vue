<script lang="ts" setup>
import { Netmask } from "netmask";
import { ref, Ref, watch } from "vue";
export interface selectItem {
  asn?: string;
  name?: string;
  network?: Array<string>;
  value: string;
  label: string;
  onselected: () => void;
}

const props = defineProps<{
  data: Array<selectItem>;
}>();

const options: Ref<Array<selectItem>> = ref([]);

const input: Ref<string> = ref("");

watch(
  () => props.data,
  () => search(input.value)
);

const onchange = (value: string) => {
  props.data
    .find((v) => {
      return v.value === value;
    })
    ?.onselected();
};

function search(val: string) {
  if (!val) {
    options.value = props.data;
    return;
  }
  options.value = props.data.filter((v) => {
    return (
      v.label.toLowerCase().includes(val.toLowerCase()) ||
      v.value.toLowerCase().includes(val.toLowerCase()) ||
      v.network?.some((n) => n.includes(val)) ||
      (() => {
        try {
          const nums = val.split(".");
          if (nums.length !== 4) {
            for (let i = nums.length; i < 4; i++) {
              nums.push("0");
            }
          }
          const sr = new Netmask(nums.join("."));
          return v.network?.some((n) => new Netmask(n).contains(sr.base));
        } catch {
          return false;
        }
      })()
    );
  });
}
</script>

<template>
  <div style="display: flex; align-items: center">
    <el-icon color="--el-color-primary" style="margin-right: 0.5rem">
      <i-ep-search />
    </el-icon>
    <el-select
      v-model="input"
      @blur="input = ''"
      filterable
      placeholder="Search"
      @change="onchange"
      :filter-method="search"
    >
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      >
        <span style="float: left; width: 40%; margin-right: 1rem">{{
          item.label
        }}</span>
        <span
          v-if="item.label !== item.value"
          style="
            float: right;
            color: var(--el-text-color-secondary);
            font-size: 13px;
          "
        >
          {{ item.value }}
        </span>
      </el-option>
    </el-select>
  </div>
</template>

<style>
@media (prefers-color-scheme: dark) {
  .el-select-dropdown__item {
    background-color: #333;
    color: #fff;
  }

  .el-select-dropdown__item:hover {
    background-color: #444;
  }
  .el-select-dropdown__item.is-hovering {
    background-color: #444;
  }
  .el-select__wrapper {
    background-color: black;
    color: #e0e0e0;
    box-shadow: 0 0 0 2px #333;
  }

  .el-select__wrapper:hover {
    box-shadow: 0 0 0 2px gray;
  }

  .el-select__wrapper:focus {
    box-shadow: 0 0 0 2px #cfe909;
  }
  .el-select__popper.el-popper,
  .el-select__popper.el-popper .el-popper__arrow:before {
    border: 1px solid #333;
  }
  .el-select__popper.el-popper .el-scrollbar__wrap {
    background-color: #333;
  }
  .el-popper.is-light,
  .el-popper.is-light > .el-popper__arrow:before {
    background: #333;
  }
}
</style>
