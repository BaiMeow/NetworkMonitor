<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { useDark, useToggle } from '@vueuse/core'

const isDark = useDark()
const toggleDark = useToggle(isDark)

const toggleTheme = (event: MouseEvent) => {
  const x = event.clientX
  const y = event.clientY
  const endRadius = Math.hypot(
    Math.max(x, innerWidth - x),
    Math.max(y, innerHeight - y),
  )

  // 兼容性处理
  if (!document.startViewTransition) {
    toggleDark()
    return
  }
  const transition = document.startViewTransition(async () => {
    toggleDark()
  })

  transition.ready.then(() => {
    const clipPath = [
      `circle(0px at ${x}px ${y}px)`,
      `circle(${endRadius}px at ${x}px ${y}px)`,
    ]
    document.documentElement.animate(
      {
        clipPath: isDark.value ? [...clipPath].reverse() : clipPath,
      },
      {
        duration: 400,
        easing: 'ease-in',
        pseudoElement: isDark.value
          ? '::view-transition-old(root)'
          : '::view-transition-new(root)',
      },
    )
  })
}
</script>

<template>
  <el-icon class="icon" @click="toggleTheme"
    ><i-ep-moon v-if="isDark" /> <i-ep-sunny v-else
  /></el-icon>
</template>

<style scoped></style>
