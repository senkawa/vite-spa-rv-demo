<script setup lang="ts"></script>

<template>
  <div class="m-auto w-60 text-sm pt-5">
    <nav class="border-b-2 pb-3">
      <router-link to="/manage">Manage</router-link> |
      <router-link to="/notes">Notes</router-link>
    </nav>
    <router-view></router-view>
  </div>
</template>

<style>
a {
  @apply text-orange-800 hover:text-orange-900;
}
</style>
