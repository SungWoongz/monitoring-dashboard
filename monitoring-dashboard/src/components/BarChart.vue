<script lang="ts">
import {defineComponent, onMounted} from 'vue';
import {Chart, BarController, BarElement, CategoryScale, LinearScale} from 'chart.js';

Chart.register(BarController, BarElement, CategoryScale, LinearScale);

export default defineComponent({
  name: 'BarChart',
  props: {
    labels: {
      type: Array as () => string[],
      required: true,
    },
    sent: {
      type: Array as () => number[],
      required: true,
    },
    recv: {
      type: Array as () => number[],
      required: true,
    },
  },
  setup(props) {
    onMounted(() => {
      console.log('Labels:', props.labels);
      console.log('Sent Data:', props.sent);
      console.log('Received Data:', props.recv);

      const ctx = document.getElementById('bar-chart') as HTMLCanvasElement;

      new Chart(ctx, {
        type: 'bar',
        data: {
          labels: props.labels, // X축 레이블
          datasets: [
            {
              label: 'Bytes Sent (MB)',
              data: props.sent, // 전송량
              backgroundColor: '#FF6384',
            },
            {
              label: 'Bytes Received (MB)',
              data: props.recv, // 수신량
              backgroundColor: '#36A2EB',
            },
          ],
        },
        options: {
          responsive: true,
          scales: {
            x: {
              title: {
                display: true,
                text: 'Time',
              },
            },
            y: {
              title: {
                display: true,
                text: 'Usage (MB)',
              },
              beginAtZero: true,
            },
          },
        },
      });
    });
  },
});
</script>
<template>
  <div>
    <canvas id="bar-chart"></canvas>
  </div>
</template>