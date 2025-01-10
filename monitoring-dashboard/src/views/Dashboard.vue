<template>
  <div>
    <h1>System Usage Dashboard</h1>

    <!-- CPU 데이터 섹션 -->
    <h2>CPU Data</h2>
    <p v-if="isCpuLoading">Loading CPU data...</p>
    <p v-if="cpuError">{{ cpuError }}</p>
    <p v-if="!isCpuLoading && !cpuData">No CPU data available</p>
    <table v-if="cpuData && Object.keys(cpuData).length > 0 && !isCpuLoading">
      <thead>
      <tr>
        <th>Metric</th>
        <th>Value</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(value, key) in cpuData" :key="key">
        <td>{{ key }}</td>
        <td>{{ value }}</td>
      </tr>
      </tbody>
    </table>

    <!-- Memory 데이터 섹션 -->
    <h2>Memory Data</h2>
    <p v-if="isMemoryLoading">Loading Memory data...</p>
    <p v-if="memoryError">{{ memoryError }}</p>
    <p v-if="!isMemoryLoading && !memoryData">No Memory data available</p>
    <table v-if="memoryData && Object.keys(memoryData).length > 0 && !isMemoryLoading">
      <thead>
      <tr>
        <th>Metric</th>
        <th>Value</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(value, key) in memoryData" :key="key">
        <td>{{ key }}</td>
        <td>{{ value }}</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref, onMounted, computed} from 'vue';
import {fetchCpuData, fetchMemData} from '@/api/index';
import {useCpuStore} from '@/store/cpu';
import {useMemoryStore} from "@/store/memory";

export default defineComponent({
  name: 'Dashboard',
  setup() {
    const cpuStore = useCpuStore();
    const memoryStore = useMemoryStore();

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
      // CPU 상태
      cpuData: computed(() => cpuStore.cpuData),
      isCpuLoading,
      cpuError,

      // Memory 상태
      memoryData: computed(() => memoryStore.memoryData),
      isMemoryLoading,
      memoryError,
    };
  },
});
</script>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
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