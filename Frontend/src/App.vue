<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="action">    
    <div class="router-view">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
.action {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  width: 100%;
}

.router-view {
  margin-top: 6px;
}

header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 60%;
  font-size: 18px;
  text-align: left;
  margin: 0;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
  font-size: 32px;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  nav {
    text-align: left;
    font-size: 1rem;
  }
}
</style>
