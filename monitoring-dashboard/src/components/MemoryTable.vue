<template>
  <table class="memory-table">
    <thead>
    <tr>
      <th>Metric</th>
      <th>Value</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="(value, key) in data" :key="key">
      <td>{{ key }}</td>
      <td>{{ value }}</td>
    </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'MemoryTable',
  props: {
    data: {
      type: Object,
      required: true,
    },
  },
});
</script>

<style scoped>
.memory-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

th,
td {
  border: 1px solid #ddd;
  padding: 8px;
}

th {
  background-color: #f4f4f4;
  text-align: left;
}
</style>
