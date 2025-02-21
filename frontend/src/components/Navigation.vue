<template>
  <v-navigation-drawer 
    image="https://cdn.vuetifyjs.com/images/backgrounds/bg-2.jpg"
    expand-on-hover
    rail
  >
    <v-list>
      <v-list-item
        prepend-avatar="https://ts1.cn.mm.bing.net/th?id=OIP-C.ZafPSQYlJ1dLRtxro-27agHaHa&w=250&h=250&c=8&rs=1&qlt=90&o=6&dpr=1.3&pid=3.1&rm=2"
        :subtitle="getUser().email"
        title="User"
      />
    </v-list>

    <v-divider />

    <v-list>
      <v-list-item
        prepend-icon="mdi-view-dashboard"
        title="我的主页"
        value="home"
        @click="router.push('./home')"
      />
      <v-list-item
        prepend-icon="mdi-folder"
        title="我的项目"
        value="folder"
        @click="router.push('./project')"
      />
      <v-list-item
        prepend-icon="mdi-account"
        title="账户中心"
        value="account"
        @click="router.push('./account')"
      />
    </v-list>
  </v-navigation-drawer>
</template>

<script setup lang="ts">
import { getUser } from '@/stores/user';
import { useRouter } from 'vue-router'

const router = useRouter()
</script>

<style scoped>
</style>