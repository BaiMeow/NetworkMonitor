<script setup lang="ts">
import { useUpdatedTime } from '@/state/updated_time'
import { useTimeAgo } from '@vueuse/core'

const updated_time = useUpdatedTime()
const timeAgo = useTimeAgo(() => updated_time.value || 0,{
  updateInterval:1000,
  showSecond:true
})
</script>

<template>
  <div
    v-if="updated_time"
    :class="
      new Date().valueOf() - updated_time.valueOf() > 1000 * 60
        ? 'old-alert'
        : ''
    "
    class="time-stamp"
  >
    <i-ep:refresh class="time-stamp-icon" /><span class="time-stamp-text">
      {{ timeAgo }}</span
    >
  </div>
</template>

<style lang="css" scoped>
.time-stamp {
  color: whitesmoke;
  font-weight: 600;
  background-color: rgb(0, 177, 0);
  padding: 2px 4px;
  border-radius: 5px;
  position: absolute;
  bottom: 0.4rem;
  right: 0.4rem;
  font-size: small;
  display: inline-flex;
  align-items: center;
  gap: 3px;
}

.dark .time-stamp {
  background-color: green;
}

.old-alert {
  background-color: rgb(204, 132, 17);
}

.dark .old-alert {
  background-color: rgb(180, 110, 0);
}

.time-stamp-icon {
  margin-top: 2px;
}
</style>
