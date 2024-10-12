<script setup lang="ts">
const upColor = '#00c853'
const downColor = '#455a64'

const props = defineProps<{
  data: Array<number>
}>()
</script>

<template>
  <div class="squares">
    <div
      v-for="d in props.data"
      class="square"
      :style="{
        backgroundColor: d ? upColor : downColor,
      }"
    ></div>
  </div>
</template>

<style scoped>
.square {
  display: inline-block;
  width: 8px;
  height: 20px;
  margin: 3px;
  border-radius: 2px;
  transition: background-color 0.5s;
  transition-timing-function: ease;
}

.squares {
  display: flex;
  justify-content: center;
}
</style>
