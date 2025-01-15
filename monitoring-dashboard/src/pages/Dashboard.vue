<template>
  <div class="container mx-auto px-4">
    <h1 class="text-3xl font-bold mb-6 text-center">System Usage Dashboard</h1>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
      <SummaryCard v-for="(value, key) in summary" :key="key" :title="key" :value="value" />
    </div>

    <!-- Charts -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- 메모리 차트 -->
      <PieChart :data="memoryData" />
      <!-- CPU 차트 -->
      <PieChart :data="cpuData" />
    </div>

    <!-- Detailed Tables -->
    <MemoryTable :data="memoryData" />
    <MemoryTable :data="cpuData" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref, onMounted } from 'vue';
import SummaryCard from '@/components/SummaryCards.vue';
import PieChart from '@/components/PieChart.vue';
import LineChart from '@/components/LineChart.vue';
import MemoryTable from '@/components/MemoryTable.vue';
import { useMemoryStore } from '@/store/memory';
import { useCpuStore } from '@/store/cpu';
import { fetchCpuData, fetchMemData } from '@/api/index';

export default defineComponent({
  name: 'Dashboard',
  components: { SummaryCard, PieChart, LineChart, MemoryTable },
  setup() {
    const memoryStore = useMemoryStore();
    const cpuStore = useCpuStore();

    // CPU 상태
    const isCpuLoading = ref<boolean>(false);
    const cpuError = ref<string | null>(null);

    // Memory 상태
    const isMemoryLoading = ref<boolean>(false);
    const memoryError = ref<string | null>(null);

    // CPU 데이터 로드
    const loadCpuData = async () => {
      isCpuLoading.value = true;
      cpuError.value = null;
      try {
        const data = await fetchCpuData();
        console.log('Fetched CpuData:', data);
        cpuStore.setCpuData(data);
      } catch (err) {
        console.error('Error fetching CpuData:', err);
        cpuError.value = 'Failed to fetch CPU data.';
      } finally {
        isCpuLoading.value = false;
      }
    };

    // Memory 데이터 로드
    const loadMemoryData = async () => {
      isMemoryLoading.value = true;
      memoryError.value = null;
      try {
        const data = await fetchMemData();
        console.log('Fetched MemoryData:', data);
        memoryStore.setMemoryData(data);
      } catch (err) {
        console.error('Error fetching MemoryData:', err);
        memoryError.value = 'Failed to fetch Memory data.';
      } finally {
        isMemoryLoading.value = false;
      }
    };

    // 마운트 후 데이터 로드
    onMounted(() => {
      console.log('Dashboard component mounted');
      loadCpuData();
      loadMemoryData();
    });

    return {
      memoryData: computed(() => memoryStore.memoryData || { used: 0, free: 0 }),
      cpuData: computed(() => cpuStore.cpuData || { used: 0, free: 0 }),
      summary: computed(() => ({
        TotalMemory: memoryStore.memoryData?.total || 'N/A',
        UsedMemory: memoryStore.memoryData?.used || 'N/A',
        FreeMemory: memoryStore.memoryData?.free || 'N/A',
        TotalCPU: cpuStore.cpuData?.total || 'N/A',
        UsedCPU: cpuStore.cpuData?.used || 'N/A',
      })),
      memoryHistory: computed(() => memoryStore.memoryHistory || []),
    };
  },
});
</script>

<style scoped>
/* 추가 스타일이 필요한 경우 */
</style>
