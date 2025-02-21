<template>
  <v-app>
    <v-card>
      <v-layout>
        <NavBar />
        <Navigation />
        <v-main style="height: 100vh;">
          <tr>
            <td>
              <v-btn
                icon="mdi-chevron-left"
                density="compact"
                style="position: absolute; top: 95vh; left: 6vw; background-color: black; z-index: 2000;"
                @click="router.go(-1)"
              />
            </td>
            <td>
              <v-btn
                icon="mdi-chevron-right"
                density="compact"
                style="position: absolute; top: 95vh; left: 10vw; background-color: black; z-index: 2000;"
                @click="router.go(+1)"
              />
            </td>
          </tr>
          
          <router-view />
        </v-main>
      </v-layout>
    </v-card>
  </v-app>
</template>

<script setup lang="ts">
  import NavBar from '@/components/NavBar.vue';
  import Navigation from '@/components/Navigation.vue';
  import router from '@/router';
</script>

<style scoped>

</style>