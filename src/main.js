import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import { createPinia } from "pinia";
import routers from "./routers/index";

createApp(App).use(createPinia()).use(routers).mount("#app");
