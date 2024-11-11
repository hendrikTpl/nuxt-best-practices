<template>
  <v-app>
    <v-navigation-drawer mobile-breakpoint="xs" :rail="rail" expand-on-hover>
      <template v-slot:prepend>
        <v-btn class="ma-1" icon="mdi-menu" @click="changeRail()"></v-btn>
      </template>
      <v-list nav>
        <v-list-item prepend-icon="mdi-view-dashboard" title="Home" value="Home" :active="route.path == '/'"
          to="/">
        </v-list-item>
        <v-list-item prepend-icon="mdi-console" title="settings" value="settings" :active="route.path == '/settings'"
          to="/settings">
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar class="bg-surface" flat height="72">
      <v-app-bar-title class="d-flex flex-1-1-100 justify-center">
        <NuxtLink to="/">
          <v-img src="/deepx.id_200x200.png" width="200"></v-img>
        </NuxtLink>
      </v-app-bar-title>
    </v-app-bar>
    <v-main>
      <slot></slot>
    </v-main>
  </v-app>
</template>
<script lang="ts" setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue';
const route = useRoute();
const router = useRouter();
const rail = ref(true);
const navigateTo = (path :any) => {
  router.push(path);
};
const changeRail = () => {
  rail.value = !rail.value;
};
</script>