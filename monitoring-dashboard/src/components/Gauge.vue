<template>
  <div class="gauge-container">
    <div
        class="gauge"
        :style="{
        background: `conic-gradient(
          #4caf50 ${percentUsed}%,
          #ddd ${percentUsed}%
        )`
      }"
    >
      <div class="gauge-cover">
        <div class="gauge-text">
          <div>{{ used.toFixed(2) }} / {{ total.toFixed(2) }} GiB</div>
          <div>{{ percentUsed }}%</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'Gauge',
  props: {
    used: {
      type: Number,
      required: true,
    },
    total: {
      type: Number,
      required: true,
    },
    percentUsed: {
      type: Number,
      required: true,
    },
  },
});
</script>

<style scoped>
.gauge-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 200px;
}

.gauge {
  position: relative;
  width: 150px;
  height: 150px;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.gauge-cover {
  position: absolute;
  width: 110px;
  height: 110px;
  background: #fff;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.gauge-text {
  font-size: 14px;
  font-weight: bold;
  color: #333;
  text-align: center;
}
</style>