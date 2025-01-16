<template>
  <div class="container mx-auto px-6 py-8 bg-gray-100">
    <h1 class="text-4xl font-extrabold text-center text-gray-800 mb-10">
      시스템 사용 대시보드
    </h1>

    <!-- 차트 섹션 -->
    <div class="flex flex-wrap justify-center gap-8 mb-12">
      <!-- CPU 사용 차트 -->
      <div class="card">
        <h2 class="card-title">CPU 사용량</h2>
        <LineChart :labels="cpuChartLabels" :data="cpuChartData" />
      </div>

      <!-- Memory 게이지 -->
      <div class="card">
        <h2 class="card-title">Memory Usage</h2>
        <Gauge :used="memoryData.used" :total="memoryData.total" :percentUsed="memoryData.percentUsed" />
      </div>

      <!-- 네트워크 사용량 차트 -->
      <div class="card">
        <h2 class="card-title">네트워크 사용량</h2>
        <BarChart :labels="networkChartLabels" :sent="networkChartSent" :recv="networkChartRecv" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import LineChart from '@/components/LineChart.vue';
import Gauge from '@/components/Gauge.vue';
import BarChart from '@/components/BarChart.vue';
import { processNetData } from '@/util/netUtils';
import { processCpuData } from '@/util/cpuUtils';
import { processMemData } from '@/util/memUtils';
import { fetchCpuData, fetchMemData, fetchNetData } from '@/api/index';

export default defineComponent({
  name: 'Dashboard',
  components: { LineChart, Gauge, BarChart },

  setup() {
    const cpuChartLabels = ref<string[]>([]);
    const cpuChartData = ref<number[]>([]);
    const memoryData = ref({ total: 0, used: 0, percentUsed: 0 });

    const networkChartLabels = ref<string[]>([]); // 시간 라벨
    const networkChartSent = ref<number[]>([]); // 전송량
    const networkChartRecv = ref<number[]>([]); // 수신량

    const loadCpuData = async () => {
      const rawData = await fetchCpuData();
      const processedData = processCpuData(rawData);
      cpuChartLabels.value = processedData.map((entry) => entry.time);
      cpuChartData.value = processedData.map((entry) => entry.averageUsage);
    };

    const loadMemoryData = async () => {
      const rawData = await fetchMemData();
      memoryData.value = processMemData(rawData.data);
    };

    const loadNetworkData = async () => {
      try {
        const rawData = await fetchNetData(); // 원본 데이터 가져오기
        console.log('Raw Network Data:', rawData); // 원본 데이터 확인
        const processedData = processNetData(rawData); // 데이터 가공
        console.log('Processed Network Data:', processedData); // 가공된 데이터 확인

        networkChartLabels.value = processedData.map((entry) => entry.time);
        networkChartSent.value = processedData.map((entry) => entry.sent);
        networkChartRecv.value = processedData.map((entry) => entry.recv);
      } catch (error) {
        console.error('네트워크 데이터를 가져오는 중 오류 발생:', error);
      }
    };

    onMounted(() => {
      loadCpuData();
      loadMemoryData();
      loadNetworkData();
    });

    return {
      cpuChartLabels,
      cpuChartData,
      memoryData,
      networkChartLabels,
      networkChartSent,
      networkChartRecv,
    };
  },
});
</script>

<style scoped>
.container {
  max-width: 90rem; /* 1440px */
  margin: 0 auto;
  padding: 2rem;
  background-color: #f8f9fa;
}

.card {
  background-color: white;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border-radius: 0.5rem;
  padding: 1.5rem;
  min-height: 21.875rem;
  width: 36rem;
}

.card-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #4a5568;
  margin-bottom: 1rem;
}
</style>