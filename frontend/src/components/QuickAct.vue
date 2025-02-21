<template>
  <!-- 快速操作 -->
  <v-row class="mt-4">
    <v-col
      v-for="action in quickActions"
      :key="action.title"
      cols="6"
      sm="3"
    >
      <v-btn
        :prepend-icon="action.icon"
        variant="tonal"
        block
        height="100"
        @click="action.handler"
      >
        <span class="text-body-1">{{ action.title }}</span>
      </v-btn>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
    const quickActions = ref([
    { title: '新建测试', icon: 'mdi-plus', handler: () => console.log('New Test') },
    { title: '历史记录', icon: 'mdi-history', handler: () => console.log('History') },
    { title: '文档中心', icon: 'mdi-file-document', handler: () => console.log('Docs') },
    { title: '系统设置', icon: 'mdi-cog', handler: () => console.log('Settings') }
    ])
</script>