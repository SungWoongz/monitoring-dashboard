<template>
  <div>
    <canvas id="pie-chart"></canvas>
    <p v-if="!hasValidData" class="no-data-message">No data available to display</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, computed, watch } from 'vue';
import { Chart, PieController, ArcElement, Tooltip, Legend } from 'chart.js';

Chart.register(PieController, ArcElement, Tooltip, Legend);

export default defineComponent({
  name: 'PieChart',
  props: {
    data: {
      type: Object,
      required: true,
    },
  },
  setup(props) {
    let chart: Chart | null = null;

    const hasValidData = computed(() => {
      return (
          props.data &&
          typeof props.data.used === 'number' &&
          typeof props.data.free === 'number' &&
          (props.data.used > 0 || props.data.free > 0)
      );
    });

    const createChart = (ctx: HTMLCanvasElement) => {
      if (chart) {
        chart.destroy();
      }

      chart = new Chart(ctx, {
        type: 'pie',
        data: {
          labels: ['Used', 'Free'],
          datasets: [
            {
              data: [props.data.used, props.data.free],
              backgroundColor: ['#FF6384', '#36A2EB'],
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
    };

    onMounted(() => {
      const ctx = document.getElementById('pie-chart') as HTMLCanvasElement;

      if (!ctx) {
        console.error('Canvas element not found');
        return;
      }

      if (hasValidData.value) {
        createChart(ctx);
      } else {
        console.warn('No valid data to render chart');
      }
    });

    watch(
        () => props.data,
        (newData) => {
          const ctx = document.getElementById('pie-chart') as HTMLCanvasElement;

          if (!ctx) {
            console.error('Canvas element not found');
            return;
          }

          if (hasValidData.value) {
            createChart(ctx);
          } else if (chart) {
            chart.destroy();
            chart = null;
          }
        },
        { deep: true }
    );

    return {
      hasValidData,
    };
  },
});
</script>

<style scoped>
canvas {
  max-width: 400px;
  margin: auto;
}
.no-data-message {
  text-align: center;
  color: #888;
  font-size: 16px;
  margin: 20px 0;
}
</style>
