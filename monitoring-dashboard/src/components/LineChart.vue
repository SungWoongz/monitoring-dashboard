<template>
  <div>
    <canvas id="line-chart"></canvas>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, watch } from 'vue';
import { Chart } from 'chart.js';

export default defineComponent({
  name: 'LineChart',
  props: {
    history: {
      type: Array,
      required: true,
    },
  },
  setup(props) {
    let chart: Chart | null = null;

    const renderChart = () => {
      const ctx = document.getElementById('line-chart') as HTMLCanvasElement;

      if (!ctx) {
        console.error('Canvas element not found');
        return;
      }

      // 기존 차트를 제거 (중복 생성 방지)
      if (chart) {
        chart.destroy();
      }

      chart = new Chart(ctx, {
        type: 'line',
        data: {
          labels: props.history.map((_, index) => `Time ${index}`),
          datasets: [
            {
              label: 'Memory Usage',
              data: props.history,
              borderColor: '#36A2EB',
              fill: false,
            },
          ],
        },
      });
    };

    onMounted(() => {
      if (props.history && props.history.length > 0) {
        renderChart();
      } else {
        console.warn('History data is empty or undefined');
      }
    });

    // props.history 변경 시 차트 다시 렌더링
    watch(
        () => props.history,
        (newHistory) => {
          if (newHistory && newHistory.length > 0) {
            renderChart();
          }
        },
        { immediate: true }
    );

    return {};
  },
});
</script>

<style scoped>
canvas {
  max-width: 800px;
  margin: auto;
}
</style>
