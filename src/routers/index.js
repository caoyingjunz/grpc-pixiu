import { createRouter, createWebHistory } from "vue-router";
import { Home, Login } from "@/views";
import Base from "@/layouts/base/index.vue";

const routes = [
  {
    path: "/",
    redirect: "/home",
  },
  {
    path: "/",
    component: Base,
    children: [
      {
        path: "home",
        component: Home,
      },
    ],
  },
  { path: "/login", component: Login },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
