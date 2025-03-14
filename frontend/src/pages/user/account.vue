<template>
  <v-app>
    <v-breadcrumbs
      :items="['账户中心','']"
      style="font-size: small;"
    />
    <v-divider />
  </v-app>
</template>

<script setup lang="ts">

</script>

<style scoped>
*{
  background-color: #fff;
  color: #111;
}
.v-card {
  transition: transform 0.2s;
}
.v-card:hover {
  transform: translateY(-2px);
}
</style>
