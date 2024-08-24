import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";

createApp(App).mount("#app");

// document.documentElement 是全局变量时
const el = document.documentElement;
el.style.setProperty("--el-color-primary", "#607d8b");
el.style.setProperty("--el-color-primary-light-3", "#78909c");
el.style.setProperty("--el-color-primary-light-5", "#90a4ae");
el.style.setProperty("--el-color-primary-light-7", "#b0bec5");
el.style.setProperty("--el-color-primary-light-8", "#cfd8dc");
el.style.setProperty("--el-color-primary-light-9", "#eceff1");
el.style.setProperty("--el-color-primary-dark-2", "#455a64");
