<template>
  <header class="header">
    <h1>VM Server Dashboard</h1>
    <p>Last updated: {{ lastUpdated }}</p>
  </header>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'Header',
  props: {
    lastUpdated: {
      type: String,
      required: true,
    },
  },
});
</script>

<style scoped>
.header {
  text-align: center;
  margin-bottom: 20px;
}
</style>
