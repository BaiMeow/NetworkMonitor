<script setup lang="ts">
import BGP from './components/BGP.vue';
import OSPF from './components/OSPF.vue';
import { getList } from './api/list';
import { ASData, loadASData } from './api/meta';
import { provide, ref, reactive } from 'vue';
import { ASDataKey } from './inject/key';

const asn = ref(0);
const graph_type = ref('');
const menu_rotate = ref('rotate-closed-margin');
const loading = ref(true);
const dataReady = ref(false);
const loaded = () => {
  loading.value = false;
};
const menu_scale = ref('');
const click_fold = () => {
  if (menu_rotate.value == 'rotate-open') {
    menu_rotate.value = 'rotate-close rotate-closed-margin';
    menu_scale.value = 'menu-fold';
  } else {
    menu_rotate.value = 'rotate-open';
    menu_scale.value = 'menu-expend';
  }
};

const asdata = ref({} as ASData | null);
provide(ASDataKey, asdata);

class bgp {
  enable() {
    graph_type.value = 'bgp';
  }
  display() {
    return 'BGP FULL GRAPH';
  }
}

class ospf {
  asn: number;
  name!: string;
  constructor(asn: number) {
    this.asn = asn;
  }
  async init() {
    this.name = asdata.value?.metadata?.[this.asn + '']?.display || '';
  }
  enable() {
    graph_type.value = 'ospf';
    asn.value = this.asn;
  }
  display() {
    return this.name ? `${this.name} Network` : `AS ${this.asn}`;
  }
}

interface graph {
  enable(): void;
  display(): string;
}

const graph_list = reactive([] as Array<graph>);

const handle_select = (idx: string) => {
  loading.value = true;
  graph_list[+idx].enable();
};

(async () => {
  const data = await loadASData();
  if (data) {
    asdata.value = data;
  }
  const list = await getList();

  list.forEach((graph) => {
    switch (graph.type) {
      case 'bgp':
        graph_list.push(new bgp());
        break;
      case 'ospf':
        const gr = new ospf(graph.asn);
        gr.init();
        graph_list.push(gr);
        break;
    }
  });

  graph_list.sort((a, b) => {
    if (a instanceof ospf && b instanceof bgp) return 1;
    if (b instanceof bgp && a instanceof ospf) return -1;
    return a.display().localeCompare(b.display());
  });

  if (graph_list.length !== 0) {
    graph_list[0]?.enable();
  } else {
    alert('no data');
  }
  dataReady.value = true;
})();
</script>
<template>
  <div class="aside">
    <div class="menu">
      <el-button
        type="primary"
        circle
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
        default-active="0"
        @select="handle_select"
      >
        <el-menu-item
          class="menu-item"
          v-for="(graph, index) in graph_list"
          :index="index.toString()"
        >
          <span>{{ graph.display() }}</span>
        </el-menu-item>
      </el-menu>
    </div>
  </div>
  <template v-if="dataReady">
    <OSPF
      v-if="graph_type === 'ospf' && asn != null && dataReady"
      :asn="asn"
      :loaded="loaded"
    />
    <BGP v-else-if="graph_type === 'bgp' && dataReady" :loaded="loaded" />
  </template>
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
  background-color: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(3px);
  box-shadow: 1.5px 1.5px 1.5px 1.5px rgba(53, 62, 51, 0.3);
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
}

.menu-expend {
  max-height: 70vh;
  max-width: 12em;
}
@media (prefers-color-scheme: dark) {
  .menu {
    background-color: rgba(0, 0, 0, 0.5);
    box-shadow: 1.5px 1.5px 1.5px 1.5px rgba(0, 0, 0, 0.3);
    color: white;
  }
  .menu-item {
    color: white;
  }
  .menu-item:hover {
    background-color: rgba(255, 255, 255, 0.1);
  }
  .el-button--primary {
    border-color: rgba(255, 255, 255, 0.5);
  }
  .menu-button {
    background-color: rgba(0, 0, 0, 0.5);
  }
}
</style>
