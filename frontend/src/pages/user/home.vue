<template>
  <v-app>
    <v-breadcrumbs
      :items="['我的主页','']"
      style="font-size: small;"
    />
    <v-divider />
    <v-main>
      <v-container class="py-8">
        <QuickAct />
      </v-container>
    </v-main>
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
