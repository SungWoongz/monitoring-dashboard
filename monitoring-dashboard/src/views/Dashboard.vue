<template>
  <div>
    <h1>CPU Usage Dashboard</h1>
    <p v-if="isLoading">Loading...</p>
    <p v-if="error">{{ error }}</p>
    <p v-if="!isLoading && !cpuData">No data available</p>
    <table v-if="cpuData && Object.keys(cpuData).length > 0 && !isLoading">
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
  </div>
</template>

<script lang="ts">
import {defineComponent, ref, onMounted, computed} from 'vue';
import {fetchCpuData} from '@/api/index';
import {useCpuStore} from '@/store/cpu';

export default defineComponent({
  name: 'Dashboard',
  setup() {
    const cpuStore = useCpuStore();
    const isLoading = ref<boolean>(false);
    const error = ref<string | null>(null);

    const loadCpuData = async () => {
      isLoading.value = true;
      error.value = null;
      try {
        const data = await fetchCpuData();
        console.log('Fetched data:', data);
        cpuStore.setCpuData(data);
      } catch (err) {
        console.error('Error fetching data:', err);
        error.value = 'Failed to fetch CPU data.';
      } finally {
        isLoading.value = false;
      }
    };

    onMounted(() => {
      console.log('Dashboard component mounted');
      loadCpuData();
    });

    return {
      cpuData: computed(() => cpuStore.cpuData), // 반응형 상태로 설정
      isLoading,
      error,
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