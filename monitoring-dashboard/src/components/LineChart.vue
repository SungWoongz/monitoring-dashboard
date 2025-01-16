<template>
  <div>
    <canvas ref="canvas" class="line-chart"></canvas>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, watch } from 'vue';
import { Chart, LineController, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend } from 'chart.js';

Chart.register(LineController, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

export default defineComponent({
  name: 'LineChart',
  props: {
    labels: {
      type: Array as () => string[],
      required: true,
    },
    data: {
      type: Array as () => number[],
      required: true,
    },
  },
  setup(props) {
    const canvas = ref<HTMLCanvasElement | null>(null);
    let chart: Chart | null = null;

    const createChart = () => {
      if (chart) {
        chart.destroy();
      }

      if (canvas.value) {
        chart = new Chart(canvas.value, {
          type: 'line',
          data: {
            labels: props.labels,
            datasets: [
              {
                label: 'CPU 사용량 평균 (%)',
                data: props.data,
                borderColor: '#36A2EB',
                backgroundColor: 'rgba(54,162,235,0.2)',
                fill: true,
              },
            ],
          },
          options: {
            responsive: true,
            plugins: {
              legend: {
                position: 'top',
              },
              tooltip: {
                enabled: true,
              },
            },
          },
        });
      }
    };

    watch(
        () => [props.labels, props.data],
        () => createChart(),
        { deep: true }
    );

    onMounted(() => {
      createChart();
    });

    return {
      canvas,
    };
  },
});
</script>

<style scoped>
.line-chart {
  max-width: 100%;
  margin: auto;
}
</style>