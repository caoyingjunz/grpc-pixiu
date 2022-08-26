<template>
  <div class="header-container">
    <!-- <div>logo</div> -->
    <div @click="menu.changeCollapse" v-if="!menu.isCollapsed">
      <el-icon :size="20"><Fold /></el-icon>
    </div>
    <div @click="menu.changeCollapse" v-else>
      <el-icon :size="20"><Expand /></el-icon>
    </div>
    <div>abc</div>
    <div>头像</div>
  </div>
</template>

<script setup>
import { Fold, Expand } from "@element-plus/icons-vue";
import { useMenuStore } from "@/store";
const menu = useMenuStore();
</script>

<style scoped>
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
}
</style>
