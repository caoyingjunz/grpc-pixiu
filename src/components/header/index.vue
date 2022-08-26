<template>
  <div class="header-container">
    <!-- <div>logo</div> -->
    <el-icon :size="20"><Fold /></el-icon>
    <div>abc</div>
    <div>头像</div>
  </div>
</template>

<script setup>
import { ref } from "vue";
</script>

<style scoped>
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
}
</style>
