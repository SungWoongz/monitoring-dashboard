<template>
  <div>
    <h1>CPU Usage Dashboard</h1>
    <!-- 로딩 상태 -->
    <p v-if="isLoading">Loading...</p>
    <!-- 에러 상태 -->
    <p v-if="error">{{ error }}</p>
    <!-- CPU 데이터 테이블 -->
    <table v-if="cpuData && !isLoading">
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
import {defineComponent, ref, onMounted} from 'vue';
import {fetchCpuData} from '@/api/index';
import {useCpuStore} from '@/store/cpu';

export default defineComponent({
  name: 'Dashboard',
  setup() {
    const cpuStore = useCpuStore();
    const isLoading = ref<boolean>(false);
    const error = ref<string | null>(null);

    // 데이터 로딩 함수
    const loadCpuData = async () => {
      console.log('Starting API call...');
      isLoading.value = true; // 로딩 상태 시작
      error.value = null; // 기존 에러 초기화
      try {
        const data = await fetchCpuData(); // API 호출
        console.log('Data fetched:', data);
        cpuStore.setCpuData(data); // 스토어 업데이트
      } catch (err) {
        console.error('Error fetching data:', err);
        error.value = 'Failed to fetch CPU data.'; // 에러 메시지 설정
      } finally {
        isLoading.value = false; // 로딩 상태 종료
      }
    };

    // 마운트 후 데이터 로드
    onMounted(() => {
      console.log('Dashboard component mounted');
      loadCpuData();
    });

    return {
      cpuData: cpuStore.cpuData, // 스토어에서 데이터 가져오기
      isLoading,
      error,
    };
  },
});
</script>

<style scoped>
/* 테이블 스타일 */
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