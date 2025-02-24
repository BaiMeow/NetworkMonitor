import { createWebHashHistory, createRouter } from 'vue-router'

import BGP from './components/BGP.vue'
import OSPF from './components/OSPF.vue'

const routes = [
  {
    path: '/bgp/:name/:asn',
    component: BGP,
  },
  {
    path: '/bgp/:name',
    component: BGP,
  },
  {
    path: '/ospf/:asn',
    component: OSPF,
  },
]

export const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})
