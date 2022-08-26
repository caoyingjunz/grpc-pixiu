import { defineStore } from "pinia";
const useMenuStore = defineStore("menu", {
  state: () => {
    return {
      isCollapsed: false,
    };
  },
  // 也可以定义为
  // state: () => ({ count: 0 })
  actions: {
    changeCollapse() {
      this.isCollapsed = !this.isCollapsed;
    },
  },
});

export default useMenuStore;
