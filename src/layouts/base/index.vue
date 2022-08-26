<template>
  <div class="container">
    <el-container class="container">
      <el-aside class="el-aside-style">
        <div
          style="
            height: 60px;
            display: flex;
            align-items: center;
            justify-content: center;
          "
        >
          logo
        </div>
        <Menu />
      </el-aside>
      <el-container>
        <el-header class="el-header-background"><Header /></el-header>
        <el-main><router-view></router-view></el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import { Header, Menu } from "@/components";
import { ref } from "vue";
</script>

<style scoped>
.el-header-background {
  background-color: #242424;
  color: rgba(255, 255, 255, 0.87);
}
.el-aside-style {
  /* background-color: rgba(0, 0, 0, 0.5); */
  width: fit-content;
  height: 100%;
  /* padding: 0 20px; */
}
.container {
  height: 100%;
  width: 100%;
}
</style>
