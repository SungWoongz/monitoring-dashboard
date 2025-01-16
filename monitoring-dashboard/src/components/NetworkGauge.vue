<template>
  <div id="network-gauge" style="width: 100%; height: 200px;"></div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, watch } from 'vue';
import * as echarts from 'echarts';

export default defineComponent({
  name: 'NetworkGauge',
  props: {
    sentSpeed: {
      type: Number,
      required: true, // 초당 송신 속도 (Mbps)
    },
    recvSpeed: {
      type: Number,
      required: true, // 초당 수신 속도 (Mbps)
    },
    maxSpeed: {
      type: Number,
      required: true, // 최대 속도 (Mbps)
    },
  },
  setup(props) {
    const chart = ref<echarts.ECharts | null>(null);

    const initChart = () => {
      const chartDom = document.getElementById('network-gauge');
      if (!chartDom) return;

      chart.value = echarts.init(chartDom);

      const option = {
        series: [
          {
            type: 'gauge',
            startAngle: 180,
            endAngle: 0,
            min: 0,
            max: props.maxSpeed,
            splitNumber: 10,
            axisLine: {
              lineStyle: {
                width: 10,
                color: [
                  [0.3, '#67e0e3'], // 안전
                  [0.7, '#ffdb5c'], // 주의
                  [1, '#fd666d'],   // 위험
                ],
              },
            },
            pointer: {
              show: true,
            },
            axisLabel: {
              distance: 15,
              fontSize: 12,
            },
            title: {
              offsetCenter: ['0%', '-30%'],
              fontSize: 16,
            },
            detail: {
              valueAnimation: true,
              formatter: '{value} Mbps',
            },
            data: [
              { value: props.recvSpeed, name: '수신 속도' },
              { value: props.sentSpeed, name: '송신 속도' },
            ],
          },
        ],
      };

      chart.value.setOption(option);
    };

    watch([() => props.sentSpeed, () => props.recvSpeed], () => {
      if (chart.value) {
        chart.value.setOption({
          series: [
            {
              data: [
                { value: props.recvSpeed, name: '수신 속도' },
                { value: props.sentSpeed, name: '송신 속도' },
              ],
            },
          ],
        });
      }
    });

    onMounted(() => {
      initChart();
    });

    return {};
  },
});
</script>

<style scoped>
#network-gauge {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>