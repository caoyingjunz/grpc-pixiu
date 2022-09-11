import {
  createRouter,
  createWebHashHistory
} from "vue-router";

const routes = [{ // 登陆界面
    path: "/",
    name: 'login',
    component: () => import( /*webpackChunkName:'Login'*/ '@/page/login/login.vue')
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// 路由导航守卫
router.beforeEach((to, from, next) => {
  if (to.path == '/') {
      next()
  } else {
      const token = window.sessionStorage.getItem('token')
      if (!token) {
          next('/')
      } else {
          next()
      }
  }
})

export default router;