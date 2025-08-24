// src/router.ts
import { createRouter, createWebHashHistory } from 'vue-router'

// 临时全部静态导入，先保证能渲染出来
import Login from './views/Login.vue'
import BGP from './components/BGP.vue'
import OSPF from './components/OSPF.vue'

const routes = [
  { path: '/login', component: Login },

  { path: '/', redirect: '/bgp' },
  { path: '/bgp', component: BGP, meta: { auth: true } },
  { path: '/bgp/:name', component: BGP, meta: { auth: true } },
  { path: '/bgp/:name/:asn', component: BGP, meta: { auth: true } },

  { path: '/ospf', component: OSPF, meta: { auth: true } },
  { path: '/ospf/:name', component: OSPF, meta: { auth: true } },
  { path: '/ospf/:name/:asn', component: OSPF, meta: { auth: true } },
]

const router = createRouter({
  history: createWebHashHistory(),   // 用 hash，生产端不用配置 SPA fallback
  routes,
})

router.beforeEach((to) => {
  console.log('[guard] to =', to.fullPath)
  const token = localStorage.getItem('auth_token')
  if (to.meta?.auth && !token && to.path !== '/login') {
    return { path: '/login', query: { redirect: to.fullPath } }
  }
  console.log('[guard] to=', to.fullPath, 'token=', localStorage.getItem('auth_token'))

})

export default router
