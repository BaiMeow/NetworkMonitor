<template>
  <div class="login-wrap">
    <div class="bg-deco"></div>
    <div class="card">
      <div class="brand">
        <div class="logo">TY</div>
        <div class="titles">
          <h1>Network Monitor</h1>
          <p class="sub">DN06甜源家庭网络基础设施监控平台</p>
        </div>
      </div>

      <form @submit.prevent="doLogin" class="form">
        <label class="field">
          <span>用户名</span>
          <input
            v-model.trim="username"
            type="text"
            autocomplete="username"
            placeholder="admin"
            required
            @keydown.enter="doLogin"
          />
        </label>

        <label class="field">
          <span>密码</span>
          <div class="pwd">
            <input
              v-model="password"
              :type="showPwd ? 'text' : 'password'"
              autocomplete="current-password"
              placeholder="••••••••"
              required
              @keydown.enter="doLogin"
            />
            <button class="toggle" type="button" @click="showPwd = !showPwd">
              {{ showPwd ? '隐藏' : '显示' }}
            </button>
          </div>
        </label>

        <div class="row">
          <label class="remember">
            <input type="checkbox" v-model="remember" />
            <span>记住我</span>
          </label>
          <button class="link" type="button" @click="fillDemo">填充示例</button>
        </div>

        <button class="submit" :disabled="loading">
          <span v-if="!loading">登录</span>
          <span v-else class="spinner"></span>
        </button>

        <p v-if="err" class="err">{{ err }}</p>
      </form>

      <div class="footer">
        <span>© {{ year }} 山东甜源创新</span>
        <span class="sep">·</span>
        <span class="muted">可信网络，安全世界</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()
const route = useRoute()

const username = ref('')
const password = ref('')
const remember = ref(true)
const showPwd  = ref(false)
const loading  = ref(false)
const err      = ref('')

const year = computed(() => new Date().getFullYear())

function fillDemo() {
  if (!username.value) username.value = 'ty'
  if (!password.value) password.value = 'admin'
}

async function doLogin() {
  if (loading.value) return
  err.value = ''
  loading.value = true
  try {
    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username: username.value, password: password.value }),
    })
    if (!res.ok) {
      const t = await res.text().catch(()=>'')
      throw new Error(t || '登录失败')
    }
    const data = await res.json() as { token: string }
    // 记住或会话：简单处理——不勾选就只存在当前会话
    if (remember.value) {
      localStorage.setItem('auth_token', data.token)
    } else {
      sessionStorage.setItem('auth_token', data.token)
      localStorage.removeItem('auth_token')
    }
    const redirect = (route.query.redirect as string) || '/bgp'
    router.replace(redirect)
  } catch (e:any) {
    err.value = e?.message || '登录失败'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
/* 背景：柔和渐变 + 光晕 */
.login-wrap {
  min-height: 100vh;
  display: grid;
  place-items: center;
  position: relative;
  background:
    radial-gradient(900px 500px at -10% -10%, #e0e7ff55, transparent 60%),
    radial-gradient(900px 500px at 110% 110%, #fde68a55, transparent 60%),
    linear-gradient(180deg, #f8fafc, #eef2ff);
}
:root.dark .login-wrap {
  background:
    radial-gradient(900px 500px at -10% -10%, #1f293755, transparent 60%),
    radial-gradient(900px 500px at 110% 110%, #0f172a55, transparent 60%),
    linear-gradient(180deg, #0b1220, #0a0f1d);
}
.bg-deco {
  position: absolute; inset: 0;
  background:
    radial-gradient(350px 220px at 20% 20%, #93c5fd55, transparent 60%),
    radial-gradient(350px 220px at 80% 80%, #fca5a555, transparent 60%);
  filter: blur(30px);
  opacity: .6;
}

/* 卡片：玻璃拟态 */
.card {
  position: relative;
  width: min(92vw, 380px);
  padding: 26px 24px 18px;
  border-radius: 18px;
  background: rgba(255,255,255,.75);
  backdrop-filter: blur(10px);
  box-shadow: 0 12px 40px rgba(0,0,0,.10), inset 0 1px rgba(255,255,255,.6);
  border: 1px solid rgba(255,255,255,.6);
}
:root.dark .card {
  background: rgba(20,24,34,.55);
  border-color: rgba(255,255,255,.08);
}

/* 标题区 */
.brand { display: flex; align-items: center; gap: 12px; margin-bottom: 10px; }
.logo {
  width: 42px; height: 42px; border-radius: 12px;
  display: grid; place-items: center; font-weight: 900;
  background: linear-gradient(135deg, #2563eb, #7c3aed);
  color: white; letter-spacing: .5px;
  box-shadow: 0 6px 18px rgba(37,99,235,.35);
}
.titles h1 { font-size: 18px; margin: 0; }
.titles .sub { margin: 2px 0 0; font-size: 12px; opacity: .65; }

/* 表单 */
.form { margin-top: 10px; display: grid; gap: 12px; }
.field span { display: block; font-size: 12px; opacity: .7; margin: 2px 2px 6px; }
input[type="text"], input[type="password"] {
  width: 100%; height: 40px; padding: 0 12px;
  border: 1px solid #e5e7eb; border-radius: 10px; outline: none;
  background: #fff; color: #111827;
}
:root.dark input[type="text"], :root.dark input[type="password"] {
  border-color: #334155; background: #0b1220; color: #e5e7eb;
}

.pwd { display: grid; grid-template-columns: 1fr auto; align-items: center; }
.toggle {
  margin-left: 8px; height: 34px; padding: 0 10px; border-radius: 8px;
  border: 1px solid #e5e7eb; background: transparent; cursor: pointer;
}
:root.dark .toggle { border-color: #334155; color: #cbd5e1; }

.row { display:flex; align-items:center; justify-content:space-between; margin-top: 2px; }
.remember { display:flex; align-items:center; gap:8px; font-size: 12px; opacity:.85; }
.link { font-size: 12px; opacity:.8; background:transparent; border:none; cursor:pointer; color:#2563eb; }

.submit {
  height: 42px; border: none; border-radius: 12px; cursor: pointer;
  background: linear-gradient(135deg, #2563eb, #7c3aed); color: white;
  font-weight: 800; letter-spacing: .3px; box-shadow: 0 10px 24px rgba(37,99,235,.35);
}
.submit:disabled { opacity: .6; cursor: not-allowed; }

.err { color: #ef4444; font-size: 12px; margin-top: 6px; text-align: center; }

.footer {
  margin-top: 16px; display:flex; justify-content:center; gap:8px;
  font-size: 12px; opacity:.6;
}
.sep { opacity:.35; }
.muted { opacity:.6; }

/* 加载中小圆环 */
.spinner {
  display:inline-block; width:18px; height:18px; border-radius:50%;
  border:2px solid currentColor; border-right-color: transparent;
  animation: spin 0.8s linear infinite;
  vertical-align: -3px;
}
@keyframes spin { to { transform: rotate(360deg); } }
</style>
