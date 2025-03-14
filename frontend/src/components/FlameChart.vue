<template>
  <v-container>
    <div
      ref="flameChart"
      class="flame-container"
    />
  </v-container>
</template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import * as d3 from 'd3';
  import { flamegraph } from 'd3-flame-graph';
  import { profileToFlameGraph } from '@/utils/pprof-converter';
  import { Profile } from '@/utils/pprof';
  import type { FlameNode } from '@/utils/pprof-converter';
  const props = defineProps({
    profile: {
      type: Profile,
      required: true,
    },
  });

  const flameChart = ref<HTMLElement | null>(null);
  
  // 如果props更新
  watch(
  () => props.profile,
  (newVal) => {
    console.log('Profile 数据更新:', newVal)
    // 在这里触发火焰图更新逻辑
    initFlameGraph(profileToFlameGraph(newVal));
  },
  { deep: true } // 深度监听对象变化
  )
  
  // 初始化火焰图
  function initFlameGraph(data: FlameNode) {
    if (!flameChart.value) return;
  
    const chart = flamegraph()
      .tooltip(true)
      .inverted(true) // Go 调用习惯更符合自顶向下
      .sort((a, b) => d3.ascending(a.name, b.name));
  
    d3.select(flameChart.value)
      .datum(data)
      .call(chart);
  }
</script>
  
  <style>

  </style>
  