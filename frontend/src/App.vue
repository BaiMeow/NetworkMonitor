<script setup lang="ts">
import { getList } from './api/list'
import { ref, watch, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useASMeta } from './state/meta'

const router = useRouter()
const menu_rotate = ref('rotate-closed-margin')
const menu_clp = ref('menu-clp')
const dataReady = ref(false)
const menu_scale = ref('')

const click_fold = () => {
  if (menu_rotate.value == 'rotate-open') {
    menu_rotate.value = 'rotate-close rotate-closed-margin'
    menu_scale.value = 'menu-fold'
    menu_clp.value = 'menu-clp'
  } else {
    menu_rotate.value = 'rotate-open'
    menu_scale.value = 'menu-expend'
    menu_clp.value = ''
  }
}

const ASMeta = useASMeta()

class bgp {
  display() {
    return 'BGP FULL GRAPH'
  }
  path() {
    return '/bgp'
  }
}

class ospf {
  asn: number
  constructor(asn: number) {
    this.asn = asn
  }
  display() {
    const display = ASMeta.value?.metadata?.[this.asn+'']?.display;
    return display? `${display} Network`: `AS ${this.asn}`
  }
  path() {
    return `/ospf/${this.asn}`
  }
}

interface graph {
  display(): string
  path(): string
}

const graph_list = ref([] as Array<graph>)

async function load_list() {
  const list = await getList()

  if (list.length === 0) {
    alert('no data')
    return
  }

  const local_list = [] as Array<graph>
  list.forEach((graph) => {
    switch (graph.type) {
      case 'bgp':
        local_list.push(new bgp())
        break
      case 'ospf': {
        local_list.push(new ospf(graph.asn))
        break
      }
    }
  })

  local_list.sort((a, b) => {
    if (a instanceof ospf && b instanceof bgp) return 1
    if (a instanceof bgp && b instanceof ospf) return -1
    return a.display().localeCompare(b.display())
  })

  graph_list.value = local_list
  if (router.currentRoute.value.path === '/') {
    router.push(local_list[0].path())
  }
  dataReady.value = true
}

watch(
  () => ASMeta.value,
  () =>
    (graph_list.value = [...graph_list.value].sort((a, b) => {
      if (a instanceof ospf && b instanceof bgp) return 1
      if (a instanceof bgp && b instanceof ospf) return -1
      return a.display().localeCompare(b.display())
    })),
)

load_list()
const refresh = setInterval(() => load_list(), 60 * 1000)
onUnmounted(() => {
  clearInterval(refresh)
})
</script>
<template>
  <div class="aside">
    <div class="menu" :class="menu_clp">
      <el-button
        type="primary"
        class="menu-button transition-06s"
        :class="menu_rotate"
        @click="click_fold"
      >
        <i-ep-arrow-right-bold />
      </el-button>
      <el-menu
        :collapse-transition="false"
        class="menu-list transition-06s"
        :class="menu_scale"
        router
      >
        <el-menu-item
          class="menu-item"
          v-for="graph in graph_list"
          :key="graph.path()"
          :index="graph.path()"
        >
          <span>{{ graph.display() }}</span>
        </el-menu-item>
      </el-menu>
    </div>
  </div>
  <Graph />
  <router-view v-if="dataReady" />
</template>

<style scoped>
.aside {
  position: absolute;
  display: flex;
  align-items: center;
  height: 100%;
  width: fit-content;
}

.menu {
  display: flex;
  align-items: center;
  z-index: 10;
  max-height: 80vh;
  flex-direction: column;
  border-style: solid;
  border-width: 1px;
  transition: border-width 0.6s;
}

.menu-clp {
  border-width: 0px;
}

@keyframes rotation-open {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(180deg);
  }
}

@keyframes rotation-close {
  from {
    transform: rotate(180deg);
  }

  to {
    transform: rotate(360deg);
  }
}

.rotate-open {
  animation: rotation-open 0.5s forwards;
}

.rotate-close {
  animation: rotation-close 0.5s forwards;
}

.rotate-closed-margin {
  margin: 0px !important;
}

.menu-item {
  border-radius: 20px;
}

.menu {
  backdrop-filter: blur(3px);
  border-radius: 20px;
}

.menu-list {
  border-right: 0px;
  background-color: rgba(255, 255, 255, 0);
  overflow-y: scroll;
  overflow-x: hidden;
  max-width: 0px;
  max-height: 0vw;
}

.transition-06s {
  transition: all, 0.6s;
  transition-timing-function: ease-in-out;
}

.menu-list::-webkit-scrollbar {
  display: none;
}

.menu-button {
  margin-top: 1vh;
  margin-bottom: 1vh;
  width: 2rem;
  height: 2rem;
  padding: 0;
  border-width: 0;
}

.menu-expend {
  max-height: 70vh;
  max-width: 12em;
}

.loading {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 2rem;
  font-weight: bold;
  color: #2242a3;
}

.graph {
  height: 100dvh;
  width: 100vw;
  top: 0;
  left: 0;
  position: absolute;
}
</style>
