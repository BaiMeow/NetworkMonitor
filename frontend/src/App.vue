<script setup lang="ts">
import BGP from "./components/BGP.vue"
import OSPF from "./components/OSPF.vue"
import { getList } from "./api/list";
import { getASMetaData } from "./api/meta";
import { ref, reactive } from "vue"

const asn = ref(0);
const graph_type = ref('');
const isCollapse = ref(true);
const menu_rotate = ref('');
const click_fold = ()=>{
  if (menu_rotate.value == 'rotate-open'){
    menu_rotate.value = 'rotate-close';
    isCollapse.value = true;
  }else{
    menu_rotate.value = 'rotate-open';
    isCollapse.value = false;
  }
}


class bgp {
  enable() {
    graph_type.value = "bgp"
  }
  display() {
    return "BGP FULL GRAPH"
  }
}

class ospf {
  asn: number;
  name!: string;
  constructor(asn: number) {
    this.asn = asn;
  }
  async init(){
    this.name = (await getASMetaData(this.asn)).display
  }
  enable() {
    graph_type.value = "ospf"
    asn.value = this.asn
  }
  display() {
    return this.name?`${this.name} Network`:`AS ${this.asn}`
  }
}

interface graph {
  enable(): void
  display(): string
}

const graph_list: Array<graph> = reactive([] as Array<graph>)

const handle_select = (idx: string) => {
  graph_list[+idx].enable()
}

getList().then((list)=>{
  list.forEach(async (graph)=>{
    switch (graph.type){
      case "bgp":
        graph_list.push(new bgp())
        break
      case "ospf":
        const gr = new ospf(graph.asn);
        gr.init()
        graph_list.push(gr)
        break
    }
  })
  if (graph_list.length!==0){
    graph_list[0]?.enable()
  }else{
    alert("no data");
  }
})

</script>
<template>
  <div class="aside">
    <div class="menu">
       <el-button type="primary" circle class="menu-buttom" :class="menu_rotate" @click="click_fold">
        <i-ep-arrow-right-bold/>
      </el-button>
      <el-menu collapse-transition="false" class="menu-list" :collapse="isCollapse" default-active="0" @select="handle_select">
        <el-menu-item class="menu-item"  v-for="(graph, index) in graph_list" :index="index.toString()">
          <span>{{graph.display()}}</span>
        </el-menu-item>
      </el-menu>
    </div>
  </div>
  <OSPF v-if="graph_type === 'ospf'" :asn="asn" />
  <BGP v-else-if="graph_type === 'bgp'" />
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
  max-height: 80vw;
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

@keyframes scale-min{
  from {
    height: fit-content;
    width: fit-content;
  }
  to {
    height: 0px;
    width: 0px;
  }
}

.rotate-open {
  animation: rotation-open 0.5s forwards;
}

.rotate-close {
  animation: rotation-close 0.5s forwards;
}

.menu-item{
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
  transition: height 1s;
  overflow: scroll;
}

.menu-button {
  margin:2px;
}

.menu-list.el-menu--collapse {
  height: 0px;
  width: 0px;
}

.menu-list.el-menu--collapse>.menu-item {
  height: 0px;
  width: 0px;
}


</style>
