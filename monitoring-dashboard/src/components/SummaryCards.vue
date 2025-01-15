<template>
  <div class="summary-cards">
    <div class="card" v-for="(value, key) in data" :key="key">
      <h2>{{ key }}</h2>
      <p>{{ value }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'SummaryCards',
  props: {
    data: {
      type: Object,
      required: true,
    },
  },
});
</script>

<style scoped>
.summary-cards {
  display: flex;
  gap: 20px;
  justify-content: center;
  margin-bottom: 30px;
}

.card {
  background-color: #f9f9f9;
  padding: 16px;
  border: 1px solid #ddd;
  border-radius: 8px;
  text-align: center;
}
</style>
