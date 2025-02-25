<script setup lang="ts">
import { useUpdatedTime } from '@/state/updated_time'

const updated_time = useUpdatedTime()
</script>

<template>
  <div
    v-if="updated_time"
    :class="new Date() - updated_time > 1000 * 60 ? 'old-alert' : ''"
    class="time-stamp"
  >
    <i-ep:refresh class="time-stamp-icon" /><span class="time-stamp-text">{{
      updated_time.toLocaleTimeString()
    }}</span>
  </div>
</template>

<style lang="css" scoped>
.time-stamp {
  color: whitesmoke;
  font-weight: 600;
  background-color: rgb(255, 157, 87);
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
  background-color: chocolate;
}

.old-alert {
  background-color: rgb(227, 38, 38);
}

.dark .old-alert {
  background-color: rgb(188, 18, 18);
}
</style>
