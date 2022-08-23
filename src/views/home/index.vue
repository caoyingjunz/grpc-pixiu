<script setup>
import { ElButton } from "element-plus";
import { useCounterStore } from "@/store";
const counter = useCounterStore();
const url = import.meta.env.VITE_BASE_URL;
</script>

<template>
  <div>{{ url }}</div>
  <div>{{ counter.count }}</div>
  <div>
    <el-button type="primary" @click="counter.increment">测试</el-button>
  </div>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
